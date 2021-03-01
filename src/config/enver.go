package config

// Enver is a function that accepts a key and returns a string value associated
// with that key.
type Enver func(key string) string

func RequiredString(enver Enver, key string) (string, error) {
	return "", nil
}

func String(enver Enver, key string) string {
	return enver(key)
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
