package interfaces

import (
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/souushinn/cardGo/domain/repository"
	"google.golang.org/grpc"
)

type ServerParams struct {
	CardRepository repository.CardRepository
}

func NewServer(parmas ServerParams) *grpc.Server {
	options := make([]grpc.ServerOption, 0)
	options = append(options, grpc.UnaryInterceptor(
		middleware.ChainUnaryServer(
			interceptor.LoggerSupplyInterceptor(),
			interceptor.RecoveryInterceptor(),
			interceptor.LoggingInterceptro(),
			interceptor.ErrorHandleInterceptor(),
		)))
	server := grpc.NewServer(options...)
}
