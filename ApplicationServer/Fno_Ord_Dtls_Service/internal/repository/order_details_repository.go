package repository

import (
	"fmt"
	"github.com/SudarshanZone/Fno_Ord_Dtls/internal/models"
	"gorm.io/gorm"
)

type OrderDetailsRepository struct {
	Db *gorm.DB
}

func (repo *OrderDetailsRepository) GetOrderDetailsByClaimMatchAccount(claimMatchAccount string) ([]models.OrderDetails, error) {
	var orderDetails []models.OrderDetails
	// query := `
	//     SELECT
	//         CASE
	//             WHEN FOD_UNDRLYNG IS NOT NULL THEN
	//                 CONCAT(FOD_UNDRLYNG, ' ', TO_CHAR(FOD_TRD_DT, 'DD-Mon-YYYY'))
	//             ELSE 'Unknown'
	//         END AS "ContractDescriptor",
	//         TO_CHAR(FOD_TRD_DT, 'DD-Mon-YYYY') AS "VTCDate",
	//         CASE
	//             WHEN FOD_ORDR_FLW = 'B' THEN 'BUY'
	//             WHEN FOD_ORDR_FLW = 'S' THEN 'SELL'
	//             ELSE ''
	//         END AS "BuySell",
	//         FOD_ORDR_TOT_QTY AS "Quantity",
	//         FOD_ORDR_STTS AS "Status",
	//         FOD_LMT_RT AS "OrderPrice",
	//         CASE
	//             WHEN FOD_ORDR_STTS = 'OPEN' THEN 'Yes'
	//             ELSE 'No'
	//         END AS "Open"
	//     FROM
	//         FOD_FO_ORDR_DTLS
	//     WHERE
	//         (FOD_ORDR_TOT_QTY != 0 OR FOD_LMT_RT IS NOT NULL)
	//         AND FOD_CLM_MTCH_ACCNT = ?;
	// `

	query := `SELECT
    CASE
        WHEN FOD.FOD_PRDCT_TYP = 'F' THEN
            'FUT-' || TRIM(FOD.FOD_UNDRLYNG) || '-' || TO_CHAR(FOD.FOD_EXPRY_DT, 'DD-Mon-YYYY')
        WHEN FOD.FOD_PRDCT_TYP = 'O' THEN
            'OPT-' || TRIM(FOD.FOD_UNDRLYNG) || '-' || TO_CHAR(FOD.FOD_EXPRY_DT, 'DD-Mon-YYYY') || '-' ||
            FOD.FOD_STRK_PRC || '-' ||
            CASE
                WHEN FOD.FOD_OPT_TYP = 'C' THEN 'CE'
                WHEN FOD.FOD_OPT_TYP = 'P' THEN 'PE'
                ELSE '*'
            END
        ELSE
            TRIM(FOD.FOD_UNDRLYNG) || '-' || TO_CHAR(FOD.FOD_EXPRY_DT, 'DD-Mon-YYYY')
    END AS "ContractDescriptor",
    TO_CHAR(FOD.FOD_TRD_DT, 'DD-Mon-YYYY') AS "VTCDate",
    CASE
        WHEN FOD.FOD_ORDR_FLW = 'B' THEN 'BUY'
        WHEN FOD.FOD_ORDR_FLW = 'S' THEN 'SELL'
        ELSE '*'
    END AS "Buy/Sell",
    FOD.FOD_ORDR_TOT_QTY AS "Quantity",
    FOD.FOD_ORDR_STTS AS "Status",
    (FOD.FOD_LMT_RT)  AS "OrderPrice",
    FCP.FCP_OPNPSTN_QTY AS "Open"
FROM
    FOD_FO_ORDR_DTLS FOD
LEFT JOIN
    FCP_FO_SPN_CNTRCT_PSTN FCP
ON
    FOD.fod_clm_mtch_accnt = FCP.fcp_clm_mtch_accnt 
WHERE
    (FOD.FOD_ORDR_TOT_QTY != 0 OR FOD.FOD_LMT_RT IS NOT NULL)
    AND FOD_CLM_MTCH_ACCNT = ?;
`
	err := repo.Db.Raw(query, claimMatchAccount).Scan(&orderDetails).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching order details: %w", err)
	}

	return orderDetails, nil
}
