package test

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Age      int    `json:"age" validate:"required,min=18"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"-"`
}

// func validateField(value interface{}, tag string) error {
// 	validate := validator.New()
// 	field := struct {
// 		Value interface{} `validate:"` + tag + `"`
// 	}{
// 		Value: value,
// 	}
//
// 	return validate.Struct(field)
// }

func TestStructValidate(t *testing.T) {
	req := User{
		Name: "123",
		Age:  18,
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		fmt.Println(err)
	}
}
