package db

import (
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
)

// InsertFactTransactionSession inserts a new session record.
func InsertFactTransactionSession(session models.FactTransactionSession) error {

	const query = `
		INSERT INTO facttransactionsession (
			subjectsk,
			propertysk,
			workareask,
			locationsk,
			gamingdatesk,
			sessionstartutcdatetime,
			sessionendutcdatetime,
			cashin,
			cashout,
			noncashin,
			noncashout,
			cashincount,
			totalcount,
			avggapseconds,
			mingapseconds,
			turnover,
			winamount,
			cashbox,
			spincount,
			username,
			clientname
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);
	`

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query,
		session.SubjectSK,
		session.PropertySK,
		session.WorkAreaSK,
		session.LocationSK,
		session.GamingDateSK,
		session.SessionStartUTCDateTime,
		session.SessionEndUTCDateTime,
		session.CashIn,
		session.CashOut,
		session.NonCashIn,
		session.NonCashOut,
		session.CashInCount,
		session.TotalCount,
		session.AvgGapSeconds,
		session.MinGapSeconds,
		session.Turnover,
		session.WinAmount,
		session.CashBox,
		session.SpinCount,
		session.Username,
		session.ClientName,
	)

	if err != nil {
		return fmt.Errorf("failed to insert facttransactionsession: %w", err)
	}
	return nil
}
