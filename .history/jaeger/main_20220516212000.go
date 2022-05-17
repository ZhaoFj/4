package main

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func main() {
	jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
		},
	}
}
