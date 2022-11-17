package server

import (
	"errors"
)

func UpdateHeightAndWeight(id int, height int, weight int) error {
	var err error
	if height == 0 && weight != 0 {
		_, err = db.Exec(`
		UPDATE Client
		SET weight = $1 WHERE
		(subscription_id = $2)`,
			weight, id)
		return errors.New("error: zero-value for height had been inputed")
	} else if weight == 0 && height != 0 {
		_, err = db.Exec(`
		UPDATE Client
		SET height = $1 WHERE
		(subscription_id = $2)`,
			height, id)
		return errors.New("error: zero-value for weight had been inputed")
	} else if height == 0 && weight == 0 {
		return errors.New("error: zero-values for weight and height had been inputed")
	}
	_, err = db.Exec(`
		UPDATE Client
		SET height = $1,
		weight = $2 WHERE
		(subscription_id = $3)`,
		height, weight, id)
	return err
}
