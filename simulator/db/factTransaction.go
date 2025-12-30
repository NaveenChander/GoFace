package db

import (
	"fmt"
	"time"

	"github.com/NaveenChander/GoFace/simulator/models"
)

// InsertFactTransaction inserts a new transaction record into the facttransaction table.
func InsertFactTransaction(txn models.FactTransaction) error {

	// Set default inserted datetime if not present by db default, but here we just pass the main fields
	const query = `
		INSERT INTO facttransaction (
			sourcetransactionnumber,
			patronid,
			locationsk,
			transactionutcdatetime,
			gamingdatesk,
			transactiondatesk,
			transactiontypesk,
			iscarded,
			issenttoincident
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	// Ensure UTC
	txnUTC := txn.TransactionUTCDateTime.UTC()

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query,
		txn.SourceTransactionNumber,
		txn.PatronID,
		txn.LocationSK,
		txnUTC,
		txn.GamingDateSK,
		txn.TransactionDateSK,
		txn.TransactionTypeSK,
		txn.IsCarded,
		txn.IsSentToIncident,
	)

	if err != nil {
		return fmt.Errorf("failed to insert facttransaction: %w", err)
	}
	return nil
}

// Helper to get current UTC time if needed
func NowUTC() time.Time {
	return time.Now().UTC()
}
