package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	e := echo.New()

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, Echo!"})
	})

	logger := initLogger()
	defer logger.Sync()

	logger = logger.With(
		zap.String("service", "my-service"),
		zap.String("environment", "development"),
	)

	go func() {
		for {
			logger.Info("Application running",
				zap.Int("count", 1),
				zap.String("event_type", "status"),
			)
			logger.Error("Example error",
				zap.String("error_code", "ERR001"),
				zap.String("event_type", "error"),
			)
			time.Sleep(5 * time.Second)
		}
	}()

	e.Logger.Fatal(e.Start(":8080"))
}

func initLogger() *zap.Logger {
	config := zap.Config{
		Encoding:      "json",
		Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:   []string{"stdout", "./logs/application.log"},
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "level"

	logger, _ := config.Build()
	return logger
}
