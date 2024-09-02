package models

// type OtpEquityPosition struct {
// 	OtpClmMtchAcct      string  // Claim match account
// 	OtpXchngCd          string  // Exchange code
// 	OtpXchngSgmntCd     string  // Exchange segment code
// 	OtpXchngSgmntSttlmnt int32   // Exchange segment settlement
// 	OtpStckCd           string  // Stock code
// 	OtpOrgnlPstnQty     int32   // Original position quantity
// 	OtpRate             float64 // Rate
// 	OtpOrgnlAmtPayble   float64 // Original amount payable
// 	OtpOrgnlMrgnAmt     float64 // Original margin amount
// 	OtpSellQty          int32   // Sell quantity
// 	OtpCvrOrdQty        int32   // Cover order quantity
// 	OtpNetMrgnAmt       float64 // Net margin amount
// 	OtpNetAmtPayble     float64 // Net amount payable
// 	OtpNetPstnQty       int32   // Net position quantity
// 	OtpCtdQty           int32   // Contracted quantity
// 	OtpPstnStts         string  // Position status
// 	OtpLpcCalcStts      string  // LPC calculation status
// 	OtpSqroffMode       string  // Square-off mode
// 	OtpPstnTrdDt        string  // Position trade date
// 	OtpMtmPrcsFlg       string  // MTM process flag
// 	OtpLastMdfcnDt      string  // Last modification date
// 	OtpInsDate          string  // Insertion date
// 	OtpCloseDate        string  // Close date
// 	OtpSysFailFlg       string  // System failure flag
// 	OtpLastPymntDt      string  // Last payment date
// 	OtpLpcCalcEndDt     string  // LPC calculation end date
// 	OtpMtmCansq         string  // MTM cancel square flag
// 	OtpExpiryDt         string  // Expiry date
// 	OtpMinMrgn          float64 // Minimum margin
// 	OtpMrgnDbcrPrcsFlg  string  // Margin debit/credit process flag
// 	OtpDpId             string  // DP ID
// 	OtpDpClntId         string  // DP client ID
// 	OtpPledgeStts       string  // Pledge status
// 	OtpBtstNetMrgnAmt   float64 // BTST net margin amount
// 	OtpBtstMrgnBlckd    float64 // BTST margin blocked
// 	OtpBtstMrgnDbcrFlg  string  // BTST margin debit/credit flag
// 	OtpBtstSgmntCd      string  // BTST segment code
// 	OtpBtstStlmnt       int32   // BTST settlement
// 	OtpBtstCshBlckd     float64 // BTST cash blocked
// 	OtpBtstSamBlckd     float64 // BTST sample blocked
// 	OtpBtstCalcDt       string  // BTST calculation date
// 	OtpDbcrCalcDt       string  // Debit calculation date
// 	OtpNsdlRefNo        string  // NSDL reference number
// 	OtpMrgnWithheldFlg  string  // Margin withheld flag
// }
type OtpEquityPosition struct {
	Otp_clm_mtch_acct       string
	Otp_xchng_cd            string
	Otp_xchng_sgmnt_cd      string
	Otp_xchng_sgmnt_sttlmnt int32
	Otp_stck_cd             string
	Otp_flw                 string
	Otp_qty                 int64
	Otp_cnvrt_dlvry_qty     int64
	Otp_cvrd_qty            int64
	Otp_rt                  string
	Otp_mrgn_amt            float64
	Otp_trd_val             float64
	Otp_rmrks               string
	Otp_xfer_mrgn_stts      string
	Otp_sell_opn_prccsd     string
	Otp_buy_opn_prccsd      string
	Otp_mrgn_sqroff_mode    string
	Otp_em_trdsplt_prcs_flg string
	Otp_mtm_flg             string
	Otp_mtm_cansq           string
	Otp_eos_can             string
	Otp_trgr_prc            float64
	Otp_16_trgr_prc         float64
	Otp_min_mrgn            float64
	Otp_t5_trdsplt_prcs_flg string
}