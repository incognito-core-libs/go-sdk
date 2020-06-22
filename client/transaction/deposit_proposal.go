package transaction

import (
	ctypes "github.com/incognito-core-libs/go-sdk/common/types"
	"github.com/incognito-core-libs/go-sdk/types"
	"github.com/incognito-core-libs/go-sdk/types/msg"
	"github.com/incognito-core-libs/go-sdk/types/tx"
)

type DepositProposalResult struct {
	tx.TxCommitResult
}

func (c *client) DepositProposal(proposalID int64, amount int64, sync bool, options ...Option) (*DepositProposalResult, error) {
	fromAddr := c.keyManager.GetAddr()
	coins := ctypes.Coins{ctypes.Coin{Denom: types.NativeSymbol, Amount: amount}}
	depositMsg := msg.NewDepositMsg(fromAddr, proposalID, coins)
	commit, err := c.broadcastMsg(depositMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &DepositProposalResult{*commit}, err

}
