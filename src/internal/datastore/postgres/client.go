package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewClient(cfg *Config) (*Client, error) {
	var err error
	var client Client

	if client.db, err = sql.Open("postgres", cfg.String()); err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	for n := range cfg.Retries {
		if err := client.db.Ping(); err != nil {
			err = fmt.Errorf("failed to ping db: %w", err)
			if n < cfg.Retries {
				log.Printf("%s: retrying( %d )", err, n)
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, err
		}
		break
	}

	return &client, nil
}

type Client struct {
	db *sql.DB
}

func (c *Client) CloseDB() error {
	return c.db.Close()
}
