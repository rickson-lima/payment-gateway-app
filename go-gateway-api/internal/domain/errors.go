package domain

import "errors"

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrDuplicatedAPIKey = errors.New("API key already exists")
	ErrInvoiceNotFound  = errors.New("invoice not found")
	ErrUnauthorized     = errors.New("unauthorized access")
)
