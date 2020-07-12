package clientapi

import (
	"context"
	"errors"
)

// KeyValue is generic entity
type KeyValue struct {
	Key   string
	Value interface{}
}

var ErrUnkKey = errors.New("key not found")

type KeyValueRepository interface {
	Store(context.Context, *KeyValue) error
	FindByKey(context.Context, string) (*KeyValue, error)
}
