package COMPONENT_TYPE

import (
	"context"

	"MODULE_NAME/internal/config"

	"github.com/google/wire"
	"github.com/goriller/ginny/logger"
	"go.uber.org/zap"
)

// REPO_NAMERepositoryProvider
var COMPONENT_NAMECOMPONENT_UP_TYPEProviderSet = wire.NewSet(
	NewCOMPONENT_NAMECOMPONENT_UP_TYPE,
	wire.Bind(new(ICOMPONENT_NAMECOMPONENT_UP_TYPE),
		new(*COMPONENT_NAMECOMPONENT_UP_TYPE)))

// ICOMPONENT_NAMECOMPONENT_UP_TYPE
type ICOMPONENT_NAMECOMPONENT_UP_TYPE interface {
	Test(ctx context.Context) error
}

// COMPONENT_NAMECOMPONENT_UP_TYPE
type COMPONENT_NAMECOMPONENT_UP_TYPE struct {
	config *config.Config
}

// NewCOMPONENT_NAMECOMPONENT_UP_TYPE
func NewCOMPONENT_NAMECOMPONENT_UP_TYPE(
	config *config.Config,
) *COMPONENT_NAMECOMPONENT_UP_TYPE {
	return &COMPONENT_NAMECOMPONENT_UP_TYPE{
		config: config,
	}
}

func (p *COMPONENT_NAMECOMPONENT_UP_TYPE) Test(ctx context.Context) error {
	log := logger.WithContext(ctx).With(zap.String("action", "Test"))
	log.Debug("xx", zap.String("xx", "xx"))
	return nil
}
