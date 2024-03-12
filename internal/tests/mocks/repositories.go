package mocks

import (
	"context"

	"github.com/bryanljx/go-rest-api/internal/models"
	"github.com/bryanljx/go-rest-api/internal/repositories"
)

// Student Repository

type mockStudentRepo struct{}

func (m mockStudentRepo) SuspendStudent(ctx context.Context, studentEmail string) error {
	return nil
}

func (m mockStudentRepo) ListSuspendedStudents(ctx context.Context) ([]models.Student, error) {
	res := []models.Student{
		models.Student{
			Person:      models.Person{ID: 1, Email: "JohnDoe@gmail.com", Metadata: models.Metadata{}},
			IsSuspended: true,
		},
		models.Student{
			Person:      models.Person{ID: 2, Email: "Guy123@gmail.com", Metadata: models.Metadata{}},
			IsSuspended: true,
		},
		models.Student{
			Person:      models.Person{ID: 4, Email: "Hello@gmail.com", Metadata: models.Metadata{}},
			IsSuspended: true,
		},
	}

	return res, nil
}

func MockStudentRepository() repositories.StudentRepository {
	return mockStudentRepo{}
}

// Teacher Repository

type mockTeacherRepo struct{}

func (m mockTeacherRepo) RegisterStudents(ctx context.Context, teacherEmail string, studentEmails []string) error {
	return nil
}

func (m mockTeacherRepo) RetrieveRegisteredStudents(ctx context.Context, teacher string) ([]string, error) {
	return mockEmails(), nil
}

func (m mockTeacherRepo) CommonStudents(ctx context.Context, teacherEmails []string) ([]string, error) {
	return mockEmails(), nil
}

func MockTeacherRepository() repositories.TeacherRepository {
	return mockTeacherRepo{}
}

// Helper functions

func mockEmails() []string {
	return []string{
		"commonstudent1@gmail.com",
		"commonstudent2@gmail.com",
		"student_only_under_teacher_ken@gmail.com",
		"studentmary@gmail.com",
		"studentbob@gmail.com",
		"studentagnes@gmail.com",
		"studentmiche@gmail.com",
	}
}
