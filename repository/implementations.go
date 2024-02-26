package repository

import (
	"context"
)

func (r *Repository) CreateUser(ctx context.Context, userData UserTable) (err error) {
	query := `
	INSERT INTO users (full_name, phone_number, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	_, err = r.Db.ExecContext(
		ctx,
		query,
		userData.FullName,
		userData.PhoneNumber,
		userData.Password,
		userData.CreatedAt,
		userData.UpdatedAt,
	)

	return err
}

func (r *Repository) GetUserByPhone(ctx context.Context, phoneNumber string) (output UserTable, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number FROM users WHERE phone_number = $1", phoneNumber).Scan(&output.Id, &output.FullName, &output.PhoneNumber)
	if err != nil {
		return
	}
	return
}
