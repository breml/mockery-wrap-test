package pkg

import (
	"context"
	"errors"
	"fmt"
)

type Exampler interface {
	Example()
	ExampleWithArgs(arg0 int, arg1 string)
	ExampleWithContext(ctx context.Context)
	ExampleWithContextAndArgs(ctx context.Context, arg0 int, arg1 string)
	ExampleWithVariadicArgs(args ...string)
	ExampleWithArgsAndVariadicArgs(arg0 int, args ...string)
	ExampleWithContextArgsAndVariadicArgs(ctx context.Context, arg0 int, args ...string)
	ExampleWithReturn() string
	ExampleWithErrorReturn() error
	ExampleWithReturnAndError() (string, error)
	FullExample(ctx context.Context, arg0 int, arg1 string, args ...string) (string, error)
}

type Example struct{}

var _ Exampler = Example{}

func (Example) Example() {
	fmt.Println("== in Example ")
}

func (Example) ExampleWithArgs(arg0 int, arg1 string) {
	fmt.Println("== in ExampleWithArgs")
}

func (Example) ExampleWithContext(ctx context.Context) {
	fmt.Println("== in ExampleWithContext")
}

func (Example) ExampleWithContextAndArgs(ctx context.Context, arg0 int, arg1 string) {
	fmt.Println("== in ExampleWithContextAndArgs")
}

func (Example) ExampleWithVariadicArgs(arg0 ...string) {
	fmt.Println("== in ExampleWithVariadicArgs")
}

func (Example) ExampleWithArgsAndVariadicArgs(arg0 int, args ...string) {
	fmt.Println("== in ExampleWithArgsAndVariadicArgs")
}

func (Example) ExampleWithContextArgsAndVariadicArgs(ctx context.Context, arg0 int, args ...string) {
	fmt.Println("== in ExampleWithContextArgsAndVariadicArgs")
}

func (Example) ExampleWithReturn() string {
	fmt.Println("== in ExampleWithReturn")

	return "ret"
}

func (Example) ExampleWithErrorReturn() error {
	fmt.Println("== in ExampleWithErrorReturn")

	return errors.New("boom")
}

func (Example) ExampleWithReturnAndError() (string, error) {
	fmt.Println("== in ExampleWithReturnAndError")

	return "ret", errors.New("boom")
}

func (Example) FullExample(ctx context.Context, arg0 int, arg1 string, args ...string) (string, error) {
	fmt.Println("== in FullExample")

	return "ret", errors.New("boom")
}
