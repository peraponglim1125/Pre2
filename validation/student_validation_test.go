package validation_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/Pre2/entity"
)

func TestStudent_Valid(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		Fullname: "Peerapong Limpasritragul",
		Age:      21,
		Email:    "test@example.com",
		GPA:      3.50,
	}

	ok, err := entity.ValidateStudent(&student)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

func TestStudent_Fullname_Required(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		Fullname: "",
		Age:      20,
		Email:    "test@example.com",
		GPA:      3.00,
	}
//วัดี
//0hk5555555555
	ok, err := entity.ValidateStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Fullname is required"))
}

func TestStudent_Age_LessThan18(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		Fullname: "Test User",
		Age:      17,
		Email:    "test@example.com",
		GPA:      3.00,
	}

	ok, err := entity.ValidateStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Age must be at least 18"))
}

func TestStudent_Email_Invalid(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		Fullname: "Test User",
		Age:      20,
		Email:    "invalid-email",
		GPA:      3.00,
	}

	ok, err := entity.ValidateStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Email is invalid"))
}

func TestStudent_GPA_OutOfRange(t *testing.T) {
	g := NewGomegaWithT(t)

	student := entity.Student{
		Fullname: "Test User",
		Age:      20,
		Email:    "test@example.com",
		GPA:      4.50,
	}

	ok, err := entity.ValidateStudent(&student)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
}
