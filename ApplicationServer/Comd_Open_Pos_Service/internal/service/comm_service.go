// ApplicationServer\Comd_Open_Pos_Service\internal\service\comm_service.go
package service

import (
	"context"
	"errors"
	"fmt"
	_"github.com/SudarshanZone/Commo_Open_Pos/internal/models"
	"github.com/SudarshanZone/Commo_Open_Pos/internal/repository"
	pb "github.com/SudarshanZone/Commo_Open_Pos/generated"
	"gorm.io/gorm"
)

type CCPService struct {
	pb.UnimplementedCCPServiceServer
	repo repository.CCPRepository
}

func NewCCPService(repo repository.CCPRepository) *CCPService {
	return &CCPService{repo: repo}
}

func (s *CCPService) GetCCPData(ctx context.Context, req *pb.CCPRequest) (*pb.CCPResponse, error) {
	fmt.Println("Read Positions", req.GetCCP_CLM_MTCH_ACCNT())

	positions, err := s.repo.GetCommodityPositions(req.GetCCP_CLM_MTCH_ACCNT())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Commodity positions not found")
			return nil, errors.New("commodity positions not found")
		}
		fmt.Println("Error querying commodity positions:", err)
		return nil, err
	}

	pbPositions := make([]*pb.CommodityPositions, len(positions))
	for i, pos := range positions {
		pbPositions[i] = &pb.CommodityPositions{
			CCP_CLM_MTCH_ACCNT:       pos.CCP_CLM_MTCH_ACCNT,
			CCP_XCHNG_CD:             pos.CCP_XCHNG_CD,
			CCP_UNDRLYNG:             pos.CCP_UNDRLYNG,
			CCP_PRDCT_TYP:            pos.CCP_PRDCT_TYP,
			CCP_INDSTK:               pos.CCP_INDSTK,
			CCP_EXPRY_DT:             pos.CCP_EXPRY_DT,
			CCP_EXER_TYP:             pos.CCP_EXER_TYP,
			CCP_STRK_PRC:             int64(pos.CCP_STRK_PRC),
			CCP_OPT_TYP:              pos.CCP_OPT_TYP,
			CCP_IBUY_QTY:             int64(pos.CCP_IBUY_QTY),
			CCP_IBUY_ORD_VAL:         pos.CCP_IBUY_ORD_VAL,
			CCP_ISELL_QTY:            int64(pos.CCP_ISELL_QTY),
			CCP_ISELL_ORD_VAL:        pos.CCP_ISELL_ORD_VAL,
			CCP_EXBUY_QTY:            int64(pos.CCP_EXBUY_QTY),
			CCP_EXBUY_ORD_VAL:        pos.CCP_EXBUY_ORD_VAL,
			CCP_EXSELL_QTY:           int64(pos.CCP_EXSELL_QTY),
			CCP_EXSELL_ORD_VAL:       pos.CCP_EXSELL_ORD_VAL,
			CCP_BUY_EXCTD_QTY:        int64(pos.CCP_BUY_EXCTD_QTY),
			CCP_SELL_EXCTD_QTY:       int64(pos.CCP_SELL_EXCTD_QTY),
			CCP_OPNPSTN_FLW:          pos.CCP_OPNPSTN_FLW,
			CCP_OPNPSTN_QTY:          int64(pos.CCP_OPNPSTN_QTY),
			CCP_OPNPSTN_VAL:          pos.CCP_OPNPSTN_VAL,
			CCP_EXRC_QTY:             int64(pos.CCP_EXRC_QTY),
			CCP_ASGND_QTY:            int64(pos.CCP_ASGND_QTY),
			CCP_OPT_PREMIUM:          pos.CCP_OPT_PREMIUM,
			CCP_MTM_OPN_VAL:          pos.CCP_MTM_OPN_VAL,
			CCP_IMTM_OPN_VAL:         pos.CCP_IMTM_OPN_VAL,
			CCP_EXTRMLOSS_MRGN_EXTRA: pos.CCP_EXTRMLOSS_MRGN_EXTRA,
			CCP_ADDNL_MRGN:           pos.CCP_ADDNL_MRGN,
			CCP_SPCL_MRGN:            pos.CCP_SPCL_MRGN,
			CCP_TNDR_MRGN:            pos.CCP_TNDR_MRGN,
			CCP_DLVRY_MRGN:           pos.CCP_DLVRY_MRGN,
			CCP_EXTRM_MIN_LOSS_MRGN:  pos.CCP_EXTRM_MIN_LOSS_MRGN,
			CCP_MTM_FLG:              pos.CCP_MTM_FLG,
			CCP_EXTRM_LOSS_MRGN:      pos.CCP_EXTRM_LOSS_MRGN,
			CCP_FLAT_VAL_MRGN:        pos.CCP_FLAT_VAL_MRGN,
			CCP_TRG_PRC:              pos.CCP_TRG_PRC,
			CCP_MIN_TRG_PRC:          pos.CCP_MIN_TRG_PRC,
			CCP_DEVOLMNT_MRGN:        pos.CCP_DEVOLMNT_MRGN,
			CCP_MTMSQ_ORDCNT:         int32(pos.CCP_MTMSQ_ORDCNT),
			CCP_AVG_PRC:              pos.CCP_AVG_PRC,
		}
	}

	return &pb.CCPResponse{COMMO: pbPositions}, nil
}
