package models

type Patron struct {
	SubjectNumber string `db:"subjectnumber" json:"subjectNumber"`
	FirstName     string `db:"firstname" json:"firstName"`
	LastName      string `db:"lastname" json:"lastName"`
}
