package config

import (
	"encoding/base64"
	"strings"
)

const (
	ReasonRequired = "required"
)

// Enver is a function that accepts a key and returns a string value associated
// with that key.
type Enver func(key string) string

func RequiredString(enver Enver, key string) (string, error) {
	s := String(enver, key)
	if len(s) == 0 {
		return "", NewEnverError(key, "string", ReasonRequired)
	}
	return s, nil
}

func String(enver Enver, key string) string {
	return enver(key)
}

func RequiredBase64Bytes(enver Enver, key string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(enver(key))
	if err != nil {
		return nil, NewEnverError(key, "base64", err.Error())
	}
	if len(b) == 0 {
		return nil, NewEnverError(key, "base64", ReasonRequired)
	}
	return b, nil
}

func StringSlice(enver Enver, key string) []string {
	return strings.Split(enver(key), ",")
}

type EnverError struct {
	Key    string
	Type   string
	Reason string
}

func NewEnverError(key, type_, reason string) *EnverError {
	return &EnverError{
		Key:    key,
		Type:   type_,
		Reason: reason,
	}
}

func (e *EnverError) Error() string {
	return ""
}
