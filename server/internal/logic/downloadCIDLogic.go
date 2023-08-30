package logic

import (
	"context"
	"fmt"

	"seal/blockchain/tool"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadCIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCIDLogic {
	return &DownloadCIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCIDLogic) DownloadCID(req *types.DownloadCidReq) (resp *types.DownloadCidResp, err error) {
	err = tool.DownloadFile(l.svcCtx.Sh, req.CID, fmt.Sprintf("%s/%s.json", req.Path, req.CID))
	if err != nil {
		return &types.DownloadCidResp{
			Res: "download cid failed",
		}, err
	}

	return &types.DownloadCidResp{
		Res: "success",
	}, nil
}
