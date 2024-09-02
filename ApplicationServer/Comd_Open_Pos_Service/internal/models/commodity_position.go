// internal/models/commodity_position.go
package models

type CommodityPosition struct {
	CCP_CLM_MTCH_ACCNT   string
	CCP_XCHNG_CD         string
	CCP_UNDRLYNG         string
	CCP_PRDCT_TYP        string
	CCP_INDSTK           string
	CCP_EXPRY_DT         string
	CCP_EXER_TYP         string
	CCP_STRK_PRC         float64
	CCP_OPT_TYP          string
	CCP_IBUY_QTY         int
	CCP_IBUY_ORD_VAL     float64
	CCP_ISELL_QTY        int
	CCP_ISELL_ORD_VAL    float64
	CCP_EXBUY_QTY        int
	CCP_EXBUY_ORD_VAL    float64
	CCP_EXSELL_QTY       int
	CCP_EXSELL_ORD_VAL   float64
	CCP_BUY_EXCTD_QTY    int
	CCP_SELL_EXCTD_QTY   int
	CCP_OPNPSTN_FLW      string
	CCP_OPNPSTN_QTY      int
	CCP_OPNPSTN_VAL      float64
	CCP_EXRC_QTY         int
	CCP_ASGND_QTY        int
	CCP_OPT_PREMIUM      float64
	CCP_MTM_OPN_VAL      float64
	CCP_IMTM_OPN_VAL     float64
	CCP_EXTRMLOSS_MRGN_EXTRA float64
	CCP_ADDNL_MRGN       float64
	CCP_SPCL_MRGN        float64
	CCP_TNDR_MRGN        float64
	CCP_DLVRY_MRGN       float64
	CCP_EXTRM_MIN_LOSS_MRGN  float64
	CCP_MTM_FLG          string
	CCP_EXTRM_LOSS_MRGN  float64
	CCP_FLAT_VAL_MRGN    float64
	CCP_TRG_PRC          float64
	CCP_MIN_TRG_PRC      float64
	CCP_DEVOLMNT_MRGN    float64
	CCP_MTMSQ_ORDCNT     int
	CCP_AVG_PRC          float64
}