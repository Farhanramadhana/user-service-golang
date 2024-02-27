package repository

import (
	"context"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, userData UserTable) (id int, err error) {
	query := `
	INSERT INTO users (full_name, phone_number, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	err = r.Db.QueryRowContext(
		ctx,
		query,
		userData.FullName,
		userData.PhoneNumber,
		userData.Password,
		userData.CreatedAt,
		userData.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *Repository) GetUserByPhone(ctx context.Context, phoneNumber string) (output UserTable, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number, password FROM users WHERE phone_number = $1", phoneNumber).Scan(&output.Id, &output.FullName, &output.PhoneNumber, &output.Password)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (output UserTable, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number, created_at, updated_at FROM users WHERE id = $1", id).Scan(&output.Id, &output.FullName, &output.PhoneNumber, &output.CreatedAt, &output.UpdatedAt)
	if err != nil {
		return
	}
	return
}

func (r *Repository) UpsertLoginLog(ctx context.Context, userID int, loginTime time.Time) error {
	var count int
	err := r.Db.QueryRowContext(ctx, "SELECT COUNT(*) FROM login_logs WHERE user_id = $1", userID).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// If it's the first time login, insert a new record
		_, err = r.Db.ExecContext(ctx, "INSERT INTO login_logs (user_id, total_login_success, last_login_at, created_at, updated_at) VALUES ($1, 1, $2, $3, $4)",
			userID, loginTime, loginTime, loginTime)
	} else {
		// update the existing record
		_, err = r.Db.ExecContext(ctx, "UPDATE login_logs SET total_login_success = total_login_success + 1, last_login_at = $1, updated_at = $2 WHERE user_id = $3",
			loginTime, loginTime, userID)
	}

	return err
}

func (r *Repository) UpdateUser(ctx context.Context, user UserTable) error {
	query := `
		UPDATE users
		SET full_name = $1, phone_number = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.Db.ExecContext(ctx, query, user.FullName, user.PhoneNumber, user.UpdatedAt, user.Id)
	if err != nil {
		return err
	}

	return nil
}
