package models

type ErrorModel struct {
	Service string
	Err error
	Message string
	Sev int
}
