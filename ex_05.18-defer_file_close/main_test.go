package main

import (
	"errors"
	"fmt"
	"testing"
)

type closeable struct {
	name string
	err  error
}

type thing struct {
	name string
	err  error
}

func (t closeable) Close() error {
	fmt.Println("closing: ", t.name)
	return t.err
}

func genCloseableNil(name string) (closeable, error) {
	return closeable{name: name}, nil
}

func genCloseableNilErr(name string) (closeable, error) {
	return closeable{name: name, err: errors.New("closeable with error")}, nil
}

func importantError(name string) (int, error) {
	return 1, errors.New(fmt.Sprintf("important: %s\n", name))
}

func importantNoError(name string) (int, error) {
	return 1, nil
}

func test() (err error) {
	t1, err := genCloseableNil("one")
	if err != nil {
		fmt.Println("one, error returning")
		return
	}
	defer t1.Close()

	t2, err := genCloseableNilErr("two")
	if err != nil {
		fmt.Println("two, error returning")
		return
	}

	_, err = importantError("test1 three")
	if closeErr := t2.Close(); err == nil {
		err = closeErr
	}
	return err
}

func test2() (err error) {
	t1, err := genCloseableNil("one")
	if err != nil {
		fmt.Println("one, error returning")
		return
	}
	defer t1.Close()

	t2, err := genCloseableNilErr("two")
	if err != nil {
		fmt.Println("two, error returning")
		return
	}
	defer func(error) error {
		if closeErr := t2.Close(); err == nil {
			err = closeErr
		}
		return err
	}(err)

	_, err = importantError("test2 three")

	return err
}

func TestFetch(t *testing.T) {
	const fname = "main"
	err1 := test()
	err2 := test2()
	if err1 != err2 {
		fmt.Printf("%s: want %q got %q", fname, err1, err2)
	}
}
