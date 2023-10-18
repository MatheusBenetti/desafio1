package grpc

import (
	"context"
	"github.com/MatheusBenetti/desafio1/application/grpc/pb"
	"github.com/MatheusBenetti/desafio1/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.CreateProduct(in.Name, in.Description, float64(in.Price))
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{Product: &pb.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       float32(product.Price),
	}}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context) (*pb.FindProductsResponse, error) {
	products, err := p.ProductUseCase.FindProducts()
	if err != nil {
		return nil, err
	}

	var productProtos []*pb.Product
	for _, p := range products {
		productProtos = append(productProtos, &pb.Product{
			// trocar os ID's por int
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       float32(p.Price),
		})
	}

	return &pb.FindProductsResponse{Products: productProtos}, nil
}

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
