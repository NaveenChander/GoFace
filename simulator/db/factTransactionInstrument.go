package db

import (
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
)

// InsertFactTransactionInstrument inserts a new instrument record.
func InsertFactTransactionInstrument(inst models.FactTransactionInstrument) error {

	const query = `
		INSERT INTO facttransactioninstrument (
			transactionsk,
			instrumenttypesk,
			amountin,
			amountout
		)
		VALUES ($1, $2, $3, $4);
	`

	// InsertedDateTime is handled by DEFAULT now() in schema,
	// unless we want to override it. For now, trusting DB default.

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query,
		inst.TransactionSK,
		inst.InstrumentTypeSK,
		inst.AmountIn,
		inst.AmountOut,
	)

	if err != nil {
		return fmt.Errorf("failed to insert facttransactioninstrument: %w", err)
	}
	return nil
}

// GetMaxTransactionSK retrieves the max ID from facttransaction to use for testing helper if needed.
func GetMaxTransactionSK() (int64, error) {
	return 0, nil // Placeholder if needed later
}
