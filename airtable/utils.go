package airtable

import (
	"context"
	"errors"

	"github.com/mehanizm/airtable"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*airtable.Client, error) {
	airtableConfig := GetConfig(d.Connection)

	if &airtableConfig == nil {
		return nil, errors.New("You must have an airtable config file")
	}

	token := ""
	databaseid := ""

	if airtableConfig.Token != nil {
		token = *airtableConfig.Token
	}
	if airtableConfig.DatabaseID != nil {
		databaseid = *airtableConfig.DatabaseID
	}
	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if databaseid == "" {
		return nil, errors.New("'databaseid' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	return airtable.NewClient(token), nil
}
