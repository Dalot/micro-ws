// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package observer

import (
	"sync"
)

// Ensure, that databaseMock does implement Database.
// If this is not the case, regenerate this file with moq.
var _ Database = &databaseMock{}

// databaseMock is a mock implementation of Database.
//
// 	func TestSomethingThatUsesDatabase(t *testing.T) {
//
// 		// make and configure a mocked Database
// 		mockedDatabase := &databaseMock{
// 			GetByObservableFunc: func(name string) ([]IObserver, error) {
// 				panic("mock out the GetByObservable method")
// 			},
// 			SetFunc: func(name string, obs IObserver) error {
// 				panic("mock out the Set method")
// 			},
// 			SetObservableFunc: func(name string, observers []IObserver) error {
// 				panic("mock out the SetObservable method")
// 			},
// 		}
//
// 		// use mockedDatabase in code that requires Database
// 		// and then make assertions.
//
// 	}
type databaseMock struct {
	// GetByObservableFunc mocks the GetByObservable method.
	GetByObservableFunc func(name string) ([]IObserver, error)

	// SetFunc mocks the Set method.
	SetFunc func(name string, obs IObserver) error

	// SetObservableFunc mocks the SetObservable method.
	SetObservableFunc func(name string, observers []IObserver) error

	// calls tracks calls to the methods.
	calls struct {
		// GetByObservable holds details about calls to the GetByObservable method.
		GetByObservable []struct {
			// Name is the name argument value.
			Name string
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Name is the name argument value.
			Name string
			// Obs is the obs argument value.
			Obs IObserver
		}
		// SetObservable holds details about calls to the SetObservable method.
		SetObservable []struct {
			// Name is the name argument value.
			Name string
			// Observers is the observers argument value.
			Observers []IObserver
		}
	}
	lockGetByObservable sync.RWMutex
	lockSet             sync.RWMutex
	lockSetObservable   sync.RWMutex
}

// GetByObservable calls GetByObservableFunc.
func (mock *databaseMock) GetByObservable(name string) ([]IObserver, error) {
	if mock.GetByObservableFunc == nil {
		panic("databaseMock.GetByObservableFunc: method is nil but Database.GetByObservable was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockGetByObservable.Lock()
	mock.calls.GetByObservable = append(mock.calls.GetByObservable, callInfo)
	mock.lockGetByObservable.Unlock()
	return mock.GetByObservableFunc(name)
}

// GetByObservableCalls gets all the calls that were made to GetByObservable.
// Check the length with:
//     len(mockedDatabase.GetByObservableCalls())
func (mock *databaseMock) GetByObservableCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockGetByObservable.RLock()
	calls = mock.calls.GetByObservable
	mock.lockGetByObservable.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *databaseMock) Set(name string, obs IObserver) error {
	if mock.SetFunc == nil {
		panic("databaseMock.SetFunc: method is nil but Database.Set was just called")
	}
	callInfo := struct {
		Name string
		Obs  IObserver
	}{
		Name: name,
		Obs:  obs,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	return mock.SetFunc(name, obs)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedDatabase.SetCalls())
func (mock *databaseMock) SetCalls() []struct {
	Name string
	Obs  IObserver
} {
	var calls []struct {
		Name string
		Obs  IObserver
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}

// SetObservable calls SetObservableFunc.
func (mock *databaseMock) SetObservable(name string, observers []IObserver) error {
	if mock.SetObservableFunc == nil {
		panic("databaseMock.SetObservableFunc: method is nil but Database.SetObservable was just called")
	}
	callInfo := struct {
		Name      string
		Observers []IObserver
	}{
		Name:      name,
		Observers: observers,
	}
	mock.lockSetObservable.Lock()
	mock.calls.SetObservable = append(mock.calls.SetObservable, callInfo)
	mock.lockSetObservable.Unlock()
	return mock.SetObservableFunc(name, observers)
}

// SetObservableCalls gets all the calls that were made to SetObservable.
// Check the length with:
//     len(mockedDatabase.SetObservableCalls())
func (mock *databaseMock) SetObservableCalls() []struct {
	Name      string
	Observers []IObserver
} {
	var calls []struct {
		Name      string
		Observers []IObserver
	}
	mock.lockSetObservable.RLock()
	calls = mock.calls.SetObservable
	mock.lockSetObservable.RUnlock()
	return calls
}
