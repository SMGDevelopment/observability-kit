package errs

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cheddartv/observability-kit/logger"
	"github.com/cheddartv/observability-kit/metrics"
)

func WrapError(funcName string, err error, args ...string) error {
	if err == nil {
		return nil
	}
	errWrapNoArgs := "%s: %w"
	errWrapArgs := "%s: %s - %w"
	if len(args) > 0 {
		return fmt.Errorf(errWrapArgs, funcName, strings.Join(args, "; "), err)
	}
	return fmt.Errorf(errWrapNoArgs, funcName, err)
}

func BaseError(err error) error {
	if err == nil {
		return nil
	}
	for {
		newErr := errors.Unwrap(err)
		if newErr == nil {
			return err
		}
		err = newErr
	}
}

func RecordErrorContext(ctx context.Context, err error, metrics metrics.Metrics, log logger.Logger, logAttrs ...logger.LogAttr) {
	if err != nil {
		//getting the base error will ensure uniformity in the error message
		metrics.ErrorInc(BaseError(err).Error())
		//want to log the entire trace for debugging
		log.ErrorContext(ctx, err.Error(), logAttrs...)
	}
}

func RecordError(err error, metrics metrics.Metrics, log logger.Logger, logAttrs ...logger.LogAttr) {
	if err != nil {
		//getting the base error will ensure uniformity in the error message
		metrics.ErrorInc(BaseError(err).Error())

		//want to log the entire trace for debugging
		log.Error(err.Error(), logAttrs...)
	}
}
