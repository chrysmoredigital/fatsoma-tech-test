package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/chrysmoredigital/fatsoma-tech-test/internal/model"
	"github.com/lib/pq"
)

const (
	codeCheckViolation = "23514"
)

func (c *Client) CreatePurchase(p model.Purchase) error {
	var err error
	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start tx: %w", err)
	}
	defer func() {
		if err != nil {
			log.Printf("%s: rolling back tx", err)
			if err := tx.Rollback(); err != nil {
				log.Printf("failed to rollback tx: %s", err)
			}
			return
		}
	}()

	// update ticket_options, fails if quantity exceeds allocation
	q := queryUpdateTicketOptionPurchase

	_, err = tx.Exec(q, p.Quantity, p.TicketOptionID)
	if err != nil {
		pqerr := &pq.Error{}
		errors.As(err, &pqerr)
		exceedsAlloc := pqerr.Code == codeCheckViolation

		if exceedsAlloc {
			err = errors.New("quantity exceeds allocation")
		}

		return &er{
			err:         fmt.Errorf("error updating ticket option: %w", err),
			clientError: exceedsAlloc,
		}
	}

	// create purchase
	q = queryCreatePurchase

	row := tx.QueryRow(q, p.Quantity, p.UserID, p.TicketOptionID)
	if err = row.Scan(&p.ID); err != nil {
		return fmt.Errorf("error creating purchase: %w", err)
	}

	// create tickets
	// Not sure if this was outside of the scope of the exercise?
	q = queryCreateTicket
	for range p.Quantity {
		_, err = tx.Exec(q, p.TicketOptionID, p.ID)
		if err != nil {
			return fmt.Errorf("error creating ticket: %w", err)
		}

	}

	// won't be caught by deferred rollback on error
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}

	return nil
}
