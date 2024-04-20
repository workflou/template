package main

import (
	"context"
	"embed"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	//go:embed static
	staticFS embed.FS
)

func newServer() *echo.Echo {
	e := echo.New()

	e.StaticFS("/static", echo.MustSubFS(staticFS, "static"))

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("requestID", v.RequestID),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	return e
}
