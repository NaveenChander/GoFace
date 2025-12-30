package models

type DIMDate struct {
	Datesk    int    `db:"datesk"`
	Fulldate  string `db:"fulldate"`
	Dayofweek int    `db:"dayofweek"`
	Month     int    `db:"month"`
	Quarter   int    `db:"quarter"`
	Year      int    `db:"year"`
}
