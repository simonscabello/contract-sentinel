package results

import (
	"context"
	"time"

	"github.com/simonscabello/contract-sentinel/internal/contracts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Saver interface {
	Save(ctx context.Context, input contracts.Contract, result contracts.ValidationResult) error
}

type SaverWithQuery interface {
	Saver
	FindAll(ctx context.Context, query QueryParams) ([]Result, error)
}

type Result struct {
	ContractPath string    `bson:"contract_path"`
	ProviderURL  string    `bson:"provider_url"`
	Consumer     string    `bson:"consumer"`
	Provider     string    `bson:"provider"`
	Version      string    `bson:"version"`
	Success      bool      `bson:"success"`
	Output       string    `bson:"output"`
	ExecutedAt   time.Time `bson:"executed_at"`
}

type QueryParams struct {
	Consumer string
	Provider string
	Success  *bool
}

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("results"),
	}
}

func (r *Repository) Save(ctx context.Context, input contracts.Contract, result contracts.ValidationResult) error {
	doc := Result{
		ContractPath: input.Path,
		ProviderURL:  input.ProviderURL,
		Consumer:     input.Consumer,
		Provider:     input.Provider,
		Version:      input.Version,
		Success:      result.Success,
		Output:       result.Output,
		ExecutedAt:   time.Now(),
	}

	_, err := r.collection.InsertOne(ctx, doc)
	return err
}

func (r *Repository) FindAll(ctx context.Context, query QueryParams) ([]Result, error) {
	filter := bson.M{}

	if query.Consumer != "" {
		filter["consumer"] = query.Consumer
	}
	if query.Provider != "" {
		filter["provider"] = query.Provider
	}
	if query.Success != nil {
		filter["success"] = *query.Success
	}

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []Result
	if err := cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

var _ Saver = (*Repository)(nil)
var _ SaverWithQuery = (*Repository)(nil)
