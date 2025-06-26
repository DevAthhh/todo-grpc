package logger

import (
	"errors"
	"log/slog"
	"os"
)

func Load(env string) (*slog.Logger, error) {
	var log *slog.Logger

	switch env {
	case "dev", "local":
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		log = slog.New(handler)
	case "prod":
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		log = slog.New(handler)
	default:
		return nil, errors.New("unknown env")
	}
	return log, nil
}
