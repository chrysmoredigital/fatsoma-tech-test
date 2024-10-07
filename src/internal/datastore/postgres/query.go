package postgres

const (
	queryGetTicketOption = `
	SELECT "id", "name", "desc", "allocation"
	  FROM ticket_options
	 WHERE id = $1	
	`

	queryCreateTicketOption = `
	INSERT INTO ticket_options("name", "desc", "allocation", "created_at", "updated_at")
	     VALUES ($1, $2, $3, current_timestamp, current_timestamp)
	  RETURNING ("id")
	`

	queryUpdateTicketOptionPurchase = `
	UPDATE ticket_options
	   SET (allocation, updated_at) = (allocation - $1, current_timestamp)
	 WHERE id = $2
	`

	queryCreatePurchase = `
	INSERT INTO purchases("quantity", "user_id", "ticket_option_id", "created_at", "updated_at")
	     VALUES ($1, $2, $3, current_timestamp, current_timestamp)
	  RETURNING ("id")
	`

	queryCreateTicket = `
	INSERT INTO tickets("ticket_option_id", "purchase_id", "created_at", "updated_at")
	     VALUES ($1, $2, current_timestamp, current_timestamp)
	`
)
