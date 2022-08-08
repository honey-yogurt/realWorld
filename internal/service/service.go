package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "realWorld/api/realworld/v1"
	"realWorld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc  *biz.UserUseCase
	log *log.Helper
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.UserUseCase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, log: log.NewHelper(logger)}
}
