package data

import (
	"fmt"
	"testing"
	"log"
)

func Test(t *testing.T) {
	agent, err := CIBuildAgent("CI-C1R1-DO")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(agent.Metadata.ExternalID)
}
