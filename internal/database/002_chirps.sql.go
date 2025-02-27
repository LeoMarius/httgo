// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 002_chirps.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createChirp = `-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING id, created_at, updated_at, body, user_id
`

type CreateChirpParams struct {
	Body   string
	UserID uuid.UUID
}

func (q *Queries) CreateChirp(ctx context.Context, arg CreateChirpParams) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, createChirp, arg.Body, arg.UserID)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const deleteAllChirps = `-- name: DeleteAllChirps :exec
DELETE FROM chirps
`

func (q *Queries) DeleteAllChirps(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllChirps)
	return err
}

const deleteChirpsByID = `-- name: DeleteChirpsByID :exec
DELETE
FROM chirps
WHERE id = $1
`

func (q *Queries) DeleteChirpsByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteChirpsByID, id)
	return err
}

const getAllChirpsASC = `-- name: GetAllChirpsASC :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
ORDER BY created_at ASC
`

func (q *Queries) GetAllChirpsASC(ctx context.Context) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getAllChirpsASC)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllChirpsDESC = `-- name: GetAllChirpsDESC :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
ORDER BY created_at DESC
`

func (q *Queries) GetAllChirpsDESC(ctx context.Context) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getAllChirpsDESC)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChirpsByID = `-- name: GetChirpsByID :one
SELECT id, created_at, updated_at, body, user_id
FROM chirps
WHERE id = $1
`

func (q *Queries) GetChirpsByID(ctx context.Context, id uuid.UUID) (Chirp, error) {
	row := q.db.QueryRowContext(ctx, getChirpsByID, id)
	var i Chirp
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Body,
		&i.UserID,
	)
	return i, err
}

const getChirpsByUserIDASC = `-- name: GetChirpsByUserIDASC :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC
`

func (q *Queries) GetChirpsByUserIDASC(ctx context.Context, userID uuid.UUID) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getChirpsByUserIDASC, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChirpsByUserIDDESC = `-- name: GetChirpsByUserIDDESC :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetChirpsByUserIDDESC(ctx context.Context, userID uuid.UUID) ([]Chirp, error) {
	rows, err := q.db.QueryContext(ctx, getChirpsByUserIDDESC, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chirp
	for rows.Next() {
		var i Chirp
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Body,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
