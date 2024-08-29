package domain

import "errors"

var (
	ErrUnauthorized        = errors.New("Authentication failed")       //401
	ErrInternalServerError = errors.New("Internal server error")       //500
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")        //422
	ErrNotFound            = errors.New("Not found")                   //404
	ErrAlredyExist         = errors.New("Follow request alredy exist") //400
)
