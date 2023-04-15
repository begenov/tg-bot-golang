package e

import "fmt"

func Wrap(msg string, e error) error {
	return fmt.Errorf("%s: %w", msg, e)
}

func WrapIfErr(msg string, e error) error {
	if e == nil {
		return nil
	}
	return Wrap(msg, e)
}
