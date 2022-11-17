package server

import (
	"errors"
)

func UpdateHeightAndWeight(id int, height int, weight int) error {
	if height == 0 {
		return errors.New("error: zero-value for height had been inputed")
	} else if weight == 0 {
		return errors.New("error: zero-value for weight had been inputed")
	}
	_, err := db.Exec(`
		UPDATE Client
		SET height = $1,
		weight = $2 WHERE
		(subscription_id = $3)`,
		height, weight, id)
	return err
}
