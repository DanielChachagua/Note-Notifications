package schemas

import (
	"fmt"

	"github.com/go-playground/validator"
	"golang.org/x/oauth2"
)

type CalendarCreate struct {
	Summary     string      `json:"title"`
	Location    *string     `json:"description"`
	Description *string     `json:"url"`
	Date        CustomDate  `json:"date" validate:"required,dateFormat"`
	Time        *CustomTime `json:"time" validate:"omitempty,timeFormat"`
}

type CalendarUpdate struct {
	ID          string      `json:"id"`
	Summary     string      `json:"summary"`
	Location    *string     `json:"location"`
	Description *string     `json:"description"`
	Date        CustomDate  `json:"date" validate:"required,dateFormat"`
	Time        *CustomTime `json:"time" validate:"omitempty,timeFormat"`
}

type CreateEvent struct {
	Token *oauth2.Token `json:"token" validate:"required"`
	Event CalendarCreate `json:"event" validate:"required"`
}

type DeleteEvent struct {
	Token *oauth2.Token `json:"token" validate:"required"`
	EventIds []string `json:"event_ids" validate:"required"`
}

func (n *CalendarCreate) Validate() error {
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

func (n *CalendarUpdate) Validate() error {
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


type UpdateEvent struct {
	Token *oauth2.Token `json:"token" validate:"required"`
	Event CalendarUpdate `json:"event" validate:"required"`
}

func (n *UpdateEvent) Validate() error {
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