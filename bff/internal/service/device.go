package service

import (
	"context"

	v4 "bff/api/device/v4"
	"bff/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// BffService is a greeter service.
type BffService struct {
	v4.UnimplementedBffDeviceServer
	buc *biz.BffUsecase
	log *log.Helper
}

func NewBffService(uc *biz.BffUsecase, logger log.Logger) *BffService {
	return &BffService{buc: uc, log: log.NewHelper(logger)}
}

func (s *BffService) ChangeName(ctx context.Context, in *v4.ChangeNameRequest) (*v4.ChangeNameReply, error) {
	// err := s.buc.CreateDevice(ctx, &biz.Device{
	// 	HardwareKey: in.GetHardwareKey(),
	// })
	return nil, nil
}
func (s *BffService) GetDevice(ctx context.Context, in *v4.GetDeviceRequest) (*v4.GetDeviceReply, error) {
	s.log.Debug("bff-开始获取数据")

	d, err := s.buc.GetDevice(ctx, in.HardwareKey)
	if err != nil {
		return nil, err
	}

	return &v4.GetDeviceReply{
		Name:            d.Name,
		HardwareKey:     d.HardwareKey,
		Defaultlayoutid: d.Defaultlayoutid,
		Status:          d.Status,
		Storenumber:     d.Storenumber,
	}, nil
}
func (s *BffService) Register(ctx context.Context, in *v4.RegisterRequest) (*v4.RegisterReply, error) {
	return nil, nil
}
func (s *BffService) GetDeviceDisplay(ctx context.Context, req *v4.GetDeviceDisplayRequest) (*v4.GetDeviceDisplayReply, error) {
	// layouts, err := s.buc.GetLayoutByDevice(ctx, req.HardwareKey)
	layouts, err := s.buc.GetDeviceDisplay(ctx, req.HardwareKey)
	if err != nil {
		return nil, err
	}
	if layouts == nil {
		return nil, nil
	}

	var data []*v4.GetDeviceDisplayReply_DisplayItem
	for _, layout := range layouts {
		data = append(data, &v4.GetDeviceDisplayReply_DisplayItem{
			Name: layout.Name,
			Md5:  layout.Md5,
		})
	}
	return &v4.GetDeviceDisplayReply{
		DisplayList: data,
	}, nil
}
