package main

import "errors"

var (
	ErrorUserNotFound  = errors.New("no account with this email address")
	ErrorWrongPassword = errors.New("password does not match")
)
