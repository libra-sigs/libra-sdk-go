package libra

import (
	"context"
	"encoding/hex"

	"time"

	"google.golang.org/grpc"

	"github.com/libra-sigs/libra-sdk-go/libra/rpc/admission_control"
	"github.com/libra-sigs/libra-sdk-go/libra/rpc/types"
)

// sdk client
type Client struct {
	address string
	conn    *grpc.ClientConn
	acc     admission_control.AdmissionControlClient
}

// close
func (c Client) Close() {
	c.conn.Close()
}

//
func NewClient(address string, dialTimeout time.Duration) (Client, error) {
	ctxWithTimeout, cancelFunc := context.WithTimeout(context.Background(), dialTimeout)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctxWithTimeout, address, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return Client{}, err
	}
	acc := admission_control.NewAdmissionControlClient(conn)
	return Client{
		address: address,
		conn:    conn,
		acc:     acc,
	}, nil
}

// query account state
func (c Client) GetAccountState(accountAddr string) (AccountState, error) {
	accountAddrBytes, err := hex.DecodeString(accountAddr)
	if err != nil {
		return AccountState{}, err
	}
	requestedItems := []*types.RequestItem{
		{
			RequestedItems: &types.RequestItem_GetAccountStateRequest{
				GetAccountStateRequest: &types.GetAccountStateRequest{
					Address: accountAddrBytes,
				},
			},
		},
	}
	knownVersion := uint64(0)
	updateLedgerRequest := types.UpdateToLatestLedgerRequest{
		ClientKnownVersion: knownVersion,
		RequestedItems:     requestedItems,
	}
	updateLedgerResponse, err := c.acc.UpdateToLatestLedger(context.Background(), &updateLedgerRequest)
	if err != nil {
		return AccountState{}, err
	}

	accStateBlob := updateLedgerResponse.GetResponseItems()[0].GetGetAccountStateResponse().GetAccountStateWithProof().GetBlob().GetBlob()

	return decodeAccountStateBlob(accStateBlob)
}

// query transaction state
func (c Client) GetTransaction(startVersion uint64, limit uint64, fetchEvents bool) error {
	requestedItems := []*types.RequestItem{
		{
			RequestedItems: &types.RequestItem_GetTransactionsRequest{
				GetTransactionsRequest: &types.GetTransactionsRequest{
					StartVersion: startVersion,
					Limit:        limit,
					FetchEvents:  fetchEvents,
				},
			},
		},
	}
	knownVersion := uint64(0)
	updateLedgerRequest := types.UpdateToLatestLedgerRequest{
		ClientKnownVersion: knownVersion,
		RequestedItems:     requestedItems,
	}
	updateLedgerResponse, err := c.acc.UpdateToLatestLedger(context.Background(), &updateLedgerRequest)
	if err != nil {
		return err
	}
	_ = updateLedgerResponse.GetResponseItems()[0].GetGetTransactionsResponse().GetTxnListWithProof()
	return nil
}

// query transaction detail by account and sequence number
func (c Client) GetAccountTransactionBySequenceNumber(accountAddr string,
	sequenceNumber uint64, fetchEvents bool) (GoLibraTransactionInfo, error) {

	accountAddrBytes, err := hex.DecodeString(accountAddr)
	if err != nil {
		return GoLibraTransactionInfo{}, err
	}
	requestedItems := []*types.RequestItem{
		{
			RequestedItems: &types.RequestItem_GetAccountTransactionBySequenceNumberRequest{
				GetAccountTransactionBySequenceNumberRequest: &types.GetAccountTransactionBySequenceNumberRequest{
					Account:        accountAddrBytes,
					SequenceNumber: sequenceNumber,
					FetchEvents:    fetchEvents,
				},
			},
		},
	}

	knownVersion := uint64(0)
	updateLedgerRequest := types.UpdateToLatestLedgerRequest{
		ClientKnownVersion: knownVersion,
		RequestedItems:     requestedItems,
	}

	updateLedgerResponse, err := c.acc.UpdateToLatestLedger(context.Background(), &updateLedgerRequest)
	if err != nil {
		return GoLibraTransactionInfo{}, err
	}

	/// TransactionWithProof
	m := updateLedgerResponse.GetResponseItems()[0].GetGetAccountTransactionBySequenceNumberResponse().GetTransactionWithProof()
	return decodeTransaction(m)
	/// AccountStateWithProof
	/// accStateBlob := updateLedgerResponse.GetResponseItems()[0].GetGetAccountTransactionBySequenceNumberResponse().GetProofOfCurrentSequenceNumber().GetBlob().GetBlob()

	///return DecodeAccountStateBlob(accStateBlob)

}
