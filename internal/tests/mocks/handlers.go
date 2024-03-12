package mocks

import (
	"github.com/bryanljx/go-rest-api/internal/handlers/notifications"
	"github.com/bryanljx/go-rest-api/internal/handlers/students"
	"github.com/bryanljx/go-rest-api/internal/handlers/teachers"
)

func MockNotificationHandler() notifications.NotificationHandler {
	studentRepo := MockStudentRepository()
	teacherRepo := MockTeacherRepository()
	s := MockService()
	return notifications.New(studentRepo, teacherRepo, &s)
}

func MockStudentHandler() students.StudentHandler {
	studentRepo := MockStudentRepository()
	s := MockService()
	return students.New(studentRepo, &s)
}

func MockTeacherHandler() teachers.TeacherHandler {
	teacherRepo := MockTeacherRepository()
	s := MockService()
	return teachers.New(teacherRepo, &s)
}
