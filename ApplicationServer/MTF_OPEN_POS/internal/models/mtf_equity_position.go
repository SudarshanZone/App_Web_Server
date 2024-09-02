package models

type MtfEquityPosition struct {
	EpbClmMtchAccnt      string  // Claim match account
	EpbXchngCd           string  // Exchange code
	EpbXchngSgmntCd      string  // Exchange segment code
	EpbXchngSgmntSttlmnt int32   // Exchange segment settlement
	EpbStckCd            string  // Stock code
	EpbOrgnlPstnQty      int32   // Original position quantity
	EpbRate              float64 // Rate
	EpbOrgnlAmtPayble    float64 // Original amount payable
	EpbOrgnlMrgnAmt      float64 // Original margin amount
	EpbSellQty           int32   // Sell quantity
	EpbCvrOrdQty         int32   // Cover order quantity
	EpbNetMrgnAmt        float64 // Net margin amount
	EpbNetAmtPayble      float64 // Net amount payable
	EpbNetPstnQty        int32   // Net position quantity
	EpbCtdQty            int32   // Contracted quantity
	EpbPstnStts          string  // Position status
	EpbLpcCalcStts       string  // LPC calculation status
	EpbSqroffMode        string  // Square-off mode
	EpbPstnTrdDt         string  // Position trade date
	EpbMtmPrcsFlg        string  // MTM process flag
	EpbLastMdfcnDt       string  // Last modification date
	EpbInsDate           string  // Insertion date
	EpbCloseDate         string  // Close date
	EpbSysFailFlg        string  // System failure flag
	EpbLastPymntDt       string  // Last payment date
	EpbLpcCalcEndDt      string  // LPC calculation end date
	EpbMtmCansq          string  // MTM cancel square flag
	EpbExpiryDt          string  // Expiry date
	EpbMinMrgn           float64 // Minimum margin
	EpbMrgnDbcrPrcsFlg   string  // Margin debit/credit process flag
	EpbDpId              string  // DP ID
	EpbDpClntId          string  // DP client ID
	EpbPledgeStts        string  // Pledge status
	EpbBtstNetMrgnAmt    float64 // BTST net margin amount
	EpbBtstMrgnBlckd     float64 // BTST margin blocked
	EpbBtstMrgnDbcrFlg   string  // BTST margin debit/credit flag
	EpbBtstSgmntCd       string  // BTST segment code
	EpbBtstStlmnt        int32   // BTST settlement
	EpbBtstCshBlckd      float64 // BTST cash blocked
	EpbBtstSamBlckd      float64 // BTST sample blocked
	EpbBtstCalcDt        string  // BTST calculation date
	EpbDbcrCalcDt        string  // Debit calculation date
	EpbNsdlRefNo         string  // NSDL reference number
	EpbMrgnWithheldFlg   string  // Margin withheld flag
}
