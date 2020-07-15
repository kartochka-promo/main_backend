package server

import (
	"2020_1_drop_table/internal/microservices/mail"
	mailProto "2020_1_drop_table/internal/microservices/mail/delivery/grpc/proto"
	proto "2020_1_drop_table/internal/microservices/mail/delivery/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	mailUseCase mail.Usecase
}

func NewMailServerGRPC(gserver *grpc.Server, mailUCase mail.Usecase) {
	mailServer := &server{
		mailUseCase: mailUCase,
	}
	mailProto.RegisterMailGRPCServer(gserver, mailServer)
	reflection.Register(gserver)
}

func (s *server) SendEmail(ctx context.Context,
	params *mailProto.EmailParams) (*proto.Empty, error) {

	err := s.mailUseCase.SendEmail(ctx, params.Recipient, params.TemplateName, params.EmailContext)

	return &proto.Empty{}, err
}
