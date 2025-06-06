package api

import (
	"context"

	okex "github.com/pefish/go-okx"
	"github.com/pefish/go-okx/api/rest"
	"github.com/pefish/go-okx/api/ws"
)

// Client is the main api wrapper of okex
type Client struct {
	Rest *rest.ClientRest
	Ws   *ws.ClientWs
	ctx  context.Context
}

// NewClient returns a pointer to a fresh Client
func NewClient(
	ctx context.Context,
	apiKey,
	secretKey,
	passphrase string,
	destination okex.Destination,
) (*Client, error) {
	restURL := okex.RestURL
	wsPubURL := okex.PublicWsURL
	wsPriURL := okex.PrivateWsURL
	switch destination {
	case okex.AwsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.AwsPublicWsURL
		wsPriURL = okex.AwsPrivateWsURL
	case okex.DemoServer:
		restURL = okex.DemoRestURL
		wsPubURL = okex.DemoPublicWsURL
		wsPriURL = okex.DemoPrivateWsURL
	case okex.CandleWsServer:
		restURL = okex.AwsRestURL
		wsPubURL = okex.HandleWsURL
		wsPriURL = okex.AwsPrivateWsURL
	}

	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, destination)
	c := ws.NewClient(ctx, apiKey, secretKey, passphrase, map[bool]okex.BaseURL{true: wsPriURL, false: wsPubURL})

	return &Client{
		Rest: r,
		Ws:   c,
		ctx:  ctx,
	}, nil
}
