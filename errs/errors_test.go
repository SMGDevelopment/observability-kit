package errs

import (
	"errors"
	"testing"
)

func TestRecordError(t *testing.T) {
	RecordError(errors.New("my test err"))
}
