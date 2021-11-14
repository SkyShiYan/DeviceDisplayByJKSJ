package service

import (
	"context"

	v4 "spaco-1103/api/device/v4"
	v1 "spaco-1103/api/helloworld/v1"
	"spaco-1103/app/layout/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v4.UnimplementedDeviceServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v4.RegisterRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetHardwareKey())

	if in.GetHardwareKey() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetHardwareKey())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetHardwareKey()}, nil
}
