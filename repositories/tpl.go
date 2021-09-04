package repositories

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
	// DATABASE_LIB 锚点请勿删除! Do not delete this line!
)

// REPO_NAMERepositoryProvider
var REPO_NAMERepositoryProvider = wire.NewSet(NewREPO_NAMERepository, wire.Bind(new(IREPO_NAMERepository), new(*REPO_NAMERepository)))

// IREPO_NAMERepository
type IREPO_NAMERepository interface {
	GetUser(ctx context.Context) (*REPO_NAMERepository, error)
}

// REPO_NAMERepository
type REPO_NAMERepository struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`

	logger *zap.Logger
	// STRUCT_ATTR 锚点请勿删除! Do not delete this line!
}

// NewREPO_NAMERepository
func NewREPO_NAMERepository(
	logger *zap.Logger,
	// FUNC_PARAM 锚点请勿删除! Do not delete this line!
) *REPO_NAMERepository {
	return &REPO_NAMERepository{
		logger: logger.With(zap.String("type", "REPO_NAMERepository")),
		// FUNC_ATTR 锚点请勿删除! Do not delete this line!
	}
}

func (p *REPO_NAMERepository) GetUser(ctx context.Context) (*REPO_NAMERepository, error) {
	r := &REPO_NAMERepository{}
	// p.mysql.Find(ctx, user, "user", nil)
	// p.mongo.Database.Collection("user").Find()
	// p.redis.DB().Get(ctx, "user").Result()
	return r, nil
}