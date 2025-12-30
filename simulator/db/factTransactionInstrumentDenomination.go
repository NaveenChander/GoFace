package db

import (
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
)

// InsertFactTransactionInstrumentDenomination inserts a new denomination record.
func InsertFactTransactionInstrumentDenomination(denom models.FactTransactionInstrumentDenomination) error {

	const query = `
		INSERT INTO facttransactioninstrumentdenomination (
			transactioninstrumentsk,
			denominationvalue,
			denominationcount
		)
		VALUES ($1, $2, $3);
	`

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query,
		denom.TransactionInstrumentSK,
		denom.DenominationValue,
		denom.DenominationCount,
	)

	if err != nil {
		return fmt.Errorf("failed to insert facttransactioninstrumentdenomination: %w", err)
	}
	return nil
}
