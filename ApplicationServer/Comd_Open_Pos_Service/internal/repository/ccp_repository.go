// ApplicationServer\Comd_Open_Pos_Service\internal\repository\ccp_repository.go
package repository

import (
	"gorm.io/gorm"
	"github.com/SudarshanZone/Commo_Open_Pos/internal/models"
)

type CCPRepository interface {
	GetCommodityPositions(account string) ([]models.CommodityPosition, error)
}

type ccpRepository struct {
	db *gorm.DB
}

func NewCCPRepository(db *gorm.DB) CCPRepository {
	return &ccpRepository{db: db}
}

//GetCommodityPositions
func (r *ccpRepository) GetCommodityPositions(account string) ([]models.CommodityPosition, error) {
	query := `
		SELECT ccp_clm_mtch_accnt, ccp_xchng_cd, CONCAT(ccp_prdct_typ, '_', ccp_undrlyng, '_', ccp_expry_dt) AS ccp_undrlyng, 
			   ccp_prdct_typ, ccp_indstk, ccp_undrlyng, ccp_expry_dt, ccp_exer_typ, ccp_strk_prc, ccp_opt_typ, 
			   ccp_ibuy_qty, ccp_ibuy_ord_val, ccp_isell_qty, ccp_isell_ord_val, ccp_exbuy_qty, ccp_exbuy_ord_val, 
			   ccp_exsell_qty, ccp_exsell_ord_val, ccp_buy_exctd_qty, ccp_sell_exctd_qty, ccp_opnpstn_flw, 
			   ccp_opnpstn_qty, ccp_opnpstn_val, ccp_exrc_qty, ccp_asgnd_qty, ccp_opt_premium, ccp_mtm_opn_val, 
			   ccp_imtm_opn_val, ccp_extrmloss_mrgn, ccp_addnl_mrgn, ccp_spcl_mrgn, ccp_tndr_mrgn, 
			   ccp_dlvry_mrgn, ccp_extrm_min_loss_mrgn, ccp_mtm_flg, ccp_extrm_loss_mrgn, ccp_flat_val_mrgn, 
			   ccp_trg_prc, ccp_min_trg_prc, ccp_devolmnt_mrgn, ccp_mtmsq_ordcnt, ccp_avg_prc 
		FROM ccp_cod_spn_cntrct_pstn 
		WHERE ccp_clm_mtch_accnt = ? AND ccp_opnpstn_flw <> 'N'`

	var positions []models.CommodityPosition
	result := r.db.Raw(query, account).Scan(&positions)
	if result.Error != nil {
		return nil, result.Error
	}
	return positions, nil
}
