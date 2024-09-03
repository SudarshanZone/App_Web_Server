package models


type OrderDetails struct {
    OrdClmMtchAccnt *string    `gorm:"column:Ord_Clm_Mtch_Accnt" json:"ord_clm_mtch_accnt"`
    OrdStckCd       *string    `gorm:"column:Ord_Stck_Cd" json:"ord_stck_cd"`
    OrdOrdrDt       *string `gorm:"column:Ord_Ordr_Dt" json:"ord_ordr_dt"`
    OrdOrdrFlw      *string    `gorm:"column:Ord_Ordr_Flw" json:"ord_ordr_flw"`
    OrdOrdrQty      *int       `gorm:"column:Ord_Ordr_Qty" json:"ord_ordr_qty"`
    OrdLmtRt        *float64   `gorm:"column:Ord_Lmt_Rt" json:"ord_lmt_rt"`
    OrdOrdrStts     *string    `gorm:"column:Ord_Ordr_Stts" json:"ord_ordr_stts"`
}
