package domain

import "errors"

var (
	// ErrConnectionFailed is returned when websocket connection fails
	ErrConnectionFailed = errors.New("connection failed")

	// ErrInvalidSymbol is returned when a symbol is not supported or malformed
	ErrInvalidSymbol = errors.New("invalid symbol")

	// ErrUpdateFailed is returned when price update fails
	ErrUpdateFailed = errors.New("update failed")

	// ErrConfigNotFound is returned when configuration file is missing
	ErrConfigNotFound = errors.New("configuration not found")
)
