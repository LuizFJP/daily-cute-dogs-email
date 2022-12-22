package model

import "daily-cute-dogs-email/backend/db"

func CreateSubscribe(email string) error {
	mdb := db.Start()
	if err := mdb.AddEmail(email); err != nil {
		return err
	}
	return nil
}