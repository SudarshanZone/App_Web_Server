package service

import (
	"context"
	"fmt"
	pb "github.com/SudarshanZone/Fno_Ord_Dtls/generated"
	"github.com/SudarshanZone/Fno_Ord_Dtls/internal/repository"
)

type OrderDetailsService struct {
	Repo repository.OrderDetailsRepository
	pb.UnimplementedOrderDetailsServiceServer
}

//getOrderDetails
func (s *OrderDetailsService) GetOrderDetails(ctx context.Context, req *pb.OrderDetailsRequest) (*pb.OrderDetailsResponse, error) {
	orderDetails, err := s.Repo.GetOrderDetailsByClaimMatchAccount(req.FOD_CLM_MTCH_ACCNT)
	if err != nil {
		return nil, fmt.Errorf("failed to get order details: %w", err)
	}

	var ordDetails []*pb.OrdDetail

	for _, ord := range orderDetails {
		ordDtls := &pb.OrdDetail{
			ContractDescriptor: ord.ContractDescriptor,
			VTCDate:            ord.VTCDate,
			BuySell:            ord.BuySell,
			Quantity:           ord.Quantity,
			Status:             ord.Status,
			OrderPrice:         ord.OrderPrice,
			Open:               ord.Open,
		}
		ordDetails = append(ordDetails, ordDtls)
	}

	return &pb.OrderDetailsResponse{
		OrdDetails: ordDetails,
	}, nil
}
