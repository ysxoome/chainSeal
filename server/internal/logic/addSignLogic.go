package logic

import (
	"context"
	"math/big"
	"time"

	"seal/server/common/constant"
	"seal/server/common/engine"
	"seal/server/common/models"
	"seal/server/common/service"
	"seal/server/common/utils"
	"seal/server/config"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddSignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSignLogic {
	return &AddSignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSignLogic) AddSign(req *types.AddSignReq) (resp *types.AddSignResp, err error) {

	credential, _ := utils.CredentialFromCID(l.svcCtx.Sh, req.CVCID)
	res := engine.QueryIdentity(credential.Hash(), big.NewInt(time.Now().UnixNano()))
	if !res {
		return nil, errors.Wrap(err, "company information can not found")
	}

	nestContract, _ := utils.ContractFromCid(l.svcCtx.Sh, req.NestCID)
	if nestContract.IsContractEmpty() {
		nestContract = nil
	}

	entity := models.Contract{
		Context:       utils.GenDefaultCredentialContext(),
		ContractHash:  req.ContrachHash,
		NestSignature: nestContract,
		Id:            utils.RandomUUID(),
		SignerVC:      *credential,
		SignDate:      time.Now().String()[:19],
		SealType:      constant.COMPANY,
		SealsClaim:    req.CompanySealClaim,
	}
	entity.Proof = utils.GenContractProof(entity, config.ADMIN_PRIVATE_KEY)

	CID, _ := utils.ContractCID(l.svcCtx.Sh, entity)

	vaild, rec, signHash, sealAddr, signAddr := service.ToChain(entity)
	if !vaild {
		return nil, errors.Wrap(err, "contract to chain failed")
	}

	return &types.AddSignResp{
		TransHash:    rec.TransactionHash,
		BlockHash:    rec.BlockHash,
		ContractHash: entity.ContractHash,
		SignHash:     signHash,
		SealAddress:  sealAddr,
		SealType:     entity.SealType,
		SignAddress:  signAddr,
		CID:          CID,
	}, nil
}
