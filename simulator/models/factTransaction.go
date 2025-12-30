package models

import "time"

type FactTransaction struct {
	TransactionSK           int64      `db:"transactionsk" json:"transactionSk"`
	SourceTransactionNumber string     `db:"sourcetransactionnumber" json:"sourceTransactionNumber"`
	PatronID                *int64     `db:"patronid" json:"patronId"`
	LocationSK              *int       `db:"locationsk" json:"locationSk"`
	TransactionUTCDateTime  time.Time  `db:"transactionutcdatetime" json:"transactionUtcDateTime"`
	GamingDateSK            int        `db:"gamingdatesk" json:"gamingDateSk"`
	TransactionDateSK       int        `db:"transactiondatesk" json:"transactionDateSk"`
	TransactionTypeSK       int64      `db:"transactiontypesk" json:"transactionTypeSk"`
	IsCarded                bool       `db:"iscarded" json:"isCarded"`
	IsSentToIncident        bool       `db:"issenttoincident" json:"isSentToIncident"`
}
