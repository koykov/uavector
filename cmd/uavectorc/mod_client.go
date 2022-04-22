package main

import "fmt"

type clientModule struct{}

func (m clientModule) Validate(input, _ string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	return nil
}

func (m clientModule) Compile(w moduleWriter, input, target string) (err error) {
	return
}
