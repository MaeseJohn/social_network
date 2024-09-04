package domain

import "errors"

var (
	ErrFollowExist         = errors.New("Follow exist")                //400
	ErrRequestExist        = errors.New("Follow request alredy exist") //400
	ErrUnauthorized        = errors.New("Authentication failed")       //401
	ErrNotFound            = errors.New("Not found")                   //404
	ErrUnprocessableEntity = errors.New("Unprocessable Entity")        //422
	ErrInternalServerError = errors.New("Internal server error")       //500
)
