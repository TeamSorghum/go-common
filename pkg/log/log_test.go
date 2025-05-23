package log_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/sainnhe/go-common/pkg/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	otellog "go.opentelemetry.io/otel/sdk/log"
)

func TestLog_NewLogger(t *testing.T) {
	t.Parallel()

	// Setup otel logger provider
	logExporter, err := stdoutlog.New()
	if err != nil {
		t.Fatal(err)
	}
	loggerProvider := otellog.NewLoggerProvider(
		otellog.WithProcessor(otellog.NewBatchProcessor(logExporter)),
	)
	global.SetLoggerProvider(loggerProvider)
	defer func() {
		if err := loggerProvider.Shutdown(context.Background()); err != nil {
			t.Fatal(err)
		}
	}()

	// Test global logger
	log.GetGlobalLogger().Debug("OTel logger provider has been initialized.")

	// Test new logger
	log.NewLogger("test").Debug("Start testing different configurations.")

	const pathPrefix = "/tmp/sainnhe-go-common-test"

	tests := []struct {
		name        string
		cfg         *log.Config
		expectError bool
	}{
		{
			"debug",
			&log.Config{
				Type:  "light",
				Level: "debug",
			},
			false,
		},
		{
			"info",
			&log.Config{
				Type:  "light",
				Level: "info",
			},
			false,
		},
		{
			"warn",
			&log.Config{
				Type:  "light",
				Level: "warn",
			},
			false,
		},
		{
			"error",
			&log.Config{
				Type:  "light",
				Level: "error",
			},
			false,
		},
		{
			"local",
			&log.Config{
				Type:  "local",
				Level: "debug",
				Local: log.LocalConfig{
					Path:       pathPrefix + "/testlog",
					MaxSizeMB:  1,
					MaxBackups: 3,
				},
			},
			false,
		},
		{
			"otel",
			&log.Config{
				Type: "otel",
			},
			false,
		},
		{
			"default type and level",
			&log.Config{},
			false,
		},
		{
			"unsupported level",
			&log.Config{
				Level: "nil",
			},
			true,
		},
		{
			"nil config",
			nil,
			true,
		},
		{
			"unsupported level",
			&log.Config{
				Level: "nil",
			},
			true,
		},
		{
			"unsupported type",
			&log.Config{
				Type: "nil",
			},
			true,
		},
	}

	output := func(logger *slog.Logger, msg string, attrs []any) {
		logger.Debug(msg, attrs...)
		logger.Info(msg, attrs...)
		logger.Warn(msg, attrs...)
		logger.Error(msg, attrs...)
	}

	msg := "Test"
	attrs := []any{"attr1", "attr1", "attr2", "attr2"}
	otelAttrs := []attribute.KeyValue{
		{Key: attribute.Key("otelAttr1"), Value: attribute.BoolValue(true)},
		{Key: attribute.Key("otelAttr2"), Value: attribute.Int64Value(10)},
	}

	for _, tt := range tests { // nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			// Set global config
			cleanup, err := log.SetGlobalConfig(tt.cfg)
			defer cleanup()
			if tt.expectError != (err != nil) {
				t.Fatalf("Expect error = %t, got %+v", tt.expectError, err)
			}
			if err != nil {
				return
			}

			// Get global logger
			logger := log.GetGlobalLogger()

			// Init otel attributes
			if log.WithOTelAttrs(nil) != nil {
				t.Fatal("Expect nil logger")
			}
			logger = log.WithOTelAttrs(logger, otelAttrs...)

			// Handle output
			output(logger, msg, attrs)

			// Cleanup
			cleanup()
		})
	}
}
