package main

import (
	"fmt"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

func main() {
	c := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		Encoding:          "json",
		DisableCaller:     true,
		DisableStacktrace: true,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	logger := zapr.NewLogger(l)

	logger.Info("hoge")

	nl := logger.WithName("named")
	nl.Info("bar")

	nnl := nl.WithName("piyo")
	nnl.Info("baz")

	nnl.Error(fmt.Errorf("ERROR"), "error")
}
