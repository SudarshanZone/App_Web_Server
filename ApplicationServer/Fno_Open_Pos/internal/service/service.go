package service

import (
	"context"
	"fmt"
	"github.com/SudarshanZone/Fno_Open_Pos/internal/logger"
	pb"github.com/SudarshanZone/Fno_Open_Pos/generated"
	"github.com/SudarshanZone/Fno_Open_Pos/internal/repository"
)

type FnoPositionService struct {
	Repo repository.FnoPositionRepository
	pb.UnimplementedFnoPositionServiceServer
}

func (s *FnoPositionService) GetFNOPosition(ctx context.Context, req *pb.FnoPositionRequest) (*pb.FcpDetailListResponse, error) {
	//log := logger.GetLogger()

	defer logger.LogFunctionExit("GetFNOPosition")
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

		s.logPositionDetails("FnoPositionService", fcpDetail)
		fcpDetails = append(fcpDetails, fcpDetail)
	}

	
	return &pb.FcpDetailListResponse{
		FcpDetails: fcpDetails,
	}, nil
}

func (s *FnoPositionService) logPositionDetails(c_ServiceName string, pos *pb.FcpDetail) {
	log := logger.GetLogger()
	log.Infof("%s: Started", c_ServiceName)
	log.Infof("%s: @@@@@@@@@@@@@@@@ FETCH RECORD @@@@@@@@@@@@@@@@:", c_ServiceName)
	log.Infof("%s: :FCP_CLM_MTCH_ACCNT:%s:", c_ServiceName, pos.FCP_CLM_MTCH_ACCNT)
	log.Infof("%s: :FFO_CONTRACT :%s:", c_ServiceName, pos.FFO_CONTRACT)
	log.Infof("%s: :FFO_PSTN :%s:", c_ServiceName, pos.FFO_PSTN)
	log.Infof("%s: :FFO_QTY :%d:", c_ServiceName, pos.FFO_QTY)
	log.Infof("%s: :FFO_AVG_PRC :%s:", c_ServiceName, pos.FFO_CONTRACT)
	// log.Infof("%s: :FCP_PRDCT_TYP:%s:", c_ServiceName, pos.FCP_PRDCT_TYP)
	// log.Infof("%s: :FCP_INDSTK:%s:", c_ServiceName, pos.FCP_INDSTK)
	// log.Infof("%s: :FCP_UNDRLYNG:%s:", c_ServiceName, pos.FCP_UNDRLYNG)
	// log.Infof("%s: :FCP_EXPRY_DT:%s:", c_ServiceName, pos.FCP_EXPRY_DT)
	// log.Infof("%s: :FCP_EXER_TYP:%s:", c_ServiceName, pos.FCP_EXER_TYP)
	// log.Infof("%s: :FCP_OPT_TYP:%s:", c_ServiceName, pos.FCP_OPT_TYP)
	// log.Infof("%s: :FCP_STRK_PRC:%f:", c_ServiceName, pos.FCP_STRK_PRC)
	// log.Infof("%s: :FCP_OPNPSTN_FLW:%s:", c_ServiceName, pos.FCP_OPNPSTN_FLW)
	// log.Infof("%s: :FFO_QTY:%d:", c_ServiceName, pos.FFO_QTY)
	// log.Println()
}

