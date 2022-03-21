package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" form:"name" validate:"nonzero"`
	CPF  string `json:"cpf" form:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" form:"rg" validate:"len=9"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}
