package service

import (
	"github.com/google/wire"
	v1 "realWorld/api/realworld/v1"
	"realWorld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.GreeterUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.GreeterUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
