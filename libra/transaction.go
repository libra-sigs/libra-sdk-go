package libra

import (
	"encoding/hex"
	"encoding/json"

	"fmt"
	"github.com/libra-sigs/libra-sdk-go/libra/rpc/types"
)

//GoLibraTransactionInfo: see more https://github.com/libra/libra/blob/master/types/src/transaction/mod.rs
//pub struct RawTransaction {
type GoLibraTransactionInfo struct {
	Version         uint64 `json:"version"`
	Sender          string `json:"sender"`
	TransactionType int32  `json:"type"`
	SequenceNo      uint64 `json:"seq_no"`
	MaxGasAmount    uint64 `json:"max_gas_amount"`
	GasUnitPrice    uint64 `json:"gas_unit_price"`
	ExpirationTime  uint64 `json:"expiration_time"`
	StateRootHash   string `json:"state_hash"`
	EventRootHash   string `json:"event_hash"`
}

//String TransactionInfo to string
func (txn GoLibraTransactionInfo) String() string {
	prettyJSON, err := json.MarshalIndent(txn, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}
	return string(prettyJSON)
}

func decodeTransaction(m *types.TransactionWithProof) (GoLibraTransactionInfo, error) {

	result := GoLibraTransactionInfo{}

	var txn *types.Transaction

	var err error
	txn = m.GetTransaction()
	blob := txn.GetTransaction()

	canonicalSerializer := NewCanonicalSerializer(blob)

	_, err = canonicalSerializer.ReadInt32()
	if err != nil {
		return result, err
	}
	result.Sender, err = canonicalSerializer.ReadXString(32)
	if err != nil {
		return result, err
	}

	result.SequenceNo, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.TransactionType, err = canonicalSerializer.ReadInt32()
	if err != nil {
		return result, err
	}

	result.MaxGasAmount, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.GasUnitPrice, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.ExpirationTime, err = canonicalSerializer.ReadUint64()
	if err != nil {
		return result, err
	}

	result.Version = m.GetVersion()
	result.StateRootHash = hex.EncodeToString(m.GetProof().GetTransactionInfo().GetStateRootHash())
	result.EventRootHash = hex.EncodeToString(m.GetProof().GetTransactionInfo().GetEventRootHash())
	if err != nil {
		return result, err
	}
	return result, nil
}
