package types

import (
	ntypes "github.com/incognito-core-libs/go-sdk/common/types"
	"github.com/incognito-core-libs/go-sdk/types/tx"
	"github.com/incognito-core-libs/go-amino"
	types "github.com/incognito-core-libs/tendermint/rpc/core/types"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}
