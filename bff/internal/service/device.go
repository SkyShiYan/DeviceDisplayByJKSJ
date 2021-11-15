package service

import (
	"context"

	v4 "bff/api/device/v4"
	"bff/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// BffService is a greeter service.
type BffService struct {
	v4.UnimplementedDeviceServer
	buc *biz.BffUsecase
	log *log.Helper
}

func NewBffService(uc *biz.BffUsecase, logger log.Logger) *BffService {
	return &BffService{buc: uc, log: log.NewHelper(logger)}
}

func (s *BffService) register(ctx context.Context, in *v4.RegisterRequest) (*v4.RegisterReply, error) {
	err := s.buc.CreateDevice(ctx, &biz.Device{
		HardwareKey: in.GetHardwareKey(),
	})
	return nil, err
}
