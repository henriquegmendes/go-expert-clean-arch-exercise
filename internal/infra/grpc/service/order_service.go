package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrdersUseCase   usecase.GetOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrdersUseCase usecase.GetOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrdersUseCase:   getOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrders(_ context.Context, _ *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	output, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	pbOrders := make([]*pb.CreateOrderResponse, len(output))
	for idx, orderDTO := range output {
		pbOrders[idx] = &pb.CreateOrderResponse{
			Id:         orderDTO.ID,
			Price:      float32(orderDTO.Price),
			Tax:        float32(orderDTO.Tax),
			FinalPrice: float32(orderDTO.FinalPrice),
		}
	}

	return &pb.GetOrdersResponse{
		Orders: pbOrders,
	}, nil
}
