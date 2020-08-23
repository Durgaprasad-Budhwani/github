package internal

import (
	"fmt"

	"github.com/pinpt/agent.next/sdk"
)

// AutoConfigure is called when a cloud integration has requested to be auto configured
func (g *GithubIntegration) AutoConfigure(autoconfig sdk.AutoConfigure) (*sdk.Config, error) {
	config := autoconfig.Config()
	logger := g.logger
	_, client, err := g.newGraphClient(logger, config)
	if err != nil {
		return nil, fmt.Errorf("error creating graphql client: %w", err)
	}
	viewer, err := g.fetchViewer(logger, client, autoconfig)
	if err != nil {
		return nil, err
	}
	var acct sdk.ConfigAccount
	acct.ID = viewer
	acct.Public = false
	acct.Type = sdk.ConfigAccountTypeUser
	accounts := make(sdk.ConfigAccounts)
	accounts[acct.ID] = &acct
	config.Accounts = &accounts
	return &config, nil
}
