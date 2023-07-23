package logger

import "fmt"

// SensibullError collect all logs of sendibull system and log in required system.
// for Simplicity purpose we are printing on terminal
type SensibullError struct {
	SensibullErr error
	Message      string
	ExtraMessage string
	ErrorCode    int
}

// Error method implement error interface.
func (senErr SensibullError) Error() string {
	return fmt.Sprintf("error occurred in system with msg : %s and err code : %d", senErr.Message, senErr.ErrorCode)
}

// Warn print all warn msg into system can be implement letter.
func (senErr SensibullError) Warn() {
	fmt.Println(senErr.Error())
}

// Info print all Info msg into system can be implement letter.
func (senErr SensibullError) Info() {
	fmt.Println(senErr.Error())
}

// Err print all Err msg into system can be implement letter.
func (senErr SensibullError) Err() {
	fmt.Println(senErr.Error())
}

// Debug print all Debug msg into system can be implement letter.
func (senErr SensibullError) Debug() {
	fmt.Println(senErr.Error())
}
