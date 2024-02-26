// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
	"time"
)

type RepositoryInterface interface {
	CreateUser(ctx context.Context, userData UserTable) (id int, err error)
	GetUserByPhone(ctx context.Context, phoneNumber string) (output UserTable, err error)
	GetUserByID(ctx context.Context, id int) (output UserTable, err error)
	UpsertLoginLog(ctx context.Context, userID int, loginTime time.Time) error
}
