package beefy

import (
	"os"
	"testing"

	"github.com/kirill-sr/go-substrate-rpc-client/v4/client"
	"github.com/kirill-sr/go-substrate-rpc-client/v4/config"
)

var testBeefy Beefy

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testBeefy = NewBeefy(cl)
	os.Exit(m.Run())
}
