package models

import "time"

type FactTransactionSession struct {
	TransactionSessionSK    int64      `db:"transactionsessionsk" json:"transactionSessionSk"`
	SubjectSK               *int64     `db:"subjectsk" json:"subjectSk"`
	PropertySK              int        `db:"propertysk" json:"propertySk"`
	WorkAreaSK              *int       `db:"workareask" json:"workAreaSk"`
	LocationSK              *int       `db:"locationsk" json:"locationSk"`
	GamingDateSK            int        `db:"gamingdatesk" json:"gamingDateSk"`
	SessionStartUTCDateTime *time.Time `db:"sessionstartutcdatetime" json:"sessionStartUtcDateTime"`
	SessionEndUTCDateTime   *time.Time `db:"sessionendutcdatetime" json:"sessionEndUtcDateTime"`
	CashIn                  *float64   `db:"cashin" json:"cashIn"`
	CashOut                 *float64   `db:"cashout" json:"cashOut"`
	NonCashIn               *float64   `db:"noncashin" json:"nonCashIn"`
	NonCashOut              *float64   `db:"noncashout" json:"nonCashOut"`
	CashInCount             int        `db:"cashincount" json:"cashInCount"`
	TotalCount              int        `db:"totalcount" json:"totalCount"`
	AvgGapSeconds           *int       `db:"avggapseconds" json:"avgGapSeconds"`
	MinGapSeconds           *int       `db:"mingapseconds" json:"minGapSeconds"`
	Turnover                float64    `db:"turnover" json:"turnover"`
	WinAmount               float64    `db:"winamount" json:"winAmount"`
	CashBox                 float64    `db:"cashbox" json:"cashBox"`
	SpinCount               int        `db:"spincount" json:"spinCount"`
	InsertedDateTime        time.Time  `db:"inserteddatetime" json:"insertedDateTime"`
	Username                string     `db:"username" json:"username"`
	ClientName              string     `db:"clientname" json:"clientName"`
}
