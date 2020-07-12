package json

import (
	"context"
	"io"

	"github.com/walkline/shippingpg/clientapi"
	"github.com/walkline/shippingpg/clientapi/scanner/json"
)

const scannerBufferSize = 100

type ReaderImporter struct {
	repo clientapi.KeyValueRepository
}

func NewReaderImporter(r clientapi.KeyValueRepository) *ReaderImporter {
	return &ReaderImporter{
		repo: r,
	}
}

func (i ReaderImporter) Import(ctx context.Context, r io.Reader) error {
	scanner := json.NewKeyValueScanner(r, scannerBufferSize)
	for scanner.Scan() {
		err := i.repo.Store(ctx, &clientapi.KeyValue{
			Key:   scanner.Key(),
			Value: scanner.Value(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
