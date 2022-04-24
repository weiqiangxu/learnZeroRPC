package handler

import (
	"net/http"

	"go-zero-demo/mall/order/api/internal/logic"
	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析请求参数为结构体对象
		var req types.OrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		// 初始化
		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
