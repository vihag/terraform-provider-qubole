package model

import (
	"encoding/json"
	"log"
	_ "strconv"
	_ "fmt"
)

type Datadisk struct {
	Count                int                `json:"count,omitempty"`
	Disktype             string             `json:"type,omitempty"`
	Encryption           bool               `json:"encryption,omitempty"`
	Size                 int                `json:"size,omitempty"`
	Ebs_upscaling_config EbsUpscalingConfig `json:"ebs_upscaling_config,omitempty"`
}

//Custom unmarshalling logic
func (u *Datadisk) UnmarshalJSON(data []byte) error {
	log.Printf("[ERR]using custom unmarshaller for umarshalling cluster object: %s", "Datadisk")
	type Alias Datadisk //some alias for the actual struct
	aux := &struct {
		Size   []interface{} `json:"size,omitempty"` // aux struct which says just leave the custom field alone
		*Alias               //rest of the actual stuff
	}{
		Alias: (*Alias)(u), //business as usual
	}
	if err := json.Unmarshal(data, &aux); err != nil { //unmarshall the rest
		//cannot unmarshal number into Go struct field .size of type string
		log.Printf("[ERR]error in umarshalling aux struct for Datadisk: %s", err.Error())
		var result map[string]interface{}
		json.Unmarshal(data, &result)
		log.Print("Reading the aux struct as a map")
		log.Println(result)
		//result : map[size:[100 1 gp2] encryption:true count:1 type:gp2]
		return err
	}

	//Now concentrate on datadisk

	log.Printf("[DEBUG]reading the size array to extract the first element: %s", int(aux.Size[0].(float64)))
	u.Size = int(aux.Size[0].(float64))
	log.Printf("[INFO]Pretty Printing Unmarshalled Response %#v", u)
	return nil
}

/*
function to flatten Datadisk
*/
func FlattenDatadisk(ia *Datadisk) []map[string]interface{} {
	attrs := map[string]interface{}{}
	result := make([]map[string]interface{}, 0)

	if &ia.Count != nil {
		attrs["count"] = ia.Count
	}

	if &ia.Disktype != nil {
		attrs["type"] = ia.Disktype
	}

	if &ia.Encryption != nil {
		attrs["encryption"] = ia.Encryption
	}

	if &ia.Size != nil {
		attrs["size"] = ia.Size
	}

	if &ia.Ebs_upscaling_config != nil {
		attrs["ebs_upscaling_config"] = FlattenEbsUpscalingConfig(&ia.Ebs_upscaling_config)
	}

	result = append(result, attrs)

	return result
}

