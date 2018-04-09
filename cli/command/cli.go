package command

import (
	"net/http"
	"github.com/stellar/go/clients/horizon"
)

type BraveCli struct {
	Network struct {
		Passphrase string
		Horizon string
	}

	Account map[string]string
}

func (cli *BraveCli) HorizonClient() *horizon.Client {
	return &horizon.Client{
		URL:  cli.Network.Horizon,
		HTTP: http.DefaultClient,
	}
}
