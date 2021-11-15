package data

import (
	"context"
	"rpcClient/app/device/service/ent/device"
	"rpcClient/app/device/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type deviceRepo struct {
	data *Data
	log  *log.Helper
}

// NewDeviceRepo .
func NewDeviceRepo(data *Data, logger log.Logger) biz.DeviceRepo {
	return &deviceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *deviceRepo) CreateDevice(ctx context.Context, g *biz.Device) error {
	_, err := r.data.dbClient.Device.Create().SetName(g.Name).SetHardwareKey(g.HardwareKey).
		SetDefaultLayoutId(g.Defaultlayoutid).SetStatus(g.Status).SetStoreNumber(g.Storenumber).Save(ctx)
	return err

	// return nil

	// _, err := r.data.dbClient.Device.Create().SetName("1").SetHardwareKey("2").
	// 	SetDefaultLayoutId(3).SetStatus(4).SetStoreNumber("5").Save(ctx)
	// return err
}

func (r *deviceRepo) UpdateDevice(ctx context.Context, g *biz.Device) error {
	_, err := r.data.dbClient.Device.UpdateOneID(int(g.Id)).SetName(g.Name).Save(ctx)
	return err
}

func (r *deviceRepo) GetDevice(ctx context.Context, key *string) (*biz.Device, error) {
	deviceDB, err := r.data.dbClient.Device.Query().Where(device.HardwareKeyEqualFold(*key)).First(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.Device{
		Id:              int64(deviceDB.ID),
		Name:            *deviceDB.Name,
		HardwareKey:     *deviceDB.HardwareKey,
		Defaultlayoutid: *deviceDB.DefaultLayoutId,
		Status:          *deviceDB.Status,
		Storenumber:     *deviceDB.StoreNumber,
	}, nil
}
