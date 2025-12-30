package models

import "time"

type FactTransactionInstrumentDetail struct {
	TransactionInstrumentDetailSK int64     `db:"transactioninstrumentdetailsk" json:"transactionInstrumentDetailSk"`
	TransactionInstrumentSK       int64     `db:"transactioninstrumentsk" json:"transactionInstrumentSk"`
	DetailSK                      int       `db:"detailsk" json:"detailSk"`
	Value                         string    `db:"value" json:"value"`
	InsertedDateTime              time.Time `db:"inserteddatetime" json:"insertedDateTime"`
}
