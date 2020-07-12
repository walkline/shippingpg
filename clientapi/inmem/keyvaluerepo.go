package inmem

import (
	"context"

	"github.com/walkline/shippingpg/clientapi"
)

type KeyValueRepo struct {
	// WARNING: not thread safe
	s map[string]clientapi.KeyValue
}

func NewKeyValueRepo() *KeyValueRepo {
	return &KeyValueRepo{
		s: map[string]clientapi.KeyValue{},
	}
}

func (r KeyValueRepo) Store(ctx context.Context, kv *clientapi.KeyValue) error {
	r.s[kv.Key] = *kv
	return nil
}

func (r KeyValueRepo) FindByKey(ctx context.Context, key string) (*clientapi.KeyValue, error) {
	v, found := r.s[key]
	if !found {
		return nil, clientapi.ErrUnkKey
	}

	return &v, nil
}
