package multithreading

import "fmt"

// WaiterError is generic type for errors thrown by waiter so they can be
// differentiated from other errors
type WaiterError string

const (
	// ErrWaiterTimeout is thrown when Waiter has timed out
	ErrWaiterTimeout WaiterError = "Waiter timed out"

	// ErrWaiterCancel is thrown when Waiter has been canceled
	ErrWaiterCancel WaiterError = "Waiter has been canceled"
)

func (err WaiterError) Error() string {
	return fmt.Sprintf("%s", err)
}
