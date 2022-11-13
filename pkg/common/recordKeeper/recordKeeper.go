package recordKeeper

import (
	"context"
)

type RecordKeeper interface {
	CreateEntity(ctx context.Context, entity interface{}) (RollBackFunc, error)
	DeleteEntity(ctx context.Context, entity interface{}) (RollBackFunc, error)
	RetrieveEntity(ctx context.Context, entity interface{}) error
	SearchEntities(ctx context.Context, entity interface{}) error
}

type RollBackFunc = func() error
