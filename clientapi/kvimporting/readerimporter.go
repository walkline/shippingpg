package kvimporting

import (
	"context"
	"io"
)

type ReaderImporter interface {
	Import(context.Context, io.Reader) error
}
