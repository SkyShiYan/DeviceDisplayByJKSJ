package biz

import (
	"context"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
	"github.com/go-kratos/kratos/v2/log"
	// "golang.org/x/sync/errgroup"
)

type DeviceRepo interface {
	CreateDevice(context.Context, *Device) error
	UpdateDevice(context.Context, *Device) error
	GetDevice(ctx context.Context, hardwareKey string) (*Device, error)
}

type Device struct {
	Name            string
	HardwareKey     string
	Defaultlayoutid int32
	Status          int32
	Storenumber     string
}

type LayoutRepo interface {
	GetLayoutByDevice(ctx context.Context, hardwareKey string) ([]*Layout, error)
}

type Layout struct {
	Name string
	Md5  string
}

type BffUsecase struct {
	repo       DeviceRepo
	layoutRepo LayoutRepo
	log        *log.Helper
}

func NewBffUsecase(deviceRepo DeviceRepo, layoutRepo LayoutRepo, logger log.Logger) *BffUsecase {
	return &BffUsecase{repo: deviceRepo, layoutRepo: layoutRepo, log: log.NewHelper(logger)}
}

func (uc *BffUsecase) CreateDevice(ctx context.Context, d *Device) error {
	return uc.repo.CreateDevice(ctx, d)
}

func (uc *BffUsecase) UpdateDevice(ctx context.Context, d *Device) error {
	return nil
}

func (uc *BffUsecase) GetDevice(ctx context.Context, hardwareKey string) (*Device, error) {
	return uc.repo.GetDevice(ctx, hardwareKey)
}

func (uc *BffUsecase) GetDeviceDisplay(ctx context.Context, hardwareKey string) ([]*Layout, error) {
	var errG = errgroup.WithContext(ctx)
	var s struct {
		*Device
		ls []*Layout
	}
	errG.Go(func(ctx context.Context) error {
		d, err := uc.GetDevice(ctx, hardwareKey)
		s.Device = d
		if err != nil {
			return err
		}
		return nil
	})
	errG.Go(func(ctx context.Context) error {
		layout, err := uc.GetLayoutByDevice(ctx, hardwareKey)
		s.ls = layout
		if err != nil {
			return err
		}
		return nil
	})

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	return s.ls, nil
}

func (uc *BffUsecase) GetLayoutByDevice(ctx context.Context, hardwareKey string) ([]*Layout, error) {
	return uc.layoutRepo.GetLayoutByDevice(ctx, hardwareKey)
}
