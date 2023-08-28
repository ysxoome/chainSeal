package logic

import (
	"context"
	"math/big"
	"time"

	"seal/server/common/engine"
	"seal/server/internal/svc"
	"seal/server/internal/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySignLogic {
	return &QuerySignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuerySignLogic) QuerySign(req *types.QuerySignReq) (resp *types.QuerySignResp, err error) {
	index := engine.Query(req.ContractHash, req.SignHash, req.SealAddress, req.SealType, req.SignAddress)
	res := engine.QueryDetail(req.ContractHash, index.Uint64()-1)
	v, ok := res.(struct {
		Cid      string
		SignHash []byte
		Sealaddr common.Address
		SealType uint8
		SignAddr common.Address
		SignTime *big.Int
	})
	if !ok {
		return nil, errors.Wrap(nil, "interface assertion failed")
	}

	return &types.QuerySignResp{
		CID:         v.Cid,
		SignHash:    hexutil.Encode(v.SignHash),
		SealAddress: v.Sealaddr.String(),
		SealType:    v.SealType,
		SignAddress: v.SignAddr.String(),
		SignTime:    BigIntToTime(v.SignTime),
	}, nil
}

func BigIntToTime(bi *big.Int) string {
	ms := bi.Int64()
	sec := ms / 1000
	nsec := (ms % 1000) * 1000000
	t := time.Unix(sec, nsec)
	return t.String()[:19]
}
