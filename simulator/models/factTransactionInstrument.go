package models

import "time"

type FactTransactionInstrument struct {
	TransactionInstrumentSK int64     `db:"transactioninstrumentsk" json:"transactionInstrumentSk"`
	TransactionSK           int64     `db:"transactionsk" json:"transactionSk"`
	InstrumentTypeSK        int       `db:"instrumenttypesk" json:"instrumentTypeSk"`
	AmountIn                *float64  `db:"amountin" json:"amountIn"`
	AmountOut               *float64  `db:"amountout" json:"amountOut"`
	InsertedDateTime        time.Time `db:"inserteddatetime" json:"insertedDateTime"`
}
