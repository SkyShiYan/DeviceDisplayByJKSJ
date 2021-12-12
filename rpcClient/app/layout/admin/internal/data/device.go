package data

import (
	"context"
	"encoding/json"
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
	data, redisErr := r.data.rClient.Get(*key).Result()
	if redisErr == nil {
		r.log.Warn("redis获取key:" + *key + "成功")
		var res biz.Device
		err := json.Unmarshal([]byte(data), &res)
		if err != nil {
			r.log.Warn("redis反序列化失败-" + data)
			return nil, err
		}
		return &res, nil
	} else {
		r.log.Warn("redis获取key:" + *key + "失败")
	}

	deviceDB, err := r.data.dbClient.Device.Query().Where(device.HardwareKeyEqualFold(*key)).First(ctx)
	if err != nil {
		return nil, err
	}

	res := &biz.Device{
		Id:              int64(deviceDB.ID),
		Name:            *deviceDB.Name,
		HardwareKey:     *deviceDB.HardwareKey,
		Defaultlayoutid: *deviceDB.DefaultLayoutId,
		Status:          *deviceDB.Status,
		Storenumber:     *deviceDB.StoreNumber,
	}

	bytes, jsonErr := json.Marshal(res)
	if jsonErr == nil {
		redisErr = r.data.rClient.Set(*key, string(bytes), 0).Err()
		if redisErr != nil {
			r.log.Warn("redis保存key:" + *key + "失败")
		}
	}

	return res, nil
}
