package errs

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/cheddartv/rmp-observability-kit/metrics"
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

func RecordError(err error) {
	if err != nil {
		//getting the base error will ensure uniformity in the error message
		metrics.MetricError(BaseError(err).Error())
		//TODO: properly log error
		//want to log the entire trace for debugging
		log.Println(err.Error())
	}
}
