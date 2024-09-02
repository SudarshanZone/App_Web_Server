package models
type OrderDetails struct {
    ContractDescriptor string  `gorm:"column:ContractDescriptor"`
    VTCDate            string  `gorm:"column:VTCDate"`
    BuySell            string  `gorm:"column:BuySell"`
    Quantity           int32   `gorm:"column:Quantity"`
    Status             string  `gorm:"column:Status"`
    OrderPrice         float32 `gorm:"column:OrderPrice"`
    Open               string  `gorm:"column:Open"`
}
