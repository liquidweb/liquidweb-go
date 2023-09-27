package network

import (
	"flag"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/liquidweb/liquidweb-go/client"
	"github.com/liquidweb/liquidweb-go/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var liveTest = flag.Bool("live", false, "set true to run live test")
var testZone = flag.String("zone", "", "given zone for dns tests")

func TestLiveListRecords(t *testing.T) {
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
	records := []string{}
	for _, each := range dnsRecords.Items {
		records = append(records, (each.Name + " => " + each.RData))
	}
	log.Printf("all zones on zone %s:\n%s\n",
		*testZone, strings.Join(records, "\n"))
}

func TestLiveZones(t *testing.T) {
	if !*liveTest {
		t.Skip("skipping live test")
	}

	username, password := os.Getenv("LWAPI_USERNAME"), os.Getenv("LWAPI_PASSWORD")
	url := "https://api.liquidweb.com"
	timeout := 60

	api, err := client.NewAPI(username, password, url, timeout)
	require.NoError(t, err, "failed to create API")

	dnsZones, err := api.NetworkDNSZone.ListAll()
	require.NoError(t, err, "got an error  sting domains")
	assert.NotZero(t, len(dnsZones.Items))
	zones := []string{}
	for _, each := range dnsZones.Items {
		zones = append(zones, each.Name)
	}
	log.Printf("all zones on account:\n%s\n", strings.Join(zones, "\n"))
}

func TestLiveRecordsListAll(t *testing.T) {
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

	dnsRecords, err := api.NetworkDNS.ListAll(*testZone)
	require.NoError(t, err, "got an error listing records")
	assert.NotZero(t, len(dnsRecords.Items))
	records := []string{}
	for _, each := range dnsRecords.Items {
		records = append(records, (each.Name + " => " + each.RData))
	}
	log.Printf("all zones on zone %s:\n%s\n",
		*testZone, strings.Join(records, "\n"))
}
