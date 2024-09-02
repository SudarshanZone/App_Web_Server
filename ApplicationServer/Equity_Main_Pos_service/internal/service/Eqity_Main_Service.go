package service

import (
    "context"

    "github.com/SudarshanZone/Equity_Main_Open_Pos/internal/repository"
    pb "github.com/SudarshanZone/Equity_Main_Open_Pos/generated"
)

type PositionService struct {
    repo repository.PositionRepository
    pb.UnimplementedPositionServiceServer
}

func NewPositionService(repo repository.PositionRepository) *PositionService {
    return &PositionService{repo: repo}
}

func (s *PositionService) GetPosition(ctx context.Context, req *pb.PositionRequest) (*pb.PositionResponse, error) {
    positions, err := s.repo.GetEquityPositions(req.GetOtpClmMtchAcct())
    if err != nil {
        return nil, err
    }

    var pbPositions []*pb.EquityPosition
    for _, position := range positions {
        pbPositions = append(pbPositions, &pb.EquityPosition{
            OtpClmMtchAcct:       position.Otp_clm_mtch_acct,
            OtpXchngCd:            position.Otp_xchng_cd,
            OtpXchngSgmntCd:      position.Otp_xchng_sgmnt_cd,
            OtpXchngSgmntSttlmnt: position.Otp_xchng_sgmnt_sttlmnt,
            OtpStckCd:             position.Otp_stck_cd,
            OtpFlw:                 position.Otp_flw,
            OtpQty:                 position.Otp_qty,
            OtpCnvrtDlvryQty:     position.Otp_cnvrt_dlvry_qty,
            OtpCvrdQty:            position.Otp_cvrd_qty,
            OtpRt:                  position.Otp_rt,
            OtpMrgnAmt:            position.Otp_mrgn_amt,
            OtpTrdVal:             position.Otp_trd_val,
            OtpRmrks:               position.Otp_rmrks,
            OtpXferMrgnStts:      position.Otp_xfer_mrgn_stts,
            OtpSellOpnPrccsd:     position.Otp_sell_opn_prccsd,
            OtpBuyOpnPrccsd:      position.Otp_buy_opn_prccsd,
            OtpMrgnSqroffMode:    position.Otp_mrgn_sqroff_mode,
            OtpEmTrdspltPrcsFlg:  position.Otp_em_trdsplt_prcs_flg,
            OtpMtmFlg:             position.Otp_mtm_flg,
            OtpMtmCansq:          position.Otp_mtm_cansq,
            OtpEosCan:             position.Otp_eos_can,
            OtpTrgrPrc:            position.Otp_trgr_prc,
            Otp_16TrgrPrc:         float32(position.Otp_16_trgr_prc),
            OtpMinMrgn:           position.Otp_min_mrgn,
            OtpT5TrdspltPrcsFlg:  position.Otp_t5_trdsplt_prcs_flg,
        })
    }

    return &pb.PositionResponse{Equity: pbPositions}, nil
}
