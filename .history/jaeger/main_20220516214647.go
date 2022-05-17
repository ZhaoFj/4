package main

import (
	"time"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "1.15.44.176:6831",
		},
		ServiceName: "hantaMall",
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}

	span := tracer.StartSpan("account_web")
	defer span.Finish()
	time.Sleep(time.Second * 1)

}
