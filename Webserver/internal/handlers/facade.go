package handlers

import (
    "context"
    "github.com/SudarshanZone/Web_Server/internal/grpc"
    pb "github.com/SudarshanZone/Fno_Open_Pos/generated"
    ob "github.com/SudarshanZone/Fno_Ord_Dtls/generated"
    cp "github.com/SudarshanZone/Commo_Open_Pos/generated"
    mt "github.com/SudarshanZone/MTF_OPEN_POS/generated" // Import new package
)

type Facade interface {
    GetFNOPosition(account string) (*pb.FcpDetailListResponse, error)
    GetOrderDetails(orderID string) (*ob.OrderDetailsResponse, error)
    GetCommodityPositions(account string) (*cp.CCPResponse, error)
    GetMtfPositions(account string) (*mt.PositionResponse, error) // Add new method
}

type facade struct {
    adapter grpc.Adapter
}

func NewFacade(adapter grpc.Adapter) Facade {
    return &facade{adapter: adapter}
}

func (f *facade) GetFNOPosition(account string) (*pb.FcpDetailListResponse, error) {
    return f.adapter.GetFNOPosition(context.Background(), account)
}

func (f *facade) GetOrderDetails(orderID string) (*ob.OrderDetailsResponse, error) {
    return f.adapter.GetOrderDetails(context.Background(), orderID)
}

func (f *facade) GetCommodityPositions(account string) (*cp.CCPResponse, error) {
    return f.adapter.GetCommodityPositions(context.Background(), account)
}

func (f *facade) GetMtfPositions(account string) (*mt.PositionResponse, error) { // Add new method
    return f.adapter.GetMtfPositions(context.Background(), account)
}
