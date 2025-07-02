package schemas

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type NoteCreate struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Url         *string    `json:"url"`
	Date        CustomDate `json:"date" validate:"required,dateFormat"`
	Time        CustomTime `json:"time" validate:"required,timeFormat"`
}

type NoteUpdate struct {
	ID          string     `json:"id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Url         *string    `json:"url"`
	Date        CustomDate `json:"date" validate:"required,dateFormat"`
	Time        CustomTime `json:"time" validate:"required,timeFormat"`
}

func validateDateFormat(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	match, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`, date)
	return match
}

func validateTimeFormat(fl validator.FieldLevel) bool {
	time := fl.Field().String()
	match, _ := regexp.MatchString(`^([01][0-9]|2[0-3]):[0-5][0-9]$`, time)
	return match
}

func (n *NoteCreate) Validate() error {
	validate := validator.New()

	_ = validate.RegisterValidation("dateFormat", validateDateFormat)
	_ = validate.RegisterValidation("timeFormat", validateTimeFormat)

	err := validate.Struct(n)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es inválido, revisar: (%s)", field, tag)
}

func (n *NoteUpdate) Validate() error {
	validate := validator.New()

	_ = validate.RegisterValidation("dateFormat", validateDateFormat)
	_ = validate.RegisterValidation("timeFormat", validateTimeFormat)

	err := validate.Struct(n)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	field := validationErr.Field()
	tag := validationErr.Tag()

	return fmt.Errorf("campo %s es inválido, revisar: (%s)", field, tag)
}

type NoteDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Date        time.Time `json:"date"`
	Time        time.Time `json:"time"`
}

type NoteResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         *string `json:"url"`
	Date        time.Time `json:"date"`
	Time        time.Time `json:"time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// type NoteCreate struct {
// 	Title       string `json:"title" validate:"required"`
// 	Description string `json:"description" validate:"required"`
// 	Url         *string `json:"url"`
// 	Date        CustomDate `json:"date" validate:"required,dateFormat"`
// 	Time        CustomTime `json:"time" validate:"required,timeFormat"`
// }

// func validateDateFormat(fl validator.FieldLevel) bool {
// 	date := fl.Field().String()
// 	// dd-mm-yyyy (01-12-2025)
// 	match, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-\d{4}$`, date)
// 	return match
// }

// func validateTimeFormat(fl validator.FieldLevel) bool {
// 	time := fl.Field().String()
// 	// hh:mm (00:00 to 23:59)
// 	match, _ := regexp.MatchString(`^([01][0-9]|2[0-3]):[0-5][0-9]$`, time)
// 	return match
// }

// func (n *NoteCreate) Validate() error {
// 	validate := validator.New()

// 	// Registrar validaciones personalizadas
// 	_ = validate.RegisterValidation("dateformat", validateDateFormat)
// 	_ = validate.RegisterValidation("timeformat", validateTimeFormat)

// 	err := validate.Struct(n)
// 	if err == nil {
// 		return nil
// 	}

// 	validationErr := err.(validator.ValidationErrors)[0]
// 	field := validationErr.Field()
// 	tag := validationErr.Tag()
// 	param := validationErr.Param()

// 	return fmt.Errorf("campo %s es inválido, revisar: (%s) (%s)", field, tag, param)
// }