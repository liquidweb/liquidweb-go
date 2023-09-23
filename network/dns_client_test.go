package network

import (
	"flag"
	"os"
	"testing"

	"github.com/liquidweb/liquidweb-go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var liveTest = flag.Bool("live", false, "set true to run live test")

func init() {
	flag.Parse()
}

func TestLivePresent(t *testing.T) {
	if !*liveTest {
		t.Skip("skipping live test")
	}

	username, password := os.Getenv("LWAPI_USERNAME"), os.Getenv("LWAPI_PASSWORD")
	url := "https://api.liquidweb.com"
	timeout := 60

	api, err := client.NewAPI(username, password, url, timeout)
	require.NoError(t, err, "failed to create API")

	reqParams := DNSRecordParams{
		Zone: os.Getenv("LWAPI_ZONE"),
	}

	dnsRecords, err := api.NetworkDNS.List(&reqParams)
	require.NoError(t, err)
	assert.NotZero(t, len(dnsRecords.Items))
}
