package models

type Notifications struct {
	Metadata

	Message    string
	Sender     Teacher
	Recipients []Student
}
