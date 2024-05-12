/*
Package logger implements the logic to generate different levels of logs within the application using custom tags field.
*/
package logger

import (
	"context"
	"log/slog"
	"os"
)

var logger *slog.Logger

type (
	MpConciliationKey struct{}
	Tags              map[string]interface{}
)

func SetupLogger() {
	options := slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelInfo,
	}

	handler := slog.NewTextHandler(os.Stdout, &options)
	logger = slog.New(handler)
	slog.SetDefault(logger)
}

func Info(ctx context.Context, msg string, tags Tags) {
	tags["action"] = ctx.Value(MpConciliationKey{})
	logger.With(getFields(tags)...).InfoContext(ctx, msg)
}

func Error(ctx context.Context, msg string, tags Tags) {
	tags["action"] = ctx.Value(MpConciliationKey{})
	logger.With(getFields(tags)...).ErrorContext(ctx, msg)
}

func getFields(tags map[string]interface{}) []any {
	fields := make([]any, len(tags))
	index := 0

	for key, value := range tags {
		fields[index] = slog.Any(key, value)
		index++
	}

	return fields
}
