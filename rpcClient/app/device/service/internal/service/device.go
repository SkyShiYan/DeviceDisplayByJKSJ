package service

import (
	"context"
	v4 "rpcClient/api/device/v4"
	"rpcClient/app/device/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// DeviceService is a Device service.
type DeviceService struct {
	v4.UnimplementedDeviceServer

	uc  *biz.DeviceUsecase
	log *log.Helper
}

// NewDeviceService new a Device service.
func NewDeviceService(uc *biz.DeviceUsecase, logger log.Logger) *DeviceService {
	return &DeviceService{uc: uc, log: log.NewHelper(logger)}
}

func (s *DeviceService) AddDeviceByKey(ctx context.Context, in *v4.AddDeviceByKeyRequest) (*v4.AddDeviceByKeyReply, error) {
	err := s.uc.Create(ctx, &biz.Device{
		Name:            in.Name,
		HardwareKey:     in.HardwareKey,
		Defaultlayoutid: in.Defaultlayoutid,
		Status:          in.Status,
		Storenumber:     in.Storenumber,
	})
	return nil, err
}
func (s *DeviceService) GetDeviceByKey(ctx context.Context, in *v4.GetDeviceByKeyRequest) (*v4.GetDeviceByKeyReply, error) {
	device, err := s.uc.Get(ctx, &in.HardwareKey)
	if err != nil {
		return nil, err
	}
	return &v4.GetDeviceByKeyReply{
		Name:            device.Name,
		HardwareKey:     device.HardwareKey,
		Defaultlayoutid: device.Defaultlayoutid,
		Status:          device.Status,
		Storenumber:     device.Storenumber,
	}, nil
}
func (s *DeviceService) UpdateDeviceByKey(ctx context.Context, in *v4.UpdateDeviceByKeyRequest) (*v4.UpdateDeviceByKeyReply, error) {
	err := s.uc.Update(ctx, &biz.Device{
		HardwareKey: in.HardwareKey,
	})
	return nil, err
}
