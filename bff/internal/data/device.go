package data

import (
	v4 "bff/api/device/v4"
	"bff/internal/biz"
	"context"
	"encoding/json"
	"time"

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
	// if conn != nil {

	// }
	// return nil
	gClient := v4.NewDeviceClient(conn)
	replay, err := gClient.UpdateDeviceByKey(ctx, &v4.UpdateDeviceByKeyRequest{
		HardwareKey: g.HardwareKey})

	if err != nil {
		return err
	}
	return v4.ErrorAddErr(replay.Message)
}

func (r *deviceRepo) UpdateDevice(ctx context.Context, g *biz.Device) error {
	return nil
}

func (r *deviceRepo) GetDevice(ctx context.Context, hardwareKey string) (*biz.Device, error) {
	r.log.Debug("111111")
	rD, err := r.data.rClient.Get("bff-" + hardwareKey).Result()
	if err == nil {
		var d biz.Device
		jErr := json.Unmarshal([]byte(rD), &d)
		if jErr == nil {
			r.log.Debug("从redis中获取数据")
			return &d, nil
		} else {
			r.log.Warn("redis获取key:" + hardwareKey + "json Unmarshal 失败")
		}
	}

	defaultConfig := consulApi.DefaultConfig()
	defaultConfig.Address = "8.130.28.195:8500"
	client, err := consulApi.NewClient(defaultConfig)
	if err != nil {
		return nil, err
	}
	dis := registry.New(client)

	// list, err = dis.ListServices()
	// if err != nil {
	// 	return nil, err
	// }

	endpoint := "discovery:///deviceRpc"
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	if err != nil {
		return nil, err
	}

	gClient := v4.NewDeviceClient(conn)
	replay, err := gClient.GetDeviceByKey(ctx, &v4.GetDeviceByKeyRequest{
		HardwareKey: hardwareKey,
	})

	if err != nil {
		return nil, err
	}

	d := &biz.Device{
		Name:            replay.Name,
		HardwareKey:     replay.HardwareKey,
		Defaultlayoutid: replay.Defaultlayoutid,
		Status:          replay.Status,
		Storenumber:     replay.Storenumber,
	}

	bytes, jErr := json.Marshal(d)
	if jErr != nil {
		r.data.rClient.Set("bff-"+hardwareKey, bytes, time.Duration(time.Minute.Minutes()))
	}

	return d, nil
}
