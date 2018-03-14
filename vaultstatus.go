package main

import (
	"encoding/json"
	"fmt"
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/metric"
	"github.com/newrelic/infra-integrations-sdk/sdk"
	"io/ioutil"
	"net/http"
	"strconv"
	"os"
)

type VaultAPIResponse struct {
	Initialized   bool   `json:"initialized"`
	Sealed        bool   `json:"sealed"`
	Standby       bool   `json:"standby"`
	ServerTimeUTC int    `json:"server_time_utc"`
	Version       string `json:"version"`
	ClusterName   string `json:"cluster_name"`
	ClusterID     string `json:"cluster_id"`
}

type argumentList struct {
	sdkArgs.DefaultArgumentList
}

const (
	integrationName    = "com.org.vaultstatus"
	integrationVersion = "0.1.0"
)

var args argumentList

func main() {
	integration, err := sdk.NewIntegration(integrationName, integrationVersion, &args)
	fatalIfErr(err)

	if args.All || args.Inventory {
		fatalIfErr(populateInventory(integration.Inventory))
	}

	if args.All || args.Metrics {
		ms := integration.NewMetricSet("VaultStatus")
		fatalIfErr(populateMetrics(ms))
	}
	fatalIfErr(integration.Publish())
}

func populateInventory(inventory sdk.Inventory) error {
	// Insert here the logic of your integration to get the inventory data
	// Ex: inventory.SetItem("softwareVersion", "value", "1.0.1")
	// -
	vaultUrl := os.Getenv("VAULT_URL")
	res, err := http.Get(vaultUrl)
        if err != nil {
                panic(err.Error())
        }

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
                panic(err.Error())
        }

	s, err := getVaultResponse([]byte(body))

	inventory.SetItem("cluster", "clustername", s.ClusterName)
	inventory.SetItem("cluster", "clusterid", s.ClusterID)
	fatalIfErr(err)
	return err
}

func populateMetrics(ms *metric.MetricSet) error {
	// Insert here the logic of your integration to get the metrics data
	// Ex: ms.SetMetric("requestsPerSecond", 10, metric.GAUGE)
	// --
	vaultUrl := os.Getenv("VAULT_URL")
	res, err := http.Get(vaultUrl)
        if err != nil {
                panic(err.Error())
        }

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
                panic(err.Error())
        }
	s, err := getVaultResponse([]byte(body))
	seal := strconv.FormatBool(s.Sealed)

	ms.SetMetric("Sealed", seal, metric.ATTRIBUTE)

	return err
}

func getVaultResponse(body []byte) (*VaultAPIResponse, error) {
	var s = new(VaultAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
