package mongoKeeper

import (
	"context"
	"fmt"

	"github.com/12153/e-commerce/pkg/common/recordKeeper"
	"github.com/12153/e-commerce/pkg/common/search"
	"github.com/12153/e-commerce/pkg/common/search/query"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	mongoBSON "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/trace"
)

// var _ recordKeeper.RecordKeeper = &RecordKeeperMongo{}

type RecordKeeperMongo struct {
	tracer            trace.Tracer
	collection        mongo.Collection
	historyCollection mongo.Collection
}

// CreateEntity implements recordKeeper.RecordKeeper
func (r *RecordKeeperMongo) CreateEntity(ctx context.Context, entity interface{}) (recordKeeper.RollBackFunc, error) {
	ctx, span := r.tracer.Start(
		ctx,
		r.collection.Name()+"Recordkeeper.CreateEntity",
	)
	defer span.End()

	result, err := r.collection.InsertOne(ctx, entity)
	// return if successful
	rollBack := func() error {
		if _, rollBackErr := r.collection.DeleteOne(ctx, map[string]interface{}{
			"_id": result.InsertedID,
		}); rollBackErr != nil {
			return fmt.Errorf("error rolling back: %s", rollBackErr)
		}
		return nil
	}
	if err == nil {
		return rollBack, nil
	}

	switch v := err.(type) {
	case mongo.WriteException:
		if len(v.WriteErrors) == 1 && v.WriteErrors[0].Code == 11000 {
			return rollBack, fmt.Errorf("write error")
		}

		// otherwise some other error occurred
		log.Ctx(ctx).Error().Err(errors.WithStack(err)).Msg("error inserting entity")
		return rollBack, fmt.Errorf("error inserting entity: %w", err)

	default:
		log.Ctx(ctx).Error().Err(errors.WithStack(err)).Msg("error inserting entity")
		return rollBack, fmt.Errorf("error inserting entity: %w", err)
	}
}

// DeleteEntity implements recordKeeper.RecordKeeper
func (*RecordKeeperMongo) DeleteEntity(ctx context.Context, entity interface{}) (func() error, error) {
	panic("unimplemented")
}

// RetrieveEntity implements recordKeeper.RecordKeeper
func (*RecordKeeperMongo) RetrieveEntity(ctx context.Context, entity interface{}) error {
	panic("unimplemented")
}

// SearchEntities implements recordKeeper.RecordKeeper
func (*RecordKeeperMongo) SearchEntities(ctx context.Context, entity interface{}) error {
	panic("unimplemented")
}

func QueryToMongoFindOptions(q query.Query) (*options.FindOptions, error) {
	// generate sorting
	sorting := mongoBSON.D{}
	for i := range q.Sorting {
		if err := q.Sorting[i].IsValid(); err != nil {
			return nil, err
		}
		if q.Sorting[i].SortOrder == query.SortOrderDescending {
			sorting = append(
				sorting,
				mongoBSON.E{
					Key:   q.Sorting[i].Field,
					Value: -1,
				},
			)
		} else {
			// assumed to be ascending since the sorting entry is valid
			sorting = append(
				sorting,
				mongoBSON.E{
					Key:   q.Sorting[i].Field,
					Value: 1,
				},
			)
		}
	}

	// create find options
	findOptions := new(options.FindOptions)

	// populate find options
	findOptions.SetSort(sorting)
	findOptions.SetSkip(q.Offset)
	if q.Limit > 0 {
		findOptions.SetLimit(q.Limit)
	}

	return findOptions, nil
}

func (r *RecordKeeperMongo) FindEntities(ctx context.Context, entities interface{}, identifier search.Identifier, query query.Query) (int64, error) {
	ctx, span := r.tracer.Start(
		ctx,
		r.collection.Name()+"Recordkeeper.FindEntities",
	)
	defer span.End()

	if err := identifier.IsValid(); err != nil {
		return 0, err
	}

	// get filter and options
	filter := identifier.ToFilter()
	findOptions, err := QueryToMongoFindOptions(query)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error converting query to find options")
		return 0, fmt.Errorf("error converting query to find options: %w", err)
	}
	// perform find
	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("error creating cursor for '%s' collection", r.collection.Name())
		return 0, fmt.Errorf("error creating cursor for '%s' collection: %w", r.collection.Name(), err)
	}

	// decode the results
	if err := cursor.All(ctx, entities); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("error decoding documents")
		return 0, fmt.Errorf("error decoding documents: %w", err)
	}

	// get document count
	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msgf("error counting documents in '%s' collection", r.collection.Name())
		return 0, fmt.Errorf("error counting documents in '%s' collection: %w", r.collection.Name(), err)
	}

	return count, nil
}
