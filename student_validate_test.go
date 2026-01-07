package entity__test

import(
	"testing"
	"github.com/onsi/gomega"
	"TestValidate/entity"
)

func TestValidateStudent(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	t.Run("ValidateStudent", func(t *testing.T) {
		student := &entity.Student{
			StudentID: "B1234567",
			CitizenID: "1234567890123",
			Name:      "John Doe",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       3.5,
			Year:      3,
		}

		err := entity.ValidateStudent(student)
		g.Expect(err).To(gomega.BeNil())

	})

	t.Run("StudentID is invalid", func(t *testing.T) {
		student := &entity.Student{
			StudentID: "A1234567",
			CitizenID: "1234567890123",
			Name:      "John Doe",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       3.5,
			Year:      3,
		}

		err := entity.ValidateStudent(student)
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("StudentID is invalid"))
	})

	t.Run ("CitizenID must be 13 digits", func (t *testing.T) {
		student := &entity.Student{
			StudentID: "B1234567",
			CitizenID: "12345678901",
			Name:      "John Doe",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       3.5,
			Year:      3,
		}

		err := entity.ValidateStudent(student)
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("CitizenID must be 13 digits"))
	})

	t.Run("GPA must be between 0.00 and 4.00", func (t *testing.T) {
		studemt := &entity.Student{
			StudentID: "B1234567",
			CitizenID: "1234567890123",
			Name:      "John Doe",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       4.50,
			Year:      3,
		}

		err := entity.ValidateStudent(studemt)
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("GPA must be between 0.00 and 4.00"))
	})

	t.Run("Year must be between 1 and 8",func (t *testing.T) {
		student := &entity.Student{
			StudentID: "B1234567",
			CitizenID: "1234567890123",
			Name:      "John Doe",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       3.5,
			Year:      10,
		}

		err := entity.ValidateStudent(student)
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Year must be between 1 and 8"))
	})

	t.Run("Name is required", func(t *testing.T) {
		student := &entity.Student{
			StudentID: "B1234567",
			CitizenID: "1234567890123",
			Name:      "",
			Faculty:   "Engineering",
			Major:     "Computer Science",
			GPA:       3.5,
			Year:      3,
		}

		err := entity.ValidateStudent(student)
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name is required"))
	})
}
