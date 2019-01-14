package qubole

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func validatePrestoVersion(validators ...schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return func(i interface{}, k string) ([]string, []error) {
		var allErrors []error
		var allWarnings []string
		for _, validator := range validators {
			validatorWarnings, validatorErrors := validator(i, k)
			if len(validatorWarnings) == 0 && len(validatorErrors) == 0 {
				return []string{}, []error{}
			}
			allWarnings = append(allWarnings, validatorWarnings...)
			allErrors = append(allErrors, validatorErrors...)
		}
		return allWarnings, allErrors
	}
}

func validatePrestoVersions() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{
		"0.157",
		"0.180",
		"0.193",
		"st1",
	}, false)
}
