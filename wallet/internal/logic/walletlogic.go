package logic

import (
	"context"

	"go-chain-infra/wallet/internal/svc"
	"go-chain-infra/wallet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLogic {
	return &WalletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WalletLogic) Wallet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{
		Message: "hello" + req.Name,
	}
	return
}
