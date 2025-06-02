package main

type SignupInput struct {
	Email          string
	Password       string
	TargetLanguage string
}

type LoginInput struct {
	Email    string
	Password string
}
