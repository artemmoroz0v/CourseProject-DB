package server

func DeleteClient(id int) error {
	_, err := db.Exec(`
		DELETE FROM Client
		WHERE (subscription_id = $1)`,
		id)
	return err
}
