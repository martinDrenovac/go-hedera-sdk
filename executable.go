package hedera

import (
	"context"
	"math"
	"time"

	"github.com/pkg/errors"

	"github.com/hashgraph/hedera-sdk-go/v2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxAttempts = 10

type _ExecutionState uint32

const (
	executionStateRetry    _ExecutionState = 0
	executionStateFinished _ExecutionState = 1
	executionStateError    _ExecutionState = 2
)

type _Method struct {
	query func(
		context.Context,
		*proto.Query,
		...grpc.CallOption,
	) (*proto.Response, error)
	transaction func(
		context.Context,
		*proto.Transaction,
		...grpc.CallOption,
	) (*proto.TransactionResponse, error)
}

type _Response struct {
	query       *proto.Response
	transaction *proto.TransactionResponse
}

type _IntermediateResponse struct {
	query       *proto.Response
	transaction TransactionResponse
}

type _ProtoRequest struct {
	query       *proto.Query
	transaction *proto.Transaction
}

type _Request struct {
	query       *Query
	transaction *Transaction
}

func _Execute(
	client *Client,
	request _Request,
	shouldRetry func(_Request, _Response) _ExecutionState,
	protoReq _ProtoRequest,
	advanceRequest func(_Request),
	getNodeAccountID func(_Request) AccountID,
	getMethod func(_Request, *_Channel) _Method,
	mapStatusError func(_Request, _Response) error,
	mapResponse func(_Request, _Response, AccountID, _ProtoRequest) (_IntermediateResponse, error),
) (_IntermediateResponse, error) {
	var maxAttempts int
	var minBackoff *time.Duration
	var maxBackoff *time.Duration

	if client.maxAttempts != nil {
		maxAttempts = *client.maxAttempts
	} else {
		if request.query != nil {
			maxAttempts = request.query.maxRetry
		} else {
			maxAttempts = request.transaction.maxRetry
		}
	}

	if request.query != nil {
		if request.query.maxBackoff == nil {
			maxBackoff = &client.maxBackoff
		} else {
			maxBackoff = request.query.maxBackoff
		}
		if request.query.minBackoff == nil {
			minBackoff = &client.minBackoff
		} else {
			minBackoff = request.query.minBackoff
		}
	} else {
		if request.transaction.maxBackoff == nil {
			maxBackoff = &client.maxBackoff
		} else {
			maxBackoff = request.transaction.maxBackoff
		}
		if request.transaction.minBackoff == nil {
			minBackoff = &client.minBackoff
		} else {
			minBackoff = request.transaction.minBackoff
		}
	}

	var attempt int64
	var errPersistent error

	for attempt = int64(0); attempt < int64(maxAttempts); attempt++ {
		protoRequest := protoReq
		nodeAccountID := getNodeAccountID(request)

		node, ok := client.network.networkNodes[nodeAccountID]
		if !ok {
			return _IntermediateResponse{}, ErrInvalidNodeAccountIDSet{nodeAccountID}
		}

		node._InUse()

		channel, err := node._GetChannel()
		if err != nil {
			node._IncreaseDelay()
			continue
		}

		method := getMethod(request, channel)

		advanceRequest(request)

		resp := _Response{}

		if !node._IsHealthy() {
			node._Wait()
		}

		if method.query != nil {
			resp.query, err = method.query(context.TODO(), protoRequest.query)
		} else {
			resp.transaction, err = method.transaction(context.TODO(), protoRequest.transaction)
		}

		if err != nil {
			errPersistent = err
			if _ExecutableDefaultRetryHandler(err) {
				node._IncreaseDelay()
				continue
			}
			return _IntermediateResponse{}, errors.Wrapf(errPersistent, "retry %d/%d", attempt, maxAttempts)
		}

		node._DecreaseDelay()

		retry := shouldRetry(request, resp)

		switch retry {
		case executionStateRetry:
			if attempt <= int64(maxAttempts) {
				_DelayForAttempt(minBackoff, maxBackoff, attempt)
				continue
			} else {
				errPersistent = mapStatusError(request, resp)
				break
			}
		case executionStateError:
			return _IntermediateResponse{}, mapStatusError(request, resp)
		case executionStateFinished:
			return mapResponse(request, resp, node.accountID, protoRequest)
		}
	}

	return _IntermediateResponse{}, errors.Wrapf(errPersistent, "retry %d/%d", attempt, maxAttempts)
}

func _DelayForAttempt(minBackoff *time.Duration, maxBackoff *time.Duration, attempt int64) {
	// 0.1s, 0.2s, 0.4s, 0.8s, ...
	ms := int64(math.Min(float64(minBackoff.Milliseconds())*math.Pow(2, float64(attempt)), float64(maxBackoff.Milliseconds())))
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func _ExecutableDefaultRetryHandler(err error) bool {
	code := status.Code(err)

	switch code {
	case codes.ResourceExhausted, codes.Unavailable:
		return true
	case codes.Internal:
		grpcErr, ok := status.FromError(err)

		if !ok {
			return false
		}

		return rstStream.FindIndex([]byte(grpcErr.Message())) != nil
	default:
		return false
	}
}
