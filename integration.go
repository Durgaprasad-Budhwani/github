package main

import (
	"github.com/pinpt/github/internal"
	"github.com/pinpt/agent/runner"
)

// Integration is used to export the integration
var Integration internal.GithubIntegration

func main() {
	runner.Main(&Integration)
}
