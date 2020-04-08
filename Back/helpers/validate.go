package helpers

import (
	"context"
	"fmt"

	"github.com/go-playground/validator"
)

//Validate struct helper
func Validate(ctx context.Context, m interface{}) error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {

		// this check is nly needed when your code could produce
		// an invald value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.Field())     // by passing alt name to ReportError like below
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return err
	}
	return nil
}
