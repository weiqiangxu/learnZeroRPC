package logic

import (
	"context"
	"errors"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

// NewGetOrderLogic 获取Order的logic层对象
func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetOrderLogic {
    return GetOrderLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

// GetOrder 获取订单信息
func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (*types.OrderReply, error) {
    // logic层调用RPC接口获取用户信息
    user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
        Id: "1",
    })
    if err != nil {
        return nil, err
    }

    if user.Name != "test" {
        return nil, errors.New("用户不存在")
    }

    return &types.OrderReply{
        Id:   req.Id,
        Name: "test order",
    }, nil
}