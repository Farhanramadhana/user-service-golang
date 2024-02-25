package repository

import "context"

func (r *Repository) GetUserByPhone(ctx context.Context, phoneNumber string) (output UserTable, err error){
	err = r.Db.QueryRowContext(ctx, "SELECT phone FROM user WHERE id = $1", phoneNumber).Scan(&output)
	if err != nil {
		return
	}
	return
}
