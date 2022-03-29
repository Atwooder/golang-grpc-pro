package global

import (
	"gin-grpc/pkg/config"
	"gin-grpc/pkg/tracer"
	"github.com/opentracing/opentracing-go"
	"log"
)

var (
	Tracer opentracing.Tracer
	Viper  *config.Setting
)

func Init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	//err = setupViper()
	//if err != nil {
	//	log.Fatalf("init.setupViper err: %v", err)
	//}
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("tour-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	Tracer = jaegerTracer
	return nil
}

func setupViper() error {
	viper, err := config.NewViper()
	if err != nil {
		return err
	}
	Viper = viper
	return nil
}
