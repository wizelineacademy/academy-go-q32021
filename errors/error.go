package errors

import "time"

type ResponseError struct {
	Code int			`json:"code"`
	Date time.Time		`json:"date"`
	Message  string		`json:"message"`
}
