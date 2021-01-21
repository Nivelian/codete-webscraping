package helpers

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func LogErr(err error, tpl string, args ...interface{}) error {
	msg := fmt.Sprintf(tpl, args...)
	err = errors.Wrap(err, msg)
	logrus.Error(err)
	return err
}
