package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
	"github.com/orzkratos/ebzkratos"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, *ebzkratos.Ebz) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, *ebzkratos.Ebz) {
	return g, nil
}

// FindByID finds a Greeter by ID
// ebzkratos.New wraps Kratos error with stack trace at call site
//
// FindByID 根据 ID 查找 Greeter
// ebzkratos.New 在调用处包装 Kratos 错误并附带堆栈信息
func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, *ebzkratos.Ebz) {
	return nil, ebzkratos.New(v1.ErrorUnknown("NOT IMPLEMENTED"))
}

// ListByHello lists Greeters by hello string
// ebzkratos.New wraps Kratos error with stack trace at call site
//
// ListByHello 根据 hello 字符串列出 Greeter
// ebzkratos.New 在调用处包装 Kratos 错误并附带堆栈信息
func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, *ebzkratos.Ebz) {
	return nil, ebzkratos.New(v1.ErrorUserNotFound("ERROR OCCURRED"))
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, *ebzkratos.Ebz) {
	return nil, nil
}
