package grpc

import (
	"context"
	"log"

	"github.com/fishkaoff/telegram-client/proto"
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

func (gc *GRPCStruct) Save(chatID int64, site string) string {
	resp, err := gc.svc.Save(gc.ctx, &proto.SaveRequest{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}

func (gc *GRPCStruct) Delete(chatID int64, site string) string {
	resp, err := gc.svc.Delete(gc.ctx, &proto.DeleteRequest{ChatID: chatID, Site: site})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}

func (gc *GRPCStruct) Get(chatID int64) []string {
	resp, err := gc.svc.Get(gc.ctx, &proto.GetRequest{ChatID: chatID})
	if err != nil {
		log.Fatal(err)
	}

	return resp.Message
}
