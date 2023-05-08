package internal

import "github.com/pkg/errors"

func handleErr(err error) {
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, ""))
	}
}

func handleErrDebug(err error) {
	if err != nil {
		log.Debug(err)
	}
}
