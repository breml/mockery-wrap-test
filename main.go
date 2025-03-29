//go:generate go tool github.com/vektra/mockery/v3 --config .mockery_fmt_print.yaml

package main

import (
	"context"

	"github.com/breml/mockery-wrap-test/middleware"
	"github.com/breml/mockery-wrap-test/pkg"
)

func main() {
	var e pkg.Exampler = pkg.Example{}

	e = middleware.NewExamplerWithFmtPrint(e)

	e.Example()
	e.ExampleWithArgs(123, "string1")
	e.ExampleWithContext(context.Background())
	e.ExampleWithContextAndArgs(context.Background(), 123, "string1")
	e.ExampleWithVariadicArgs("opt1", "opt2")
	e.ExampleWithArgsAndVariadicArgs(123, "opt1", "opt2")
	e.ExampleWithContextArgsAndVariadicArgs(context.Background(), 123, "opt1", "opt2")
	_ = e.ExampleWithReturn()
	_ = e.ExampleWithErrorReturn()
	_, _ = e.ExampleWithReturnAndError()
	_, _ = e.FullExample(context.Background(), 123, "string1", "opt1", "opt2")
}
