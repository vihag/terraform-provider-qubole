package model

import (
	"encoding/json"
	_ "fmt"
	"log"
	_ "strconv"
)

type Datadisk struct {
	Count            int             `json:"count,omitempty"`
	Disktype         string          `json:"type,omitempty"`
	Encryption       bool            `json:"encryption,omitempty"`
	Size             int             `json:"size,omitempty"`
	Upscaling_config UpscalingConfig `json:"upscaling_config,omitempty"`
}

//Custom unmarshalling logic
func (u *Datadisk) UnmarshalJSON(data []byte) error {
	log.Printf("[DEBUG]using custom unmarshaller for umarshalling cluster object: %s", "Datadisk")
	type Alias Datadisk //some alias for the actual struct
	aux := &struct {
		Size             []interface{}   `json:"size,omitempty"`
		Upscaling_config UpscalingConfig `json:"ebs_upscaling_config,omitempty"`
		*Alias                           //rest of the actual stuff
	}{
		Alias: (*Alias)(u), //business as usual
	}
	if err := json.Unmarshal(data, &aux); err != nil { //unmarshall the rest
		//cannot unmarshal number into Go struct field .size of type string
		log.Printf("[DEBUG]error in umarshalling aux struct for Datadisk: %s", err.Error())
		var result map[string]interface{}
		json.Unmarshal(data, &result)
		return err
	}

	//Now concentrate on Size

	if len(aux.Size) > 0 {
		log.Printf("[DEBUG]reading the size array to extract the first element: %s", int(aux.Size[0].(float64)))
		u.Size = int(aux.Size[0].(float64))
	}

	//Now concentrate on Ebs Upscaling Config
	if &aux.Upscaling_config != nil {
		log.Printf("[DEBUG]Translating Ebs_Upscaling_Config to Upscaling config %s", aux.Upscaling_config)
		u.Upscaling_config = aux.Upscaling_config
	}

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

	if &ia.Upscaling_config != nil {
		attrs["upscaling_config"] = FlattenUpscalingConfig(&ia.Upscaling_config)
	}

	result = append(result, attrs)

	return result
}

func ReadDatadiskFromTf(datadisk *Datadisk, datadiskConfig []interface{}) bool {

	if len(datadiskConfig) > 0 {
		configs := datadiskConfig[0].(map[string]interface{})
		if v, ok := configs["count"]; ok {
			datadisk.Count = v.(int)
		}
		if v, ok := configs["type"]; ok {
			datadisk.Disktype = v.(string)
		}
		if v, ok := configs["encryption"]; ok {
			datadisk.Encryption = v.(bool)
		}
		if v, ok := configs["size"]; ok {
			datadisk.Size = v.(int)
		}
		//Read disk upscaling config
		var upscaling_config UpscalingConfig
		if v, ok := configs["upscaling_config"]; ok {
			ebsUpscalingConfigs := v.([]interface{})
			ReadUpscalingConfigFromTf(&upscaling_config, ebsUpscalingConfigs)
			datadisk.Upscaling_config = upscaling_config
		}
	}

	return true
}
