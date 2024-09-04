package models

// type FnoPosition struct {
// 	Contract          string  //`gorm:"column:FFO_CONTRACT"`
// 	Position          string  //`gorm:"column:FFO_PSTN"`
// 	TotalQty          int     //`gorm:"column:FFO_QTY"`
// 	AvgCostPrice      float64 //`gorm:"column:FFO_AVG_PRC"`
// 	ExchangeCode      string  //`gorm:"column:FCP_XCHNG_CD"`
// 	BuyQty            int     //`gorm:"column:FCP_IBUY_QTY"`
// 	ClaimMatchAccount string  //`gorm:"column:FCP_CLM_MTCH_ACCNT"`
// 	ProductType       string  //`gorm:"column:FCP_PRDCT_TYP"`
// 	IndexStock        string  //`gorm:"column:FCP_INDSTK"`
// 	Underlying        string  //`gorm:"column:FCP_UNDRLYNG"`
// 	ExpiryDate        string  //`gorm:"column:FCP_EXPRY_DT"`
// 	ExerciseType      string  //`gorm:"column:FCP_EXER_TYP"`
// 	OptionType        string  //`gorm:"column:FCP_OPT_TYP"`
// 	StrikePrice       float64 //`gorm:"column:FCP_STRK_PRC"`
// 	UCCCode           string //`gorm:"column:FCP_UCC_CD"`
// 	OpenPstnFlow      string  //`gorm:"column:FCP_OPNPSTN_FLW"`
// }
type FnoPosition struct {
    Contract          string  `json:"contract" gorm:"column:Contract"`
    Position          string  `json:"position" gorm:"column:Position"`
    TotalQty          int     `json:"totalQty" gorm:"column:TotalQty"`
    AvgCostPrice      float64 `json:"avgCostPrice" gorm:"column:AvgCostPrice"`
    ExchangeCode      string  `json:"exchangeCode" gorm:"column:ExchangeCode"`
    BuyQty            int     `json:"buyQty" gorm:"column:BuyQty"`
    ClaimMatchAccount string  `json:"claimMatchAccount" gorm:"column:ClaimMatchAccount"`
    ProductType       string  `json:"productType" gorm:"column:ProductType"`
    IndexStock        string  `json:"indexStock" gorm:"column:IndexStock"`
    Underlying        string  `json:"underlying" gorm:"column:Underlying"`
    ExpiryDate        string  `json:"expiryDate" gorm:"column:ExpiryDate"`
    ExerciseType      string  `json:"exerciseType" gorm:"column:ExerciseType"`
    OptionType        string  `json:"optionType" gorm:"column:OptionType"`
    StrikePrice       float64 `json:"strikePrice" gorm:"column:StrikePrice"`
    UCCCode           string  `json:"uccCode" gorm:"column:UccCode"`
    OpenPstnFlow      string  `json:"openPstnFlow" gorm:"column:OpenPstnFlow"`
}
