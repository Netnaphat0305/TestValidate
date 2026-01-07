package entity

import "github.com/asaskevich/govalidator"

type Student struct {
	StudentID string  `valid:"required,student_id~StudentID is invalid"`
	CitizenID string  `valid:"required,citizen_id~CitizenID must be 13 digits"`
	Name      string  `valid:"required~Name is required"`
	Faculty   string  `valid:"required~Faculty is required"`
	Major     string  `valid:"required~Major is required"`
	GPA       float64 `valid:"gpa_range~GPA must be between 0.00 and 4.00"`
	Year      int     `valid:"year_range~Year must be between 1 and 8"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("student_id", func(i interface{}, _ interface{}) bool {
		studentid, ok := i.(string)
		if !ok {
			return false
		}
		return govalidator.Matches(studentid, `^B\d{7}$`)
	})

	govalidator.CustomTypeTagMap.Set("citizen_id", func(i interface{}, _ interface{}) bool {
		citizenid, ok := i.(string)
		if !ok {
			return false
		}
		return govalidator.Matches(citizenid, `^\d{13}$`)
	})

	govalidator.CustomTypeTagMap.Set("gpa_range", func(i interface{}, _ interface{}) bool {
		gpa, ok := i.(float64)
		return ok && gpa >= 0.00 && gpa <= 4.00
	})

	govalidator.CustomTypeTagMap.Set("year_range", func(i interface{}, _ interface{}) bool {
		year, ok := i.(int)
		return ok && year >= 1 && year <= 8

	})
}

func ValidateStudent(s *Student) error {
	_, err := govalidator.ValidateStruct(s)
	return err
}
