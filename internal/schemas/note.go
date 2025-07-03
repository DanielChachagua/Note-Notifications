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
	Warn        bool       `json:"warn"`
}

type NoteUpdate struct {
	ID          string     `json:"id" validate:"required"`
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description" validate:"required"`
	Url         *string    `json:"url"`
	Date        CustomDate `json:"date" validate:"required,dateFormat"`
	Time        CustomTime `json:"time" validate:"required,timeFormat"`
	Warn        *bool      `json:"warn"`
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
	ID    string     `json:"id"`
	Title string     `json:"title"`
	Date  CustomDate `json:"date"`
	Time  CustomTime `json:"time"`
	Warn  bool       `json:"warn"`
}

type NoteResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Url         *string    `json:"url"`
	Date        CustomDate `json:"date"`
	Time        CustomTime `json:"time"`
	Warn        bool       `json:"warn"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}


