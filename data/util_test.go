package data

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	agent, _ := CIBuildAgent("CI-R1")
	fmt.Println(agent.Details.ExternalID)
}
