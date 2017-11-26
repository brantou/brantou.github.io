package goreturnsdemo

import "errors"

func F() (*MyType, int, error) { return errors.New("foo") }
