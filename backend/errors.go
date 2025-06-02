package main

import "errors"

var (
	ErrorUserNotFound = errors.New(
		"no account with this email address")
	ErrorWrongPassword = errors.New(
		"password does not match")
	ErrorEmailFormatting = errors.New(
		"email address is not correctly formatted")
	ErrorEmailUsed = errors.New(
		"email address is already in use")
)
