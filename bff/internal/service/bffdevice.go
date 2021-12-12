package service

import (
	"context"

	pb "bff/api/device/v4"
)

type BffDeviceService struct {
	pb.UnimplementedBffDeviceServer
}

func NewBffDeviceService() *BffDeviceService {
	return &BffDeviceService{}
}

func (s *BffDeviceService) registeDevice(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *BffDeviceService) changeName(ctx context.Context, req *pb.ChangeNameRequest) (*pb.ChangeNameReply, error) {
	return &pb.ChangeNameReply{}, nil
}
func (s *BffDeviceService) getDevice(ctx context.Context, req *pb.GetDeviceRequest) (*pb.GetDeviceReply, error) {
	return &pb.GetDeviceReply{}, nil
}
func (s *BffDeviceService) getDeviceDisplay(ctx context.Context, req *pb.GetDeviceDisplayRequest) (*pb.GetDeviceDisplayReply, error) {
	return &pb.GetDeviceDisplayReply{}, nil
}
