package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type GetOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) Execute() ([]*GetOrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	ordersDTO := make([]*GetOrderOutputDTO, len(orders))
	for idx, order := range orders {
		ordersDTO[idx] = &GetOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return ordersDTO, nil
}
