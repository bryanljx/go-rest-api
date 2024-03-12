package repositories

import (
	"context"
	"log"
	"time"

	"github.com/bryanljx/go-rest-api/internal/models"
	"github.com/jmoiron/sqlx"
)

const (
	ErrStudentAlreadySuspended = "student is already suspended"
)

type StudentRepository interface {
	SuspendStudent(ctx context.Context, studentEmail string) error
	ListSuspendedStudents(ctx context.Context) ([]models.Student, error)
}

type studentRepo struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepo{
		db: db,
	}
}

func (sr studentRepo) SuspendStudent(ctx context.Context, studentEmail string) error {
	tx, err := sr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	var student models.Student
	getStudentQuery := `SELECT * FROM students s WHERE s.email = $1;`
	err = tx.GetContext(ctx, &student, getStudentQuery, studentEmail)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if student.IsSuspended {
		// Is is an error in this case if student is alr suspended?
		// return errorresponse.NewError(ErrStudentAlreadySuspended, http.StatusPreconditionFailed)

		tx.Commit()
		return nil
	}

	suspendStudentQuery := `UPDATE students s
		SET is_suspended = true,
			updated_at = $1
		WHERE s.email = $2;`

	_, err = tx.ExecContext(ctx, suspendStudentQuery, time.Now(), studentEmail)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	err = tx.Commit()
	return err
}

func (sr studentRepo) ListSuspendedStudents(ctx context.Context) ([]models.Student, error) {
	var suspendedStudents []models.Student
	query := `SELECT * FROM students s WHERE s.is_suspended IS TRUE;`

	err := sr.db.SelectContext(ctx, &suspendedStudents, query)
	if err != nil {
		return nil, err
	}

	return suspendedStudents, nil
}
