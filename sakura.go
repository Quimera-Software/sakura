// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import "github.com/pkg/errors"

var cfg Config

func Setup(config Config) error {
	cfg = config

	err := newLogFile()
	if err != nil {
		return errors.Wrap(err, "unable to create log file")
	}

	return nil
}
