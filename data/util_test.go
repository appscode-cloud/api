package data

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	agent, _ := CIBuildAgent("CI-C1R1-DO")
	fmt.Println(agent.Metadata.ExternalID)
}
