package orcaError

import (
	"errors"
)

type OrcaError struct {
	Err error `json:"-"`
}

func (err OrcaError) Error() string {

	return err.Err.Error() + "\n this is wdwdq error111 file 2"
}

func New(e error) OrcaError {
	e.Error = 
	return OrcaError{
		Err: e,
	}
}