package errs

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cheddartv/rmp-observability-kit/logger"
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

func RecordErrorContext(context context.Context, err error, logger logger.Logger) {
	if err != nil {
		//getting the base error will ensure uniformity in the error message
		//TODO: metrics
		//Metrics().RED.Errors.WithLabelValues(BaseError(err).Error()).Inc()
		//want to log the entire trace for debugging
		logger.ErrorContext(context, err.Error())
	}
}

func RecordError(err error, logger logger.Logger) {
	if err != nil {
		//getting the base error will ensure uniformity in the error message
		//TODO: metrics
		//Metrics().RED.Errors.WithLabelValues(BaseError(err).Error()).Inc()
		//want to log the entire trace for debugging
		logger.Error(err.Error())
	}
}
