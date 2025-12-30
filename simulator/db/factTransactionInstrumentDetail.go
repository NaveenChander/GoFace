package db

import (
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
)

// InsertFactTransactionInstrumentDetail inserts a new instrument detail record.
func InsertFactTransactionInstrumentDetail(detail models.FactTransactionInstrumentDetail) error {

	const query = `
		INSERT INTO facttransactioninstrumentdetail (
			transactioninstrumentsk,
			detailsk,
			value
		)
		VALUES ($1, $2, $3);
	`

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query,
		detail.TransactionInstrumentSK,
		detail.DetailSK,
		detail.Value,
	)

	if err != nil {
		return fmt.Errorf("failed to insert facttransactioninstrumentdetail: %w", err)
	}
	return nil
}
