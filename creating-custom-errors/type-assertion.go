/* Type Assertion and Custom Errors */
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type RequesError struct {
	StatusCode int
	Err        error
}

func (r *RequesError) Error() string {
	return r.Err.Error()
}

func (r *RequesError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func doReques() error {
	return &RequesError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func main() {
	err := doReques()
	if err != nil {
		fmt.Println(err)
		re, ok := err.(*RequesError)
		if ok {
			if re.Temporary() {
				fmt.Println("This request can be tried again")
			} else {
				fmt.Println("This request cannot be tried again")
			}
		}
		os.Exit(1)
	}

	fmt.Println("Success!")
}
