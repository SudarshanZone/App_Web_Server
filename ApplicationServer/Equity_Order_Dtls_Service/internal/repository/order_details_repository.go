package repository

import (
    "fmt"
    "gorm.io/gorm"
    "github.com/SudarshanZone/Equity_Ord_Dtls/internal/models" 
)

type OrderDetailsRepository struct {
    Db *gorm.DB
}

func (r *OrderDetailsRepository) GetOrderDetails(ordClmMtchAccnt string) ([]*models.OrderDetails, error) {
    var orderDetailsList []*models.OrderDetails

    // Use raw SQL query
    query := `SELECT Ord_Clm_Mtch_Accnt, Ord_Stck_Cd, Ord_Ordr_Dt, Ord_Ordr_Flw, Ord_Ordr_Qty, Ord_Lmt_Rt, Ord_Ordr_Stts 
              FROM ORD_ORDR_DTLS WHERE Ord_Clm_Mtch_Accnt = ?`

    rows, err := r.Db.Raw(query, ordClmMtchAccnt).Rows()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var od models.OrderDetails
        if err := rows.Scan(
            &od.OrdClmMtchAccnt,
            &od.OrdStckCd,
            &od.OrdOrdrDt,
            &od.OrdOrdrFlw,
            &od.OrdOrdrQty,
            &od.OrdLmtRt,
            &od.OrdOrdrStts,
        ); err != nil {
            return nil, err
        }
        orderDetailsList = append(orderDetailsList, &od)
    }

    if len(orderDetailsList) == 0 {
        return nil, gorm.ErrRecordNotFound
    }

    fmt.Printf("Fetched order details: %+v\n", orderDetailsList)
    return orderDetailsList, nil
}
