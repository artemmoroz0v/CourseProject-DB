package server

func DeleteGroup(id int) error {
	_, err := db.Exec(`
		DELETE FROM FC_Group
		WHERE (group_id = $1)`,
		id)
	return err
}
