// ApplicationServer\Equity_Pos_Service\internal\service\equity_service.go
package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/SudarshanZone/MTF_OPEN_POS/internal/repository"
	pb "github.com/SudarshanZone/MTF_OPEN_POS/generated"
	"gorm.io/gorm"
)

type EquityService struct {
	pb.UnimplementedMtfPositionServiceServer
	repo repository.EquityRepository
}

func NewEquityService(repo repository.EquityRepository) *EquityService {
	return &EquityService{repo: repo}
}

func (s *EquityService) GetMtfPosition(ctx context.Context, req *pb.PositionRequest) (*pb.PositionResponse, error) {
	fmt.Println("Read Positions", req.GetEpbClmMtchAccnt())

	positions, err := s.repo.GetEquityPositions(req.GetEpbClmMtchAccnt())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Equity positions not found")
			return nil, errors.New("equity positions not found")
		}
		fmt.Println("Error querying equity positions:", err)
		return nil, err
	}

	pbPositions := make([]*pb.EquityPosition, len(positions))
	for i, pos := range positions {
		pbPositions[i] = &pb.EquityPosition{
			EpbClmMtchAccnt:      pos.EpbClmMtchAccnt,
			EpbXchngCd:           pos.EpbXchngCd,
			EpbXchngSgmntCd:      pos.EpbXchngSgmntCd,
			EpbXchngSgmntSttlmnt: pos.EpbXchngSgmntSttlmnt,
			EpbStckCd:            pos.EpbStckCd,
			EpbOrgnlPstnQty:      pos.EpbOrgnlPstnQty,
			EpbRate:              pos.EpbRate,
			EpbOrgnlAmtPayble:    pos.EpbOrgnlAmtPayble,
			EpbOrgnlMrgnAmt:      pos.EpbOrgnlMrgnAmt,
			EpbSellQty:           pos.EpbSellQty,
			EpbCvrOrdQty:         pos.EpbCvrOrdQty,
			EpbNetMrgnAmt:        pos.EpbNetMrgnAmt,
			EpbNetAmtPayble:      pos.EpbNetAmtPayble,
			EpbNetPstnQty:        pos.EpbNetPstnQty,
			EpbCtdQty:            pos.EpbCtdQty,
			EpbPstnStts:          pos.EpbPstnStts,
			EpbLpcCalcStts:       pos.EpbLpcCalcStts,
			EpbSqroffMode:        pos.EpbSqroffMode,
			EpbPstnTrdDt:         pos.EpbPstnTrdDt,
			EpbMtmPrcsFlg:        pos.EpbMtmPrcsFlg,
			EpbLastMdfcnDt:       pos.EpbLastMdfcnDt,
			EpbInsDate:           pos.EpbInsDate,
			EpbCloseDate:         pos.EpbCloseDate,
			EpbSysFailFlg:        pos.EpbSysFailFlg,
			EpbLastPymntDt:       pos.EpbLastPymntDt,
			EpbLpcCalcEndDt:      pos.EpbLpcCalcEndDt,
			EpbMtmCansq:          pos.EpbMtmCansq,
			EpbExpiryDt:          pos.EpbExpiryDt,
			EpbMinMrgn:           pos.EpbMinMrgn,
			EpbMrgnDbcrPrcsFlg:   pos.EpbMrgnDbcrPrcsFlg,
			EpbDpId:              pos.EpbDpId,
			EpbDpClntId:          pos.EpbDpClntId,
			EpbPledgeStts:        pos.EpbPledgeStts,
			EpbBtstNetMrgnAmt:    pos.EpbBtstNetMrgnAmt,
			EpbBtstMrgnBlckd:     pos.EpbBtstMrgnBlckd,
			EpbBtstMrgnDbcrFlg:   pos.EpbBtstMrgnDbcrFlg,
			EpbBtstSgmntCd:       pos.EpbBtstSgmntCd,
			EpbBtstStlmnt:        pos.EpbBtstStlmnt,
			EpbBtstCshBlckd:      pos.EpbBtstCshBlckd,
			EpbBtstSamBlckd:      pos.EpbBtstSamBlckd,
			EpbBtstCalcDt:        pos.EpbBtstCalcDt,
			EpbDbcrCalcDt:        pos.EpbDbcrCalcDt,
			EpbNsdlRefNo:         pos.EpbNsdlRefNo,
			EpbMrgnWithheldFlg:   pos.EpbMrgnWithheldFlg,
		}
	}

	response := &pb.PositionResponse{
		Equity: pbPositions,
	}

	return response, nil
}
