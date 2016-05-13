package data

import (
	"encoding/json"
	"strings"
	"fmt"
	"github.com/appscode/api/data/files"
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


