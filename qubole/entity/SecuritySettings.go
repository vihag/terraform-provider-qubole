package entity

import (
	_ "fmt"
)

type SecuritySettings struct {
	Encrypted_ephemerals      bool		`json:"encrypted_ephemerals,omitempty"`
	Customer_ssh_key          string	`json:"customer_ssh_key,omitempty"`
	Persistent_security_group string	`json:"persistent_security_group,omitempty"`
}
