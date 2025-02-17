package datatype

import (
	"github.com/google/uuid"
)

// FIXME: im not sure why the DTO (data access object) goes here. This maps DataTypes from the DB to the program
type DTO struct {
    Id uuid.UUID `json:"id"`
    FieldName string `json:"field_name"`
}

type DataType struct {
    FieldName string `json:"field_name"`
}



