package service

import "daily-cute-dogs-email/backend/api/model"

func CreateSubscribe(email string) error {
	if err := model.CreateSubscribe(email); err != nil {
		return err
	}
	return nil
}