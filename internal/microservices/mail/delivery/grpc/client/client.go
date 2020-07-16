package client

import (
	mailProto "2020_1_drop_table/internal/microservices/mail/delivery/grpc/proto"
	"context"
	"google.golang.org/grpc"
)

type MailClient struct {
	client mailProto.MailGRPCClient
}

func NewMailClient(conn *grpc.ClientConn) MailClient {
	c := mailProto.NewMailGRPCClient(conn)
	return MailClient{
		client: c,
	}
}

func (m MailClient) SendEmail(ctx context.Context,
	recipient, templateName string, emailContext map[string]string) error {

	params := mailProto.EmailParams{
		Recipient:    recipient,
		TemplateName: templateName,
		EmailContext: emailContext,
	}

	_, err := m.client.SendEmail(ctx, &params)
	return err
}
