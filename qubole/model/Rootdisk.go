package model

import (
	_ "fmt"
)

type Rootdisk struct {
	Size int `json:"size,omitempty"`
}

/*
function to flatten Rootdisk
*/
func FlattenRootdisk(ia *Rootdisk) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Size != nil {
		attrs["size"] = ia.Size
	}

	result = append(result, attrs)

	return result
}

func ReadRootdiskFromTf(rootdisk *Rootdisk, rootdiskConfig []interface{}) bool {

	if len(rootdiskConfig) > 0 {
		configs := rootdiskConfig[0].(map[string]interface{})

		if v, ok := configs["size"]; ok {
			rootdisk.Size = v.(int)
		}
	}

	return true
}
