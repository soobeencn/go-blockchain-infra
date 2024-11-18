package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chain-infra/wallet/internal/logic"
	"go-chain-infra/wallet/internal/svc"
	"go-chain-infra/wallet/internal/types"
)

func WalletHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewWalletLogic(r.Context(), svcCtx)
		resp, err := l.Wallet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
