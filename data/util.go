package data

import (
	"encoding/json"
	"fmt"
	"github.com/appscode/api/data/files"
	"strings"
)

func DBGenericName(dbName, sku string) (string, error) {
	bytes, err := files.Asset("data/files/db_products.json")
	if err != nil {
		return "", err
	}

	dbs := make(map[string]DBProduct)
	err = json.Unmarshal(bytes, &dbs)
	if err != nil {
		return "", err
	}
	if pg, ok := dbs[dbName]; ok {
		for _, dt := range pg.DbTypes {
			if dt.Sku == strings.ToUpper(sku) {
				return dt.Name, nil
			}
		}
	}
	return "", fmt.Errorf("Can't detect generic name for db %v and sku %v", dbName, sku)
}

/*
curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer <token>" "https://api.digitalocean.com/v2/sizes"
*/
func BuildAgentExternalID(sku string) (string, error) {
	bytes, err := files.Asset("data/files/ci_products.json")
	if err != nil {
		return "", err
	}

	var ci CIProduct
	err = json.Unmarshal(bytes, &ci)
	if err != nil {
		return "", err
	}
	for _, agent := range ci.BuildAgents {
		if agent.Sku == strings.ToUpper(sku) {
			return agent.Details.ExternalID, nil
		}
	}
	return "", fmt.Errorf("Can't detect external id for CIProduct sku %v", sku)
}
