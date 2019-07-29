// Global state of the application.

package main

import "errors"

// Path
var (
	dirHome string
)

// Errors
var (
	ErrNoURL = errors.New("url is empty")
)
