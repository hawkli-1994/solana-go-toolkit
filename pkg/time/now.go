package datetime

import (
	"context"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
)

func Now(client *rpc.Client) (*time.Time, error) {
	blockInfo, err := client.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}
	slot := blockInfo.Context.Slot
	zero := uint64(0)
	opts := &rpc.GetBlockOpts{
		MaxSupportedTransactionVersion: &zero,
	}
	blockResult, err := client.GetBlockWithOpts(context.Background(), slot, opts)
	if err != nil {
		return nil, err
	}
	blockTime := blockResult.BlockTime.Time()
	return &blockTime, nil
}
