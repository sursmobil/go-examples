package multithreading

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestWaiterSuccess(t *testing.T) {
	sut := NewWaiter(successJob(1, time.Millisecond))
	sut.Start()
	result, err := sut.Wait(time.Second)
	require.Nil(t, err)
	require.Equal(t, 1, result, "Waiter hasn't returned job result")
}

func TestWaiterFail(t *testing.T) {
	sut := NewWaiter(failJob())
	sut.Start()
	result, err := sut.Wait(time.Second)
	require.NotNil(t, err, "Job error should be passed as error")
	_, ok := err.(WaiterError)
	require.False(t, ok, "Job error should not be of type WaiterError")
	require.Equal(t, nil, result, "Waiter result should be nil after job error")
}

func TestWaiterTimeout(t *testing.T) {
	sut := NewWaiter(successJob(1, time.Second))
	sut.Start()
	result, err := sut.Wait(time.Millisecond)
	require.Equal(t, err, ErrWaiterTimeout, "Waiter did not timed out")
	_, ok := err.(WaiterError)
	require.True(t, ok, "Timeout error should be of type WaiterError")
	require.Nil(t, result, "Result should be nil after job error")
}

type waitResult struct {
	result interface{}
	err    error
}

func TestWaiterCancel(t *testing.T) {
	sut := NewWaiter(successJob(1, time.Second))
	sut.Start()
	resultCh := make(chan waitResult)
	go func() {
		result, err := sut.Wait(time.Second)
		resultCh <- waitResult{result, err}
	}()
	sut.Cancel()
	r := <-resultCh
	require.Equal(t, r.err, ErrWaiterCancel, "Waiter should be canceled")
	_, ok := r.err.(WaiterError)
	require.True(t, ok, "Cancel error should be of type WaiterError")
	require.Nil(t, r.result, "Result should be nil after timeout")
}

func successJob(expectedResult interface{}, after time.Duration) WaiterWork {
	return func(result WaiterWorkResult) {
		time.Sleep(after)
		result <- expectedResult
	}
}

func failJob() WaiterWork {
	return func(result WaiterWorkResult) {
		result <- errors.New("Job error")
	}
}
