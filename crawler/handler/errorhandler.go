package handler

import "fmt"

type ErrorHandler struct {
	errorChan chan error
}

func (this *ErrorHandler) Build() {
	this.errorChan = make(chan error)
	go func() {
		for {
			err := <-this.errorChan
			this.handleError(err)
		}
	}()
}

func (this *ErrorHandler) GetErrorChan() chan error {
	return this.errorChan
}

func (this *ErrorHandler) handleError(err error) {
	fmt.Println(err.Error())
}
