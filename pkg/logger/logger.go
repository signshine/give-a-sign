package logger

import (
	"log/slog"
	"os"

	"github.com/google/uuid"
)

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	
	return slog.New(slog.NewJSONHandler(os.Stdout, opts)).
		With("trace_id", uuid.NewString())
}