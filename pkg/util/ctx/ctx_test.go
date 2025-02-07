package ctxutil_test

import (
	"context"
	"reflect"
	"testing"

	ctxutil "github.com/teamsorghum/go-common/pkg/util/ctx"
)

func TestContext_GetContextFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		ctx            context.Context
		expectedFields map[any]any
	}{
		{
			"empty fields",
			context.Background(),
			map[any]any{},
		},
		{
			"non empty fields",
			context.WithValue(context.Background(), ctxutil.Key, ctxutil.ValueType{
				"key": "value",
			}),
			map[any]any{
				"key": "value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fields := ctxutil.GetFields(tt.ctx)
			if !reflect.DeepEqual(fields, tt.expectedFields) {
				t.Fatalf("expectedFields = %+v, get %+v", tt.expectedFields, fields)
			}
		})
	}
}

func TestContext_PutContextFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		ctx         context.Context
		fields      map[any]any
		expectedCtx context.Context
	}{
		{
			"nil context",
			nil,
			map[any]any{},
			nil,
		},
		{
			"nil fields",
			context.Background(),
			nil,
			context.Background(),
		},
		{
			"add kv",
			context.Background(),
			map[any]any{
				"key": "value",
			},
			context.WithValue(context.Background(), ctxutil.Key, ctxutil.ValueType{
				"key": "value",
			}),
		},
		{
			"overwrite kv",
			context.WithValue(context.Background(), ctxutil.Key, ctxutil.ValueType{
				"key1": "value1",
				"key2": "value2",
			}),
			map[any]any{
				"key1": "aloha",
			},
			context.WithValue(context.Background(), ctxutil.Key, ctxutil.ValueType{
				"key1": "aloha",
				"key2": "value2",
			}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := ctxutil.PutFields(tt.ctx, tt.fields)
			if tt.expectedCtx == nil && ctx != nil {
				t.Fatalf("Expect ctx to be nil, get %+v", ctx)
			}
			expectedFields := ctxutil.GetFields(tt.expectedCtx)
			actualFields := ctxutil.GetFields(ctx)
			if len(expectedFields) == 0 && len(actualFields) == 0 {
				return
			}
			if !reflect.DeepEqual(expectedFields, actualFields) {
				t.Fatalf("expectedCtx == %+v, actualFields == %+v", expectedFields, actualFields)
			}
		})
	}
}
