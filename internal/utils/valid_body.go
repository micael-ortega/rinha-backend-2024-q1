package utils

import (
	"fmt"

	"github.com/micael-ortega/crebito/internal/dto/request"
)

func CheckValidBody(body *request.TransactionRequest) error {
	if len(body.Description) == 0 || len(body.Description) > 10 {
		return fmt.Errorf("description length must be between 1 and 10 characters")
	}

	if body.Value <= 0 {
		return fmt.Errorf("value must be greater than zero")
	}

	if len(body.Kind) < 1 {
		return fmt.Errorf("kind length must be 1 character")
	}

	if body.Kind != "c" && body.Kind != "d" {
		return fmt.Errorf("kind must be 'c' or 'd'")
	}

	return nil
}
