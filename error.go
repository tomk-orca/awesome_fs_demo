package orcaError

import (
	"errors"
)

type OrcaError struct {
	Err error `json:"-"`
}

func (err OrcaError) Error() string {

			return err.Err.Error() + "\n ths ss111 ewssswdw ssss file 2"
}
func New(e error) OrcaError {
	e.Error = 
	return OrcaError{
		Err: e,
	}
}