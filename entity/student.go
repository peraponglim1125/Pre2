package entity

import "github.com/asaskevich/govalidator"

type Student struct {
	Fullname string  `valid:"required~Fullname is required"`
	Age      uint    `valid:"range(18|150)~Age must be at least 18"`
	Email    string  `valid:"required~Email is required,email~Email is invalid"`
	GPA      float32 `valid:"float,range(0|4)~GPA must be between 0.00 and 4.00"`
}

func ValidateStudent(student *Student) (bool, error) {
	return govalidator.ValidateStruct(student)
}
