package data

import (
	"encoding/json"
	"fmt"
	"github.com/appscode/api/data/files"
	"log"
	"testing"
)

func Test(t *testing.T) {
	agent, err := CIBuildAgent("CI-C1R1-DO")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(agent.UnitPriceUSD.String())
	fmt.Println(agent.UnitPriceUSD.Float())
}

func TestGenericParsing(t *testing.T) {
	bytes, err := files.Asset("data/files/subscription.latest.json")
	if err != nil {
		log.Fatal(err)
	}

	p := &struct {
		Subscription []*Product `json:"subscription"`
	}{}
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Subscription[0].DisplayName)
	fmt.Println(len(p.Subscription[0].DisplayPriceUSD))
	fmt.Println(string(p.Subscription[0].Metadata))
}
