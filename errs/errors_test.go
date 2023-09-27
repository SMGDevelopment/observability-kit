package errs

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cheddartv/rmp-observability-kit/logger"
)

var myTestError = errors.New("my test error")

func TestWrapError(t *testing.T) {
	funcName := "TestWrapError"

	tests := []struct {
		name            string
		err             error
		args            []string
		expectErrString string
	}{
		{
			name:            "No Args",
			err:             myTestError,
			args:            nil,
			expectErrString: "TestWrapError: my test error",
		},
		{
			name:            "One Args",
			err:             myTestError,
			args:            []string{"one"},
			expectErrString: "TestWrapError: one - my test error",
		},
		{
			name:            "Multiple Args",
			err:             myTestError,
			args:            []string{"one", "two", "three"},
			expectErrString: "TestWrapError: one; two; three - my test error",
		},
		{
			name: "Nil Error",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WrapError(funcName, tt.err, tt.args...)
			if tt.err == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tt.expectErrString, err.Error())
			}
		})
	}
}

func TestBaseError(t *testing.T) {
	funcName := "TestBaseError"

	tests := []struct {
		name string
		err  error
		wrap func(error) error
	}{
		{
			name: "Original",
			err:  myTestError,
			wrap: func(err error) error {
				return err
			},
		},
		{
			name: "Single Wrap",
			err:  myTestError,
			wrap: func(err error) error {
				return WrapError(funcName, err)
			},
		},
		{
			name: "Single Wrap With Args",
			err:  myTestError,
			wrap: func(err error) error {
				return WrapError(funcName, err, "help")
			},
		},
		{
			name: "Deep Wrap",
			err:  myTestError,
			wrap: func(err error) error {
				return WrapError(funcName,
					WrapError(funcName,
						WrapError(funcName, err, "help")),
					"help")
			},
		},
		{
			name: "Nil",
			err:  nil,
			wrap: func(err error) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wrappedError := tt.wrap(tt.err)
			err := BaseError(wrappedError)
			if tt.err == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, tt.err.Error(), err.Error())
			}
		})
	}
}

func TestRecordError(t *testing.T) {
	l := logger.InitLogger("DEV", os.Stdout)
	RecordErrorContext(context.TODO(), myTestError, l)
	RecordError(myTestError, l)
}
