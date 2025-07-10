package service

import (
	"StudentService/cache"
	"StudentService/dao"
	"StudentService/model"
)

func ListStudents() ([]model.Student, error) {
	return dao.GetAllStudents()
}

func GetStudent(id string) (*model.Student, error) {
	student, err := cache.GetStudentFromCache(id)
	if err == nil {
		return student, nil
	}

	student, err = dao.GetStudentByID(id)
	if err != nil {
		return nil, err
	}
	cache.SetStudentToCache(student)
	return student, nil
}

func CreateStudent(student *model.Student) error {
	err := dao.CreateStudent(student)
	if err == nil {
		cache.SetStudentToCache(student)
	}
	return err
}

func UpdateStudent(student *model.Student) error {
	err := dao.UpdateStudent(student)
	if err == nil {
		cache.SetStudentToCache(student)
	}
	return err
}

func DeleteStudent(id string) error {
	err := dao.DeleteStudent(id)
	if err == nil {
		cache.DeleteStudentFromCache(id)
	}
	return err
}
