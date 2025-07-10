package dao

import (
	"StudentService/database"
	"StudentService/model"
)

func GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	result := database.DB.Find(&students)
	return students, result.Error
}

func GetStudentByID(id string) (*model.Student, error) {
	var student model.Student
	result := database.DB.First(&student, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func CreateStudent(student *model.Student) error {
	return database.DB.Create(student).Error
}

func UpdateStudent(student *model.Student) error {
	return database.DB.Save(student).Error
}

func DeleteStudent(id string) error {
	return database.DB.Delete(&model.Student{}, "id = ?", id).Error
}
