package rpc_server

import (
	"MODULE_NAME/api/proto"
	"MODULE_NAME/internal/constants"
	"context"
	"errors"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// SERVER_NAMEServerProvider
var SERVER_NAMEServerProvider = wire.NewSet(NewSERVER_NAMEServer, wire.Bind(new(iSERVER_NAMEServer), new(*SERVER_NAMEServer)))

// iSERVER_NAMEServer
type iSERVER_NAMEServer interface {
	Get(ctx context.Context, req *proto.GetReq) (*proto.GetRes, error)
}

// SERVER_NAMEServer
type SERVER_NAMEServer struct {
	logger *zap.Logger
	// Introduce new dependencies here, exp:
	// testService  *services.TestService
}

// NewSERVER_NAMEServer
func NewSERVER_NAMEServer(
	logger *zap.Logger,
	// testService *services.TestService,
) (*SERVER_NAMEServer, error) {
	return &SERVER_NAMEServer{
		logger: logger.With(zap.String("type", "SERVER_SERVER_NAMEServer")),
		// testService:  testService,
	}, nil
}

func (s *SERVER_NAMEServer) Get(ctx context.Context, req *proto.GetReq) (*proto.GetRes, error) {
	if req == nil {
		return nil, errors.New(constants.GetErrMsg(constants.PARAMS_INVALID))
	}
	// p, err := s.testService.Get(ctx, req.Id)
	// if err != nil {
	//	p.logger.Error("Get", zap.Error(err))
	// 	return nil, errors.Wrap(err, constants.GetErrMsg(constants.INTERNAL_ERR))
	// }

	resp := &proto.GetRes{
		// ...
	}

	return resp, nil
}
