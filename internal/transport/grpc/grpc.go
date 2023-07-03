package grpc

import (
	"context"
	"log"

	"github.com/fishkaoff/monitor-system-proto-files/proto"
	"google.golang.org/grpc"
)

func NewGRPCClient(remoteAddr string) (proto.StorageClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewStorageClient(conn)
	return c, nil
}

type GRPCStruct struct {
	svc proto.StorageClient
	ctx context.Context
}

func NewGRPCStruct(svc proto.StorageClient, ctx context.Context) *GRPCStruct {
	return &GRPCStruct{svc: svc, ctx: ctx}
}

func (gc *GRPCStruct) SaveUrl(chatID int64, site string) string {
	resp, err := gc.svc.SaveUrl(gc.ctx, &proto.SaveUrlRequest{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}

func (gc *GRPCStruct) DeleteUrl(chatID int64, site string) string {
	resp, err := gc.svc.DeleteUrl(gc.ctx, &proto.DeleteUrlRequest{ChatID: chatID, Site: site})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}

func (gc *GRPCStruct) Get(chatID int64) []string {
	resp, err := gc.svc.GetUrl(gc.ctx, &proto.GetUrlRequest{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}


func (gc *GRPCStruct) SaveUser(chatID int64, token string) string {
	resp, err := gc.svc.SaveUser(gc.ctx, &proto.SaveUserRequest{ChatID: chatID, Token: token})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}