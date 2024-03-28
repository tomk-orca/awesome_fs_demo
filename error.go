package orcaError

import (
	"errors"
)

type OrcaError struct {
	Err error `json:"-"`
}

func (err OrcaError) Error() string {

	return err.Err.Error() + "\n this is orca error file 7"
}

func New(e error) OrcaError {
	e.Error = 
	return OrcaError{
		Err: e,
	}
}