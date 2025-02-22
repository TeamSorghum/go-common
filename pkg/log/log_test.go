package log_test

import (
	"context"
	"testing"

	loadconfig "github.com/teamsorghum/go-common/pkg/load_config"
	"github.com/teamsorghum/go-common/pkg/log"
	ctxutil "github.com/teamsorghum/go-common/pkg/util/ctx"
)

// nolint:paralleltest
func TestLog_Logger(t *testing.T) {
	defaultCfg, err := loadconfig.Load[log.Config](nil, loadconfig.TypeNil)
	if err != nil {
		t.Fatal(err)
	}
	slogCfg := *defaultCfg
	slogCfg.Type = "slog"

	tests := []struct {
		name string
		cfg  *log.Config
	}{
		{
			"Default config",
			defaultCfg,
		},
		{
			"Slog",
			&slogCfg,
		},
	}

	output := func(logger log.Logger, msg string, attrs, args []any) {
		logger.Debug(msg, attrs...)
		logger.Debugf(msg, args...)
		logger.Info(msg, attrs...)
		logger.Infof(msg, args...)
		logger.Warn(msg, attrs...)
		logger.Warnf(msg, args...)
		logger.Error(msg, attrs...)
		logger.Errorf(msg, args...)
	}

	msg := "Test %s"
	attrs := []any{"attr1", "attr1", "attr2", "attr2"}
	args := []any{"arg"}

	for _, tt := range tests {
		logger, cleanup, err := log.ProvideLogger(tt.cfg)
		if err != nil {
			t.Fatal(err)
		}
		if cleanup != nil {
			defer cleanup() // nolint:gocritic
		}

		t.Run(tt.name+" default output", func(_ *testing.T) {
			output(logger, msg, attrs, args)
		})

		t.Run(tt.name+" with attrs", func(_ *testing.T) {
			output(
				logger.WithAttrs("k1", "v1").WithAttrs("k2", "v2"),
				msg, attrs, args)
		})

		t.Run(tt.name+" with context", func(_ *testing.T) {
			wrongCtx := ctxutil.PutFields(context.Background(), map[any]any{"wrong": "wrong"})
			ctx := ctxutil.PutFields(context.Background(), map[any]any{"k": "v"})
			output(
				logger.WithContext(wrongCtx).WithContext(ctx),
				msg, attrs, args)
		})

		t.Run(tt.name+" with attrs and context", func(_ *testing.T) {
			output(
				logger.WithAttrs("k1", "v1").WithAttrs("k2", "v2").WithContext(
					ctxutil.PutFields(context.Background(), map[any]any{"k": "v"})),
				msg, attrs, args)
		})
	}
}
