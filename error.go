package orcaError

import (
	"errors"
)

type OrcaError struct {
	Err error `json:"-"`
}

func (err OrcaError) Error() string {

	return err.Err.Error() + "\n this ihhhs 1esse11orca error file 10"
}

func New(e error) OrcaError {
	e.Error = 
	return OrcaError{
		Err: e,
	}
}