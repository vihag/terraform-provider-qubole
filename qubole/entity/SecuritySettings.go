package entity

import (
	_ "fmt"
)

type SecuritySettings struct {
	Encrypted_ephemerals      bool
	Customer_ssh_key          string
	Persistent_security_group string
}
