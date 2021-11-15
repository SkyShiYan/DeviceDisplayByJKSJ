package data

import (
	v4 "bff/api/device/v4"
	"bff/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	registry "github.com/go-kratos/consul/registry"
	consulApi "github.com/hashicorp/consul/api"
)

type deviceRepo struct {
	data *Data
	log  *log.Helper
}

// NewBffRepo .
func NewBffRepo(data *Data, logger log.Logger) biz.DeviceRepo {
	return &deviceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *deviceRepo) CreateDevice(ctx context.Context, g *biz.Device) error {
	defaultConfig := consulApi.DefaultConfig()
	defaultConfig.Address = "8.130.28.195:8500"
	client, err := consulApi.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	dis := registry.New(client)

	endpoint := "discovery:///device"
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	if err != nil {
		return err
	}
	gClient := v4.NewClientDeviceClient(conn)
	replay, err := gClient.AddDeviceByKey(ctx, &v4.AddDeviceByKeyRequest{Name: g.Name,
		HardwareKey:     g.HardwareKey,
		Defaultlayoutid: g.Defaultlayoutid,
		Status:          g.Status,
		Storenumber:     g.Storenumber})

	if err != nil {
		return err
	}
	return v4.ErrorAddErr(replay.Message)
}

func (r *deviceRepo) UpdateDevice(ctx context.Context, g *biz.Device) error {
	return nil
}

func (r *deviceRepo) GetDevice(ctx context.Context, hardwareKey *string) (*biz.Device, error) {
	return nil, nil
}
