package server

func DeleteTrainer(id int) error {
	_, err := db.Exec(`
		DELETE FROM Trainer
		WHERE (trainer_id = $1)`,
		id)
	return err
}
