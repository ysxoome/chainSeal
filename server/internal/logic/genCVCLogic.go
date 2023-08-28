package logic

import (
	"context"

	"seal/blockchain/tool"
	"seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/service"
	"seal/server/common/utils"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenCVCLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenCVCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenCVCLogic {
	return &GenCVCLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenCVCLogic) GenCVC(req *types.CVCReq) (resp *types.CVCResp, err error) {
	company := models.CompanyVcEntity{
		LegalDid:    req.LegalDid,
		Name:        req.Name,
		LegalName:   req.LegalName,
		CreditCode:  req.CreditCode,
		CompanyType: req.CompanyType,
	}

	credential := service.GenCompanyVC(company)

	session := tool.InitIdentity()
	b, err := hexutil.Decode(credential.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "identity decode failed")
	}

	_, rec, err := session.AddIdentity(b)
	if err != nil && rec.Status != 0 {
		return nil, errors.Wrap(err, "identity tochain failed")
	}

	suc := engine.CreateSeal(req.LegalDid, req.CreditCode)
	if !suc {
		return nil, errors.Wrap(err, "create sealaddr failed")
	}

	CID, _ := utils.CredentialCID(l.svcCtx.Sh, credential)

	return &types.CVCResp{
		TransHash:   rec.TransactionHash,
		BlockHash:   rec.BlockHash,
		Identity:    credential.Hash(),
		SealAddress: engine.CalculateSeal(req.LegalDid, req.CreditCode),
		CVCID:       CID,
	}, nil
}
