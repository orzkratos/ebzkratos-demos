# Changes

Code differences compared to source project demokratos.

## internal/biz/greeter.go (+11 -6)

```diff
@@ -6,6 +6,7 @@
 	"github.com/go-kratos/kratos/v2/errors"
 	"github.com/go-kratos/kratos/v2/log"
 	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
+	"github.com/orzkratos/ebzkratos"
 )
 
 var (
@@ -19,12 +20,16 @@
 }
 
 // GreeterRepo is a Greater repo.
+// Use *ebzkratos.Ebz instead of error to get stack trace on failures
+//
+// GreeterRepo 是 Greeter 仓库接口
+// 使用 *ebzkratos.Ebz 替代 error 以便在失败时获取堆栈信息
 type GreeterRepo interface {
-	Save(context.Context, *Greeter) (*Greeter, error)
-	Update(context.Context, *Greeter) (*Greeter, error)
-	FindByID(context.Context, int64) (*Greeter, error)
-	ListByHello(context.Context, string) ([]*Greeter, error)
-	ListAll(context.Context) ([]*Greeter, error)
+	Save(context.Context, *Greeter) (*Greeter, *ebzkratos.Ebz)
+	Update(context.Context, *Greeter) (*Greeter, *ebzkratos.Ebz)
+	FindByID(context.Context, int64) (*Greeter, *ebzkratos.Ebz)
+	ListByHello(context.Context, string) ([]*Greeter, *ebzkratos.Ebz)
+	ListAll(context.Context) ([]*Greeter, *ebzkratos.Ebz)
 }
 
 // GreeterUsecase is a Greeter usecase.
@@ -39,7 +44,7 @@
 }
 
 // CreateGreeter creates a Greeter, and returns the new Greeter.
-func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
+func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, *ebzkratos.Ebz) {
 	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
 	return uc.repo.Save(ctx, g)
 }
```

## internal/data/greeter.go (+19 -7)

```diff
@@ -4,7 +4,9 @@
 	"context"
 
 	"github.com/go-kratos/kratos/v2/log"
+	v1 "github.com/orzkratos/demokratos/demo2kratos/api/helloworld/v1"
 	"github.com/orzkratos/demokratos/demo2kratos/internal/biz"
+	"github.com/orzkratos/ebzkratos"
 )
 
 type greeterRepo struct {
@@ -20,22 +22,32 @@
 	}
 }
 
-func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
+func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, *ebzkratos.Ebz) {
 	return g, nil
 }
 
-func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
+func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, *ebzkratos.Ebz) {
 	return g, nil
 }
 
-func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
-	return nil, nil
+// FindByID finds a Greeter by ID
+// ebzkratos.New wraps Kratos error with stack trace at call site
+//
+// FindByID 根据 ID 查找 Greeter
+// ebzkratos.New 在调用处包装 Kratos 错误并附带堆栈信息
+func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, *ebzkratos.Ebz) {
+	return nil, ebzkratos.New(v1.ErrorUnknown("NOT IMPLEMENTED"))
 }
 
-func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
-	return nil, nil
+// ListByHello lists Greeters by hello string
+// ebzkratos.New wraps Kratos error with stack trace at call site
+//
+// ListByHello 根据 hello 字符串列出 Greeter
+// ebzkratos.New 在调用处包装 Kratos 错误并附带堆栈信息
+func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, *ebzkratos.Ebz) {
+	return nil, ebzkratos.New(v1.ErrorUserNotFound("ERROR OCCURRED"))
 }
 
-func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
+func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, *ebzkratos.Ebz) {
 	return nil, nil
 }
```

## internal/service/greeter.go (+9 -3)

```diff
@@ -20,10 +20,16 @@
 }
 
 // SayHello implements helloworld.GreeterServer.
+// ebz is *ebzkratos.Ebz which wraps error with stack trace
+// ebz.Erk returns the original Kratos error for gRPC/HTTP response
+//
+// SayHello 实现 helloworld.GreeterServer 接口
+// ebz 是 *ebzkratos.Ebz 类型，包装错误并附带堆栈信息
+// ebz.Erk 返回原始 Kratos 错误用于 gRPC/HTTP 响应
 func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
-	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
-	if err != nil {
-		return nil, err
+	g, ebz := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
+	if ebz != nil {
+		return nil, ebz.Erk // Return original Kratos error // 返回原始 Kratos 错误
 	}
 	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
 }
```

