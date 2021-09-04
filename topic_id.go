package hedera

import (
	"fmt"

	"github.com/hashgraph/hedera-sdk-go/v2/proto"
	protobuf "google.golang.org/protobuf/proto"
)

// TopicID is a unique identifier for a topic (used by the  service)
type TopicID struct {
	Shard    uint64
	Realm    uint64
	Topic    uint64
	checksum *string
}

// TopicIDFromString constructs a TopicID from a string formatted as `Shard.Realm.Topic` (for example "0.0.3")
func TopicIDFromString(data string) (TopicID, error) {
	shard, realm, num, checksum, err := _IdFromString(data)
	if err != nil {
		return TopicID{}, err
	}

	return TopicID{
		Shard:    uint64(shard),
		Realm:    uint64(realm),
		Topic:    uint64(num),
		checksum: checksum,
	}, nil
}

func (id *TopicID) Validate(client *Client) error {
	if !id._IsZero() && client != nil && client.network.networkName != nil {
		tempChecksum, err := _ChecksumParseAddress(client.network.networkName._LedgerID(), fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.Topic))
		if err != nil {
			return err
		}
		err = _ChecksumVerify(tempChecksum.status)
		if err != nil {
			return err
		}
		if id.checksum == nil {
			id.checksum = &tempChecksum.correctChecksum
			return nil
		}
		if tempChecksum.correctChecksum != *id.checksum {
			return errNetworkMismatch
		}
	}

	return nil
}

func (id *TopicID) _SetNetworkWithClient(client *Client) {
	if client.network.networkName != nil {
		id._SetNetwork(*client.network.networkName)
	}
}

func (id *TopicID) _SetNetwork(name NetworkName) {
	checksum := _CheckChecksum(name._LedgerID(), fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.Topic))
	id.checksum = &checksum
}
func (id TopicID) _IsZero() bool {
	return id.Shard == 0 && id.Realm == 0 && id.Topic == 0
}

// String returns the string representation of a TopicID in `Shard.Realm.Topic` (for example "0.0.3")
func (id TopicID) String() string {
	return fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.Topic)
}

func (id TopicID) ToStringWithChecksum(client Client) (string, error) {
	if client.network.networkName == nil {
		return "", errNetworkNameMissing
	}
	checksum, err := _ChecksumParseAddress(client.network.networkName._LedgerID(), fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.Topic))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d-%s", id.Shard, id.Realm, id.Topic, checksum.correctChecksum), nil
}

func (id TopicID) _ToProtobuf() *proto.TopicID {
	return &proto.TopicID{
		ShardNum: int64(id.Shard),
		RealmNum: int64(id.Realm),
		TopicNum: int64(id.Topic),
	}
}

func _TopicIDFromProtobuf(topicID *proto.TopicID) *TopicID {
	if topicID == nil {
		return nil
	}

	return &TopicID{
		Shard: uint64(topicID.ShardNum),
		Realm: uint64(topicID.RealmNum),
		Topic: uint64(topicID.TopicNum),
	}
}

func (id TopicID) ToBytes() []byte {
	data, err := protobuf.Marshal(id._ToProtobuf())
	if err != nil {
		return make([]byte, 0)
	}

	return data
}

func TopicIDFromBytes(data []byte) (TopicID, error) {
	if data == nil {
		return TopicID{}, errByteArrayNull
	}
	pb := proto.TopicID{}
	err := protobuf.Unmarshal(data, &pb)
	if err != nil {
		return TopicID{}, err
	}

	return *_TopicIDFromProtobuf(&pb), nil
}
