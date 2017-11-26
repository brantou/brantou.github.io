package goreturnsdemo


import "errors"

func F() (int, error) { return errors.New("foo") }
