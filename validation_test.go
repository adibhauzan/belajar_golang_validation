package belajar_golang_validation

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
	"testing"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("")
	}

}

func TestValidationField(t *testing.T) {
	validate := validator.New()

	var user string = ""

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "adib"
	confirmPassword := "adib"

	err := validate.VarWithValueCtx(context.Background(), password, confirmPassword, "eqfield")
	if err != nil {
		log.Panic(err)
	}
}

func TestValidationMultipleField(t *testing.T) {
	validate := validator.New()

	var user int = 9132409132

	err := validate.Var(user, "required,number")
	if err != nil {
		log.Panic(err)
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	angka := "11111"

	err := validate.Var(angka, "required,numeric,min=5,max=10")
	if err != nil {
		log.Panic(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Name     string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := &LoginRequest{
		Name:     "Adib@gmail.com",
		Password: "Adibas",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		log.Panic()
	}

	fmt.Printf("Name = %s \nPassword = %s \n", loginRequest.Name, loginRequest.Password)
}

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Name     string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}
	validate := validator.New()

	loginRequest := &LoginRequest{
		Name:     "adib@gmail.com",
		Password: "Adibadib",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationError := err.(validator.ValidationErrors)

		for _, fieldError := range validationError {
			t.Errorf("error %s on tag %s with error %s \n", fieldError.Field(), fieldError.Tag(), fieldError.Error())
		}
	} else {
		fmt.Printf("Name = %s \nPassword = %s \n", loginRequest.Name, loginRequest.Password)
	}

}

func TestCrossField(t *testing.T) {
	type RegisterUser struct {
		UserName        string `validate:"required,email"`
		Password        string `validate:"required,min=6"`
		ConfirmPassword string `validate:"required,min=6,eqfield=Password"`
	}
	validate := validator.New()
	registerUser := &RegisterUser{
		UserName:        "adibhauzan48@gmail.com",
		Password:        "adib123123",
		ConfirmPassword: "adib123123",
	}

	err := validate.Struct(registerUser)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		for _, fieldError := range validationError {
			t.Errorf("error %s on tag %s with error %s \n", fieldError.Field(), fieldError.Tag(), fieldError.Error())
		}
	} else {
		fmt.Println("User Registered")
	}
}

func TestValidationNestedStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string   `validate:"required"`
		Name    string   `validate:"required"`
		Address *Address `validate:"required"`
	}
	validate := validator.New()

	user := &User{
		Id:   "1",
		Name: "Adib Hauzan",
		Address: &Address{
			City:    "Bandung",
			Country: "Indonesia",
		},
	}
	err := validate.Struct(user)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("Success")
	}
}

func TestValidationCollectionStruct(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string     `validate:"required"`
		Name      string     `validate:"required"`
		Addresses []*Address `validate:"required,dive"`
	}
	validate := validator.New()

	user := &User{
		Id:   "1",
		Name: "adib hauzan",
		Addresses: []*Address{
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
			{
				City:    "Padang",
				Country: "Indonesia",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("Success")
	}
}

func TestValidationBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string     `validate:"required"`
		Name      string     `validate:"required"`
		Addresses []*Address `validate:"required,dive"`
		Hobbies   []string   `validate:"dive,required,min=3"`
	}

	validate := validator.New()

	user := &User{
		Id:   "1",
		Name: "Adib Hauzan",
		Addresses: []*Address{
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
			{
				City:    "Padang",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Main Game",
			"Belajar",
			"Berenang",
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("Success")
	}
}

func TestMapValidation(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}
	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []*Address        `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=3"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	}

	validate := validator.New()

	user := &User{
		Id:   "1",
		Name: "Adib Hauzan Sofyan",
		Addresses: []*Address{
			{
				City:    "Bandung",
				Country: "Indonesia",
			},
			{
				City:    "Padang",
				Country: "Indonesia",
			},
		},
		Hobbies: []string{
			"Main Game",
			"Belajar",
			"Berenang",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDN 14 Laing",
			},
			"SMP": {
				Name: "SMPN 30 Padang",
			},
			"SMA": {
				Name: "SMAN 9 Padang",
			},
			"Kuliah": {
				Name: "Universitas Logistik Dan Bisnis Internasional",
			},
		},
	}

	err := validate.Struct(user)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("Success")
	}
}

func TestValidationTagAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := &Seller{
		Id:     "",
		Name:   "Adib Hauzan",
		Owner:  "Gojek",
		Slogan: "Hidup Bak Diperkosa, mau gamau siap ga siap, jalani aja",
	}
	err := validate.Struct(seller)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println("success")
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}

		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidationFunction(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	user := &LoginRequest{
		Username: "ADIBHHH",
		Password: "adibadibadib",
	}

	err := validate.Struct(user)
	if err != nil {
		t.Log(err.Error())
	} else {
		fmt.Println("Success")
	}

}
