package models

import "time"

type FactTransactionInstrumentDenomination struct {
	TransactionInstrumentDenominationSK int64     `db:"transactioninstrumentdenominationsk" json:"transactionInstrumentDenominationSk"`
	TransactionInstrumentSK             int64     `db:"transactioninstrumentsk" json:"transactionInstrumentSk"`
	DenominationValue                   int       `db:"denominationvalue" json:"denominationValue"`
	DenominationCount                   int       `db:"denominationcount" json:"denominationCount"`
	InsertedDateTime                    time.Time `db:"inserteddatetime" json:"insertedDateTime"`
}
