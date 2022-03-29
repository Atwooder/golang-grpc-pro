package main

import (
	"context"
	loggerSystem "gin-grpc/pkg/log/system"
	"log"

	"gin-grpc/global"
	"gin-grpc/internal/middleware"
	"gin-grpc/pkg/tracer"
	pb "gin-grpc/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	global.Init()
}

func main() {
	loggerSystem.InfoWithFields("错误", loggerSystem.Fields{
		"error": "-----------err--------",
	})
	ctx := context.Background()
	newCtx := metadata.AppendToOutgoingContext(ctx, "eddycjy", "Go语言编程之旅")
	clientConn, err := GetClientConn(newCtx, "localhost:8004", []grpc.DialOption{grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware.UnaryContextTimeout(),
			middleware.ClientTracing(),
		),
	)})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(newCtx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

func setupTracer() error {
	jaegerTracer, closer, err := tracer.NewJaegerTracer("article-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	defer closer.Close()
	global.Tracer = jaegerTracer
	return nil
}
