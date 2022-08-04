package model

import (
    "gorm.io/gorm"
)

type Employee struct {
    gorm.Model           // Adds some metadata fields to the table
    ID          int `gorm:"not null"` // Explicitly specify the type to be uuid
    Name	    string
    Gender   	string
    City        string
}
