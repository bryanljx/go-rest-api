package models

type Student struct {
	Person

	Teachers []Teacher

	IsSuspended bool `db:"is_suspended"`
}
