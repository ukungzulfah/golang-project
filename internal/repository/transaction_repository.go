package repository

import (
	"database/sql"
	"errors"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) InsertUserAndMenu(
	userName string,
	fullName string,
	userPassword string,
	menuName string,
	menuURL string,
) error {

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(`
		INSERT INTO tb_user (user_name, user_fullname, user_password, user_menu)
		VALUES (?, ?, ?, ?)
	`, userName, fullName, userPassword, 0)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to insert user data")
	}

	result, err := tx.Exec(`
		INSERT INTO tb_menu (menu_name, menu_url)
		VALUES (?, ?)
	`, menuName, menuURL)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to insert menu data")
	}

	menuID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return errors.New("failed to get menu ID")
	}

	_, err = tx.Exec(`
		UPDATE tb_user 
		SET user_menu = ?
		WHERE user_name = ?
	`, menuID, userName)
	if err != nil {
		tx.Rollback()
		return errors.New("failed to update user menu")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
