package postgres

import (
	"fmt"

	"github.com/chrysmoredigital/fatsoma-tech-test/internal/model"
	"github.com/google/uuid"
)

func (c *Client) GetTicketOption(id uuid.UUID) (*model.TicketOption, error) {
	q := queryGetTicketOption
	var to model.TicketOption

	row := c.db.QueryRow(q, id)
	if err := row.Scan(
		&to.ID,
		&to.Name,
		&to.Desc,
		&to.Allocation,
	); err != nil {
		return nil, fmt.Errorf("error scanning rows: %w", err)
	}

	return &to, nil
}

func (c *Client) CreateTicketOption(to model.TicketOption) (uuid.UUID, error) {
	q := queryCreateTicketOption

	row := c.db.QueryRow(q, to.Name, to.Desc, to.Allocation)
	if err := row.Scan(&to.ID); err != nil {
		return uuid.UUID{}, fmt.Errorf("error running query: %w", err)
	}

	return to.ID, nil
}
