package internal

import (
	"strings"

	"github.com/pinpt/agent.next/sdk"
)

func isBot(name string) bool {
	return strings.HasSuffix(name, "[bot]") || strings.HasSuffix(name, "-bot") || strings.HasSuffix(name, " Bot") || name == "GitHub"
}

type author struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avatar string `json:"avatarUrl"`
	Login  string `json:"login"`
	URL    string `json:"url"`
	Type   string `json:"type"`
}

func (a author) ToModel(customerID string) *sdk.SourceCodeUser {
	user := &sdk.SourceCodeUser{}
	user.CustomerID = customerID
	user.RefID = a.RefID(customerID)
	user.RefType = refType
	user.ID = sdk.NewSourceCodeUserID(customerID, refType, user.RefID)
	user.URL = sdk.StringPointer(a.URL)
	user.AvatarURL = sdk.StringPointer(a.Avatar)
	user.Email = sdk.StringPointer(a.Email)
	if user.Email != nil {
		id := sdk.Hash(customerID, a.Email)
		if id != user.RefID {
			user.AssociatedRefID = sdk.StringPointer(id)
		}
	}
	user.Name = a.Name
	switch a.Type {
	case "Bot":
		user.Type = sdk.SourceCodeUserTypeBot
	case "User":
		user.Type = sdk.SourceCodeUserTypeHuman
		user.Username = sdk.StringPointer(a.Login)
	case "Mannequin":
	}
	if user.RefID == "" || isBot(a.Name) {
		user.Type = sdk.SourceCodeUserTypeBot
	}
	return user
}

func (a author) RefID(customerID string) string {
	// FIXME: review how we do this in current agent to match
	switch a.Type {
	case "Bot":
		return ""
	case "User":
		return a.ID
	case "Mannequin":
	}
	if a.Email != "" {
		return sdk.Hash(customerID, a.Email)
	}
	return "" // FIXME
}

type gitUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatarUrl"`
	User   author `json:"user"`
}

func (a gitUser) RefID(customerID string) string {
	// FIXME
	if a.User.ID != "" {
		return a.User.ID
	}
	if a.Email != "" {
		return sdk.Hash(customerID, a.Email)
	}
	return ""
}

func (a gitUser) ToModel(customerID string, integrationID string) *sdk.SourceCodeUser {
	user := &sdk.SourceCodeUser{}
	user.CustomerID = customerID
	user.RefID = a.RefID(customerID)
	user.RefType = refType
	if a.Email != "" {
		id := sdk.Hash(customerID, a.Email)
		if id != user.RefID {
			user.AssociatedRefID = sdk.StringPointer(id)
		}
	}
	user.IntegrationInstanceID = sdk.StringPointer(integrationID)
	user.URL = sdk.StringPointer(a.User.URL)
	user.AvatarURL = sdk.StringPointer(a.Avatar)
	user.Email = sdk.StringPointer(a.Email)
	user.Name = a.Name
	switch a.User.Type {
	case "Bot":
		user.Type = sdk.SourceCodeUserTypeBot
		user.Username = sdk.StringPointer(a.User.Login)
	case "User":
		user.Type = sdk.SourceCodeUserTypeHuman
		user.Username = sdk.StringPointer(a.User.Login)
	case "Mannequin":
	}
	if user.RefID == "" || isBot(a.Name) {
		user.Type = sdk.SourceCodeUserTypeBot
	}
	return user
}
