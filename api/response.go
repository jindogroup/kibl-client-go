package api

import (
	"time"

	"github.com/pkg/errors"
)

type Response[T any] struct {
	Code        int       `json:"code,omitempty"`
	Description string    `json:"description:,omitempty"`
	APIKey      string    `json:"api_key,omitempty"`
	Result      []T       `json:"result,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (r *Response[T]) Validate() error {
	if r.Code != 200 {
		return errors.Errorf("%s failed with status code %d", r.APIKey, r.Code)
	}
	return nil
}

func (r *Response[T]) Empty() bool {
	return len(r.Result) == 0
}

