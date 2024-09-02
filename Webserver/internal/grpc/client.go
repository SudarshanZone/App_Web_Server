package grpc

import (
    "log"
    "google.golang.org/grpc"
    pb "github.com/SudarshanZone/Fno_Open_Pos/generated" // Fno_open_Pos
    ob "github.com/SudarshanZone/Fno_Ord_Dtls/generated" //Fno_Ord_dtls
    cp "github.com/SudarshanZone/Commo_Open_Pos/generated" //Commodities_Open_Pos
    mt "github.com/SudarshanZone/MTF_OPEN_POS/generated"  //Mtf_Open_Pos
    eq "github.com/SudarshanZone/Equity_Main_Open_Pos/generated"
)

type Client struct {
    conn1                        *grpc.ClientConn
    conn2                        *grpc.ClientConn
    conn3                        *grpc.ClientConn
    conn4                        *grpc.ClientConn 
    conn5                        *grpc.ClientConn
    FnoPositionServiceClient     pb.FnoPositionServiceClient
    OrderDetailsServiceClient    ob.OrderDetailsServiceClient
    CommodityPositionServiceClient cp.CCPServiceClient
    MtfPositionServiceClient     mt.MtfPositionServiceClient 
    EquityMainServiceClient      eq.PositionServiceClient
}

func NewClient(address1, address2, address3, address4 ,address5 string) (Client, error) {
    conn1, err := grpc.Dial(address1, grpc.WithInsecure())
    if err != nil {
        return Client{}, err
    }
    conn2, err := grpc.Dial(address2, grpc.WithInsecure())
    if err != nil {
        return Client{}, err
    }
    conn3, err := grpc.Dial(address3, grpc.WithInsecure())
    if err != nil {
        return Client{}, err
    }
    conn4, err := grpc.Dial(address4, grpc.WithInsecure()) 
    if err != nil {
        return Client{}, err
    }

    conn5,err := grpc.Dial(address5,grpc.WithInsecure())
    if err != nil{
        return Client{},err
    }

    return Client{
        conn1:                        conn1,
        conn2:                        conn2,
        conn3:                        conn3,
        conn4:                        conn4, 
        conn5:                        conn5,   
        FnoPositionServiceClient:     pb.NewFnoPositionServiceClient(conn1),
        OrderDetailsServiceClient:    ob.NewOrderDetailsServiceClient(conn2),
        CommodityPositionServiceClient: cp.NewCCPServiceClient(conn3),
        MtfPositionServiceClient:     mt.NewMtfPositionServiceClient(conn4), 
        EquityMainServiceClient:      eq.NewPositionServiceClient(conn5),
    }, nil
}

func (c *Client) Close() {
    if err := c.conn1.Close(); err != nil {
        log.Printf("Failed to close gRPC connection 1: %v", err)
    }
    if err := c.conn2.Close(); err != nil {
        log.Printf("Failed to close gRPC connection 2: %v", err)
    }
    if err := c.conn3.Close(); err != nil {
        log.Printf("Failed to close gRPC connection 3: %v", err)
    }
    if err := c.conn4.Close(); err != nil { 
        log.Printf("Failed to close gRPC connection 4: %v", err)
    }
    if err := c.conn5.Close(); err != nil { 
        log.Printf("Failed to close gRPC connection 5: %v", err)
    }
}
