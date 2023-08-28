package logic

import (
	"context"

	"seal/blockchain/tool"
	"seal/server/common/models"
	"seal/server/common/service"
	"seal/server/common/utils"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenPVCLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenPVCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenPVCLogic {
	return &GenPVCLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenPVCLogic) GenPVC(req *types.PVCReq) (resp *types.PVCResp, err error) {
	keyStore, _ := utils.CreateCredentials(req.PPriKeyPath)

	personVC := models.PersonVC{
		KeyStore: keyStore,
		Name:     req.Name,
		IdNumber: req.IdNumber,
	}

	credential := service.GenPersonVC(personVC)

	session := tool.InitIdentity()
	b, err := hexutil.Decode(credential.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "identity decode failed")
	}

	_, rec, err := session.AddIdentity(b)
	if err != nil && rec.Status != 0 {
		return nil, errors.Wrap(err, "identity tochain failed")
	}

	CID, _ := utils.CredentialCID(l.svcCtx.Sh, credential)

	return &types.PVCResp{
		TransHash: rec.TransactionHash,
		BlockHash: rec.BlockHash,
		Identity:  credential.Hash(),
		PVCID:     CID,
	}, nil
}
