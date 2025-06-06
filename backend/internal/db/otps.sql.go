// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: otps.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteOTP = `-- name: DeleteOTP :exec
DELETE FROM otps WHERE user_pid =$1
`

func (q *Queries) DeleteOTP(ctx context.Context, userPid int64) error {
	_, err := q.db.Exec(ctx, deleteOTP, userPid)
	return err
}

const getOTP = `-- name: GetOTP :one
SELECT user_pid, created_at, secret_key, expiry
FROM otps WHERE user_pid = $1
`

func (q *Queries) GetOTP(ctx context.Context, userPid int64) (Otp, error) {
	row := q.db.QueryRow(ctx, getOTP, userPid)
	var i Otp
	err := row.Scan(
		&i.UserPid,
		&i.CreatedAt,
		&i.SecretKey,
		&i.Expiry,
	)
	return i, err
}

const insertOTP = `-- name: InsertOTP :exec
INSERT INTO otps (user_pid, created_at, secret_key, expiry)
VALUES ($1, $2, $3, $4)
ON CONFLICT (user_pid)
DO UPDATE
SET created_at = $2, secret_key = $3, expiry = $4
`

type InsertOTPParams struct {
	UserPid   int64
	CreatedAt pgtype.Timestamptz
	SecretKey string
	Expiry    pgtype.Timestamptz
}

func (q *Queries) InsertOTP(ctx context.Context, arg InsertOTPParams) error {
	_, err := q.db.Exec(ctx, insertOTP,
		arg.UserPid,
		arg.CreatedAt,
		arg.SecretKey,
		arg.Expiry,
	)
	return err
}
