package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
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
			LocalAgentHostPort: "0.0.0.0:6831",
		},
		ServiceName: "hantaMall",
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}

	parantSpan := tracer.StartSpan("order_web")

	cartSpan := tracer.StartSpan("cart_srv", opentracing.ChildOf(parantSpan.Context()))
	time.Sleep(time.Second * 1)
	cartSpan.Finish()

	producntSpan := tracer.StartSpan("product_srv", opentracing.ChildOf(parantSpan.Context()))
	time.Sleep(time.Second * 1)
	producntSpan.Finish()

}
