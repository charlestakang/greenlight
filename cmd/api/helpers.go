package main

import (
	"errors"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app application) readIDParam(ps httprouter.Params) (int64, error) {
	params := ps.ByName("id")
	id, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}