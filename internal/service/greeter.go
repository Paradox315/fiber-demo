package service

import (
	"context"
	"fiber-demo/internal/biz"

	pb "fiber-demo/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

func (s *GreeterService) List(ctx context.Context, req *pb.PageRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("List Received: %v", req.String())
	return &pb.HelloReply{
		Message: req.String(),
	}, nil
}
func (s *GreeterService) Get(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Get Received: %v", req.String())
	return &pb.HelloReply{
		Message: req.String(),
	}, nil
}
func (s *GreeterService) Add(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Add Received: %v", req.String())
	return &pb.HelloReply{
		Message: req.String(),
	}, nil
}
func (s *GreeterService) Edit(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Edit Received: %v", req.String())
	return &pb.HelloReply{
		Message: req.String(),
	}, nil
}
func (s *GreeterService) Del(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.log.WithContext(ctx).Infof("Del Received: %v", req.String())
	return &pb.HelloReply{
		Message: req.String(),
	}, nil
}
