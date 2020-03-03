package libra

import (
	"encoding/json"
	"fmt"
)

const (
	accResourceKey = "0116608f05d24742a043e6fd12d3b32735f6bfcba287bea92b28a175cd4f3eee32"
)

// AccountState account state, see https://github.com/libra/libra/blob/master/types/src/account_config.rs
type AccountState struct {
	Blob            []byte
	AccountResource AccountResource
}

// AccountResource account state detail, see https://github.com/libra/libra/blob/master/types/src/account_config.rs
type AccountResource struct {
	AuthKey            string `json:"authkey"`
	Balance            uint64 `json:"balance"`
	KeyRotationCap     bool   `json:"key_rotation_cap"`
	WithdrawalCap      bool   `json:"with_drawal_cap"`
	ReceivedEventCount uint64 `json:"recv_event_count"`
	ReceivedEventData  string `json:"recv_event_data"`
	SentEventCount     uint64 `json:"sent_event_count"`
	SentEventData      string `json:"sent_event_data"`
	SequenceNo         uint64 `json:"seq_no"`
	EventNo            uint64 `json:"event_no"`
}

// String account struct as string
func (ar AccountResource) String() string {
	prettyJSON, err := json.MarshalIndent(ar, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	return string(prettyJSON)
}

func decodeAccountStateBlob(accountStateBlob []byte) (AccountState, error) {
	result := AccountState{
		Blob: accountStateBlob,
	}

	canonicalSerializer := NewCanonicalSerializer(accountStateBlob)

	m, err := canonicalSerializer.ReadMap()
	if err != nil {
		return result, err
	}
	accResource, err := DecodeAccountResourceBlob(m[accResourceKey])
	if err != nil {
		return result, err
	}
	result.AccountResource = accResource

	return result, nil
}

func decodeAccountResourceBlob(accountResourceBlob []byte) (AccountResource, error) {
	result := AccountResource{}

	canonicalSerializer := NewCanonicalSerializer(accountResourceBlob)
	var err error
	result.AuthKey, err = canonicalSerializer.ReadString()
	if err != nil {
		return result, err
	}

	result.Balance, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}
	result.KeyRotationCap, err = canonicalSerializer.ReadBool()
	if err != nil {
		return result, err
	}
	result.WithdrawalCap, err = canonicalSerializer.ReadBool()
	if err != nil {
		return result, err
	}

	result.ReceivedEventCount, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.ReceivedEventData, err = canonicalSerializer.ReadString()
	if err != nil {
		return result, err
	}

	result.SentEventCount, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}
	result.SentEventData, err = canonicalSerializer.ReadString()
	if err != nil {
		return result, err
	}

	result.SequenceNo, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.EventNo, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	return result, nil
}
