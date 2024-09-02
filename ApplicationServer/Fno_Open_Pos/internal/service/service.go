package service

import (
	"context"
	"fmt"
	pb"github.com/SudarshanZone/Fno_Open_Pos/generated"
	"github.com/SudarshanZone/Fno_Open_Pos/internal/repository"
)

type FnoPositionService struct {
	Repo repository.FnoPositionRepository
	pb.UnimplementedFnoPositionServiceServer
}

func (s *FnoPositionService) GetFNOPosition(ctx context.Context, req *pb.FnoPositionRequest) (*pb.FcpDetailListResponse, error) {
	positions, err := s.Repo.GetPositionsByClaimMatchAccount(req.FCP_CLM_MTCH_ACCNT)
	if err != nil {
		return nil, fmt.Errorf("failed to get positions: %w", err)
	}

	var fcpDetails []*pb.FcpDetail

	for _, pos := range positions {
		fcpDetail := &pb.FcpDetail{
			FFO_CONTRACT:     pos.Contract,
			FFO_PSTN:         pos.Position,
			FFO_QTY:          int32(pos.TotalQty),
			FFO_AVG_PRC:      float32(pos.AvgCostPrice),
			FCP_XCHNG_CD:     pos.ExchangeCode,
			FCP_IBUY_QTY:     int32(pos.BuyQty),
			FCP_CLM_MTCH_ACCNT: pos.ClaimMatchAccount,
			FCP_PRDCT_TYP:    pos.ProductType,
			FCP_INDSTK:       pos.IndexStock,
			FCP_UNDRLYNG:     pos.Underlying,
			FCP_EXPRY_DT:     pos.ExpiryDate,
			FCP_EXER_TYP:     pos.ExerciseType,
			FCP_OPT_TYP:      pos.OptionType,
			FCP_STRK_PRC:     float32(pos.StrikePrice),
			FCP_UCC_CD:       pos.UCCCode,
			FCP_OPNPSTN_FLW:  pos.OpenPstnFlow,
		}

		fcpDetails = append(fcpDetails, fcpDetail)
	}

	return &pb.FcpDetailListResponse{
		FcpDetails: fcpDetails,
	}, nil
}

