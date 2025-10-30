package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
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

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, *ebzkratos.Ebz) {
	return nil, ebzkratos.New(v1.ErrorUnknown("NOT IMPLEMENTED"))
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, *ebzkratos.Ebz) {
	return nil, ebzkratos.New(v1.ErrorUserNotFound("ERROR OCCURRED"))
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, *ebzkratos.Ebz) {
	return nil, nil
}
