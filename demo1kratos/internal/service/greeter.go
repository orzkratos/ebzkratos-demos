package service

import (
	"context"

	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
// ebz is *ebzkratos.Ebz which wraps error with stack trace
// ebz.Erk returns the original Kratos error for gRPC/HTTP response
//
// SayHello 实现 helloworld.GreeterServer 接口
// ebz 是 *ebzkratos.Ebz 类型，包装错误并附带堆栈信息
// ebz.Erk 返回原始 Kratos 错误用于 gRPC/HTTP 响应
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, ebz := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if ebz != nil {
		return nil, ebz.Erk // Return original Kratos error // 返回原始 Kratos 错误
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
