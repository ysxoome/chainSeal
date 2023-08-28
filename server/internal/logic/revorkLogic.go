package logic

import (
	"context"
	"encoding/base64"
	"time"

	"seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/service"
	"seal/server/common/utils"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RevorkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRevorkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RevorkLogic {
	return &RevorkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RevorkLogic) Revork(req *types.RevokeReq) (resp *types.RevokeResp, err error) {
	keyStore, _ := utils.CreateCredentials(req.CPriKeyPath)

	signatureAuth := models.SignatureAuth{
		SealAddress: req.SealAddress,
		TypeAuth:    1,
		TargetDid:   req.TargetDid,
		AuthDid:     req.AuthDid,
		AuthTime:    time.Now().UnixMilli() / 1000,
	}

	encodePacked := engine.EncodePacked(signatureAuth.SealAddress,
		signatureAuth.TypeAuth,
		signatureAuth.TargetDid,
		signatureAuth.AuthTime)
	hash := crypto.Keccak256Hash(encodePacked)
	signatureData, err := crypto.Sign(hash[:], keyStore.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "签名失败")
	}

	signatureAuth.AuthSign = base64.StdEncoding.EncodeToString(signatureData)
	res, rec := service.RevokeSeal(signatureAuth)
	if !res {
		return nil, errors.Wrap(err, "签章授权失败")
	}

	if rec == nil {
		return &types.RevokeResp{
			TransHash: "",
			BlockHash: "",
			Res:       "have been revorked, no option to do",
		}, nil
	}

	return &types.RevokeResp{
		TransHash: rec.TransactionHash,
		BlockHash: rec.BlockHash,
		Res:       "revoke success",
	}, nil
}
