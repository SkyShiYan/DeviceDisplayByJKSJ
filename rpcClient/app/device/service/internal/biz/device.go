package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Device struct {
	Id              int64
	Name            string
	HardwareKey     string
	Defaultlayoutid int32
	Status          int32
	Storenumber     string
}

type DeviceRepo interface {
	CreateDevice(context.Context, *Device) error
	UpdateDevice(context.Context, *Device) error
	GetDevice(ctx context.Context, key *string) (*Device, error)
}

type DeviceUsecase struct {
	repo DeviceRepo
	log  *log.Helper
}

func NewDeviceUsecase(repo DeviceRepo, logger log.Logger) *DeviceUsecase {
	return &DeviceUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DeviceUsecase) Create(ctx context.Context, g *Device) error {
	return uc.repo.CreateDevice(ctx, g)
}

func (uc *DeviceUsecase) Update(ctx context.Context, g *Device) error {
	return uc.repo.UpdateDevice(ctx, g)
}

func (uc *DeviceUsecase) Get(ctx context.Context, key *string) (*Device, error) {
	return uc.repo.GetDevice(ctx, key)
}
