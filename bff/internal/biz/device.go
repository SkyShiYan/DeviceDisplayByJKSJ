package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type DeviceRepo interface {
	CreateDevice(context.Context, *Device) error
	UpdateDevice(context.Context, *Device) error
	GetDevice(ctx context.Context, hardwareKey *string) (*Device, error)
}

type Device struct {
	Name            string
	HardwareKey     string
	Defaultlayoutid int32
	Status          string
	Storenumber     string
}

type BffUsecase struct {
	repo DeviceRepo
	log  *log.Helper
}

func NewBffUsecase(repo DeviceRepo, logger log.Logger) *BffUsecase {
	return &BffUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *BffUsecase) CreateDevice(ctx context.Context, d *Device) error {
	return uc.repo.CreateDevice(ctx, d)
}

func (uc *BffUsecase) UpdateDevice(ctx context.Context, d *Device) error {
	return nil
}

func (uc *BffUsecase) GetDevice(ctx context.Context, d *Device) (*Device, error) {
	return nil, nil
}
