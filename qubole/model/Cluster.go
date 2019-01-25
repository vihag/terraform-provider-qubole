package model

import (
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "strconv"
	_ "encoding/gob"
    _ "bytes"
)

type Cluster struct {
	Id            int          `json:"id,omitempty"`
	State         string       `json:"state,omitempty"`
	Cloud_config  CloudConfig  `json:"cloud_config,omitempty"`
	Cluster_info  ClusterInfo  `json:"cluster_info,omitempty"`
	Engine_config EngineConfig `json:"engine_config,omitempty"`
	Monitoring    Monitoring   `json:"monitoring,omitempty"`
	Internal      Internal     `json:"internal,omitempty"`
}

//Custom unmarshalling logic
/*func (u *Cluster) UnmarshalJSON(data []byte) error {
	log.Printf("[ERR]using custom unmarshaller for umarshalling cluster object: %s", "Cluster")
	type Alias Cluster //some alias for the actual struct
	aux := &struct {
		Cluster_info ClusterInfo `json:"cluster_info,omitempty"` // aux struct which says just leave the custom field alone
		*Alias                   //rest of the actual stuff
	}{
		Alias: (*Alias)(u), //business as usual
	}
	if err := json.Unmarshal(data, &aux); err != nil { //unmarshall the rest
		return err
	}

	//Now concentrate on cluster_info
	var cluster_info *ClusterInfo
	err := json.Unmarshal(GetBytes(aux.Cluster_info), &cluster_info)
	if err != nil {
		log.Printf("[ERR]There was an error unmarshalling cluster_info: %s", err.Error())
		return fmt.Errorf("There was an error unmarshalling cluster_info %s", err.Error())
	}
	u.Cluster_info = *cluster_info
	return nil
}

func GetBytes(key interface{}) ([]byte) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil
    }
    return buf.Bytes()
}*/
