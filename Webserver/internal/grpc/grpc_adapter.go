package grpc

import (
	"context"
	"fmt"
	cp "github.com/SudarshanZone/Commo_Open_Pos/generated" //cmmd
	pb "github.com/SudarshanZone/Fno_Open_Pos/generated"   //fno
	ob "github.com/SudarshanZone/Fno_Ord_Dtls/generated"   //fno
	mt "github.com/SudarshanZone/MTF_OPEN_POS/generated"   //MTF
	eq "github.com/SudarshanZone/Equity_Main_Open_Pos/generated"
)


// Adapter 
type Adapter interface {

	//Fno_Open_Pos
	GetFNOPosition(ctx context.Context, account string) (*pb.FcpDetailListResponse, error)

	//Fno_Ord_Dtls
	GetOrderDetails(ctx context.Context, orderID string) (*ob.OrderDetailsResponse, error)

	// Commodities_Ope_Pos
	GetCommodityPositions(ctx context.Context, account string) (*cp.CCPResponse, error)

	//MTF Open Positions
	GetMtfPositions(ctx context.Context, account string) (*mt.PositionResponse, error)

	//Equity_Main positions
	GetEquityMainPosition(ctx context.Context,account string)(*eq.PositionResponse,error)
}

type grpcAdapter struct {
	fnoClient   pb.FnoPositionServiceClient
	orderClient ob.OrderDetailsServiceClient
	commClient  cp.CCPServiceClient
	mtfClient   mt.MtfPositionServiceClient
	eqClient    eq.PositionServiceClient
}



// Adapter Method Implmentation
func (a *grpcAdapter) GetFNOPosition(ctx context.Context, account string) (*pb.FcpDetailListResponse, error) {
	req := &pb.FnoPositionRequest{
		FCP_CLM_MTCH_ACCNT: account,
	}
	resp, err := a.fnoClient.GetFNOPosition(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch FNO positions: %w", err)
	}
	return resp, nil
}


func (a *grpcAdapter) GetOrderDetails(ctx context.Context, orderID string) (*ob.OrderDetailsResponse, error) {
	req := &ob.OrderDetailsRequest{
		FOD_CLM_MTCH_ACCNT: orderID,
	}
	return a.orderClient.GetOrderDetails(ctx, req)
}

func (a *grpcAdapter) GetCommodityPositions(ctx context.Context, account string) (*cp.CCPResponse, error) {
	req := &cp.CCPRequest{
		CCP_CLM_MTCH_ACCNT: account,
	}
	resp, err := a.commClient.GetCCPData(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commodity positions: %w", err)
	}
	return resp, nil
}

func (a *grpcAdapter) GetMtfPositions(ctx context.Context, account string) (*mt.PositionResponse, error) { // Add new method
	req := &mt.PositionRequest{
		EpbClmMtchAccnt: account,
	}
	resp, err := a.mtfClient.GetMtfPosition(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch MTF positions: %w", err)
	}
	return resp, nil
}

func(a *grpcAdapter) GetEquityMainPosition(ctx context.Context,account string)(*eq.PositionResponse,error){
	req := &eq.PositionRequest{
		OtpClmMtchAcct: account,
	}
	resp,err := a.eqClient.GetPosition(ctx,req)
	if err != nil{
		return nil,fmt.Errorf("failed to fetch Equity Main positions: %w",err)
	}
	return resp,err
}


func NewGrpcAdapter(fnoClient pb.FnoPositionServiceClient, orderClient ob.OrderDetailsServiceClient, commClient cp.CCPServiceClient, mtfClient mt.MtfPositionServiceClient,eqClient eq.PositionServiceClient) Adapter {
	return &grpcAdapter{
		fnoClient:   fnoClient,
		orderClient: orderClient,
		commClient:  commClient,
		mtfClient:   mtfClient,
		eqClient: eqClient,
	}
}