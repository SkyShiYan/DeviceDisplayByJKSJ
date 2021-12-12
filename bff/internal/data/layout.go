package data

import (
	"bff/internal/biz"
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
)

type layoutRepo struct {
	data *Data
	log  *log.Helper
}

// NewBffLayoutRepo .
func NewBffLayoutRepo(data *Data, logger log.Logger) biz.LayoutRepo {
	return &layoutRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (lRepo *layoutRepo) GetLayoutByDevice(ctx context.Context, hardwareKey string) (layouts []*biz.Layout, err error) {
	lRepo.log.Debug("asdasdasdad-------hardwareKey=" + hardwareKey + "--hardwareKey==1:" + strconv.FormatBool(hardwareKey == "1"))

	if hardwareKey == "1" {
		var d = []*biz.Layout{{
			Name: "1",
			Md5:  "2",
		}, {
			Name: "11",
			Md5:  "22",
		}}

		return d, nil
	}
	return nil, nil
}
