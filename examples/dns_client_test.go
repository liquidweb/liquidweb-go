package network

import (
	"flag"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/liquidweb/liquidweb-go/client"
	"github.com/liquidweb/liquidweb-go/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var liveTest = flag.Bool("live", false, "set true to run live test")
var testZone = flag.String("zone", "", "given zone for dns tests")

func TestLivePresent(t *testing.T) {
	if !*liveTest {
		t.Skip("skipping live test")
	}
	if strings.TrimSpace(*testZone) == "" {
		t.Skip("skipping zone test without zone")
	}

	username, password := os.Getenv("LWAPI_USERNAME"), os.Getenv("LWAPI_PASSWORD")
	url := "https://api.liquidweb.com"
	timeout := 60

	api, err := client.NewAPI(username, password, url, timeout)
	require.NoError(t, err, "failed to create API")

	reqParams := network.DNSRecordParams{
		Zone: *testZone,
	}

	dnsRecords, err := api.NetworkDNS.List(&reqParams)
	require.NoError(t, err, "got an error using parameters %#v", reqParams)
	assert.NotZero(t, len(dnsRecords.Items))
	log.Printf("got back %s", spew.Sdump(dnsRecords))
}
