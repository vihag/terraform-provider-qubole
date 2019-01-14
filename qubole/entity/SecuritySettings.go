package entity

import (
	"fmt"
)

type SecuritySettings struct {
	encrypted_ephemerals      bool
	customer_ssh_key          string
	persistent_security_group string
}
