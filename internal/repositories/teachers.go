package repositories

import (
	"context"

	"github.com/bryanljx/go-rest-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const ()

type TeacherRepository interface {
	RegisterStudents(ctx context.Context, teacherEmail string, studentEmails []string) error
	RetrieveRegisteredStudents(ctx context.Context, teacher string) ([]string, error)
	CommonStudents(ctx context.Context, teacherEmails []string) ([]string, error)
}

type teacherRepo struct {
	db *sqlx.DB
}

func NewTeacherRepository(db *sqlx.DB) TeacherRepository {
	return teacherRepo{
		db: db,
	}
}

func (tr teacherRepo) CommonStudents(ctx context.Context, teacherEmails []string) ([]string, error) {
	var commonStudents []models.Student

	query := `SELECT s.email FROM students s
		JOIN teacher_student ts ON s.id = ts.student_id
		JOIN teachers t ON ts.teacher_id = t.id
		WHERE t.email IN (:emails)
		GROUP BY s.email
		HAVING COUNT(t.id) = :count;`

	m := map[string]interface{}{"count": len(teacherEmails), "emails": teacherEmails}

	query, args, err := sqlx.Named(query, m)
	if err != nil {
		return nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}

	query = tr.db.Rebind(query)

	err = tr.db.SelectContext(ctx, &commonStudents, query, args...)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(commonStudents))
	for _, s := range commonStudents {
		res = append(res, s.Email)
	}

	return res, nil
}

func (tr teacherRepo) RegisterStudents(ctx context.Context, teacherEmail string, studentEmails []string) error {
	query := `INSERT INTO teacher_student (teacher_id, student_id)
		SELECT t.id, s.id
		FROM teachers t, students s
		WHERE t.email = :teacher
		AND s.email IN (:student)
		AND NOT EXISTS (SELECT ts.teacher_id, ts.student_id FROM teacher_student ts);`

	arg := map[string]interface{}{"teacher": teacherEmail, "student": studentEmails}

	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return err
	}

	query = tr.db.Rebind(query)

	_, err = tr.db.ExecContext(ctx, query, args...)
	return err
}

func (tr teacherRepo) RetrieveRegisteredStudents(ctx context.Context, teacherEmail string) ([]string, error) {
	var registeredStudents []models.Student
	query := `SELECT s.email
		FROM teachers t
		JOIN teacher_student ts on t.id = ts.teacher_id
		JOIN students s on ts.student_id = s.id
		WHERE t.email = $1
		AND s.is_suspended IS FALSE;`

	err := tr.db.SelectContext(ctx, &registeredStudents, query, teacherEmail)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(registeredStudents))
	for _, s := range registeredStudents {
		res = append(res, s.Email)
	}

	return res, nil
}
