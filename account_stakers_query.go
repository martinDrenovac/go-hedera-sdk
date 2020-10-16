package hedera

import "github.com/hashgraph/hedera-sdk-go/proto"

// AccountStakersQuery gets all of the accounts that are proxy staking to this account. For each of  them, the amount
// currently staked will be given. This is not yet implemented, but will be in a future version of the API.
type AccountStakersQuery struct {
	Query
	pb *proto.CryptoGetStakersQuery
}

// NewAccountStakersQuery creates an AccountStakersQuery query which can be used to construct and execute
// an AccountStakersQuery.
//
// It is recommended that you use this for creating new instances of an AccountStakersQuery
// instead of manually creating an instance of the struct.
func NewAccountStakersQuery() *AccountStakersQuery {
	header := proto.QueryHeader{}
	query := newQuery(true, &header)
	pb := proto.CryptoGetStakersQuery{Header: &header}
	query.pb.Query = &proto.Query_CryptoGetProxyStakers{
		CryptoGetProxyStakers: &pb,
	}

	return &AccountStakersQuery{
		Query: query,
		pb:    &pb,
	}
}

// SetAccountID sets the Account ID for which the stakers should be retrieved
func (query *AccountStakersQuery) SetAccountID(id AccountID) *AccountStakersQuery {
	query.pb.AccountID = id.toProtobuf()
	return query
}

func accountStakersQuery_mapResponseStatus(_ request, response response) Status {
	return Status(response.query.GetCryptoGetProxyStakers().Header.NodeTransactionPrecheckCode)
}

func accountStakersQuery_getMethod(_ request, channel *channel) method {
	return method{
		query: channel.getCrypto().GetStakersByAccountID,
	}
}

func (query *AccountStakersQuery) Execute(client *Client) ([]Transfer, error) {
	if client == nil || client.operator == nil {
		return []Transfer{}, errNoClientProvided
	}

	query.queryPayment = NewHbar(2)
	query.paymentTransactionID = TransactionIDGenerate(client.operator.accountID)

	var cost Hbar
	if query.queryPayment.tinybar != 0 {
		cost = query.queryPayment
	} else {
		cost = client.maxQueryPayment

		// actualCost := CostQuery()
	}

	err := query_generatePayments(&query.Query, client, cost)
	if err != nil {
		return []Transfer{}, err
	}

	resp, err := execute(
		client,
		request{
			query: &query.Query,
		},
		query_shouldRetry,
		query_makeRequest,
		query_advanceRequest,
		query_getNodeId,
		accountStakersQuery_getMethod,
		accountStakersQuery_mapResponseStatus,
		query_mapResponse,
	)

	if err != nil {
		return []Transfer{}, err
	}

	var stakers = make([]Transfer, len(resp.query.GetCryptoGetProxyStakers().Stakers.ProxyStaker))

	for i, element := range resp.query.GetCryptoGetProxyStakers().Stakers.ProxyStaker {
		stakers[i] = Transfer{
			AccountID: accountIDFromProtobuf(element.AccountID),
			Amount:    HbarFromTinybar(element.Amount),
		}
	}

	return stakers, err
}

// SetMaxQueryPayment sets the maximum payment allowed for this Query.
func (query *AccountStakersQuery) SetMaxQueryPayment(maxPayment Hbar) *AccountStakersQuery {
	query.Query.SetMaxQueryPayment(maxPayment)
	return query
}

// SetQueryPayment sets the payment amount for this Query.
func (query *AccountStakersQuery) SetQueryPayment(paymentAmount Hbar) *AccountStakersQuery {
	query.Query.SetQueryPayment(paymentAmount)
	return query
}

func (query *AccountStakersQuery) SetNodeAccountID(accountID AccountID) *AccountStakersQuery {
	query.Query.SetNodeAccountID(accountID)
	return query
}

func (query *AccountStakersQuery) GetNodeAccountId() AccountID {
	return query.Query.GetNodeAccountId()
}