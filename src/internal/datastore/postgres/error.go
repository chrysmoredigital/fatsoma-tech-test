package postgres

import (
	"errors"
)

// er is the package's internal error type. Can
// be used for reflection on the type of error
// returned by postgres
type er struct {
	err         error
	clientError bool
}

func (e *er) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

// ClientError is a callback function that can be used to
// determine if the error returned by postgres equates
// to a client error (as opposed to server error)
func (c *Client) ClientError(err error) bool {
	e := &er{}
	return errors.As(err, &e) && e.clientError
}
