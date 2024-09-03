package service

import (
	"context"
	"fmt"
	"github.com/SudarshanZone/Equity_Ord_Dtls/internal/repository"
	pb "github.com/SudarshanZone/Equity_Ord_Dtls/generated"
)

type OrderDetailsService struct {
	repo repository.OrderDetailsRepository
	pb.UnimplementedOrderServiceServer
}

func NewOrderDetailsService(repo repository.OrderDetailsRepository) pb.OrderServiceServer {
	return &OrderDetailsService{repo: repo}
}

func (s *OrderDetailsService) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	orderDetails, err := s.repo.GetOrderDetails(req.GetOrdClmMtchAccnt())
	if err != nil {
		return nil, fmt.Errorf("failed to get order details: %w", err)
	}

	var ordDtlsList []*pb.OrdOrdrDtls
	for _, od := range orderDetails {
		ordDtls := &pb.OrdOrdrDtls{
			OrdClmMtchAccnt: *od.OrdClmMtchAccnt,
			OrdStckCd:       *od.OrdStckCd,
			OrdOrdrDt:       *od.OrdOrdrDt,
			OrdOrdrFlw:      *od.OrdOrdrFlw,
			OrdOrdrQty:      int32(*od.OrdOrdrQty),
			OrdLmtRt:        *od.OrdLmtRt,
			OrdOrdrStts:     *od.OrdOrdrStts,
		}
		ordDtlsList = append(ordDtlsList, ordDtls)
	}

	return &pb.OrderResponse{OrdDtls: ordDtlsList}, nil
}
