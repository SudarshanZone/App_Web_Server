package handlers

import (
    "context"
    "github.com/SudarshanZone/Web_Server/internal/grpc"
    pb "github.com/SudarshanZone/Fno_Open_Pos/generated"
    ob "github.com/SudarshanZone/Fno_Ord_Dtls/generated"
    cp "github.com/SudarshanZone/Commo_Open_Pos/generated"
    mt "github.com/SudarshanZone/MTF_OPEN_POS/generated" 
    eq "github.com/SudarshanZone/Equity_Main_Open_Pos/generated"
)


// Facade Design Pattern
type Facade interface {
    GetFNOPosition(account string) (*pb.FcpDetailListResponse, error)
    GetOrderDetails(orderID string) (*ob.OrderDetailsResponse, error)
    GetCommodityPositions(account string) (*cp.CCPResponse, error)
    GetMtfPositions(account string) (*mt.PositionResponse, error)
    GetEquityPosition(account string) (*eq.PositionResponse,error)
}

type facade struct {
    adapter grpc.Adapter
}

func NewFacade(adapter grpc.Adapter) Facade {
    return &facade{adapter: adapter}
}



// Implmentations of Methods 
func (f *facade) GetFNOPosition(account string) (*pb.FcpDetailListResponse, error) {
    return f.adapter.GetFNOPosition(context.Background(), account)
}

func (f *facade) GetOrderDetails(orderID string) (*ob.OrderDetailsResponse, error) {
    return f.adapter.GetOrderDetails(context.Background(), orderID)
}

func (f *facade) GetCommodityPositions(account string) (*cp.CCPResponse, error) {
    return f.adapter.GetCommodityPositions(context.Background(), account)
}

func (f *facade) GetMtfPositions(account string) (*mt.PositionResponse, error) {
    return f.adapter.GetMtfPositions(context.Background(), account)
}

func (f *facade) GetEquityPosition(account string)(*eq.PositionResponse,error){
    return f.adapter.GetEquityMainPosition(context.Background(),account)
}
