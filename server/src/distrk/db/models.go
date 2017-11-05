package db

import (
	"github.com/jinzhu/gorm"
)

type AdmissionType struct {
	gorm.Model
	Name string
}

type Patient struct {
	gorm.Model
	AdmissionType   AdmissionType
	AdmissionTypeID int
	Name            string
	Consults        []Consult `gorm:"many2many:patients_consults;"`
}

type Department struct {
	gorm.Model
	Name     string
	Location string
	Priority int
}

type Consult struct {
	gorm.Model
	Department   Department
	DepartmentID int
	Completed    bool
}
