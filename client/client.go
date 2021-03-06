package client

import (
	"gopkg.in/resty.v1"

	"github.com/incognito-core-libs/go-sdk/client/basic"
	"github.com/incognito-core-libs/go-sdk/client/query"
	"github.com/incognito-core-libs/go-sdk/client/transaction"
	"github.com/incognito-core-libs/go-sdk/client/websocket"
	"github.com/incognito-core-libs/go-sdk/common/types"
	"github.com/incognito-core-libs/go-sdk/keys"
)

// dexClient wrapper
type dexClient struct {
	query.QueryClient
	websocket.WSClient
	transaction.TransactionClient
	basic.BasicClient
}

// DexClient methods
type DexClient interface {
	basic.BasicClient
	query.QueryClient
	websocket.WSClient
	transaction.TransactionClient
}

func init() {
	resty.DefaultClient.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
}

func NewDexClient(baseUrl string, network types.ChainNetwork, keyManager keys.KeyManager) (DexClient, error) {
	types.Network = network
	c := basic.NewClient(baseUrl)
	w := websocket.NewClient(c)
	q := query.NewClient(c)
	n, err := q.GetNodeInfo()
	if err != nil {
		return nil, err
	}
	t := transaction.NewClient(n.NodeInfo.Network, keyManager, q, c)
	return &dexClient{BasicClient: c, QueryClient: q, TransactionClient: t, WSClient: w}, nil
}
