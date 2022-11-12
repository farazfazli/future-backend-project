// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package futuredb

import (
	"context"
	"time"
)

const listAvailableAppointments = `-- name: ListAvailableAppointments :many
SELECT starts_at FROM list_available_appointments($1, $2, $3)
`

type ListAvailableAppointmentsParams struct {
	TrainerID int32     `json:"trainer_id"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
}

func (q *Queries) ListAvailableAppointments(ctx context.Context, arg ListAvailableAppointmentsParams) ([]time.Time, error) {
	rows, err := q.db.QueryContext(ctx, listAvailableAppointments, arg.TrainerID, arg.StartsAt, arg.EndsAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []time.Time
	for rows.Next() {
		var starts_at time.Time
		if err := rows.Scan(&starts_at); err != nil {
			return nil, err
		}
		items = append(items, starts_at)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTrainerScheduledAppointments = `-- name: ListTrainerScheduledAppointments :many
SELECT id, user_id, starts_at, ends_at, trainer_id FROM appointments WHERE trainer_id = $1
`

func (q *Queries) ListTrainerScheduledAppointments(ctx context.Context, trainerID int32) ([]Appointment, error) {
	rows, err := q.db.QueryContext(ctx, listTrainerScheduledAppointments, trainerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Appointment
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.StartsAt,
			&i.EndsAt,
			&i.TrainerID,
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

const postNewAppointment = `-- name: PostNewAppointment :one
INSERT INTO appointments(trainer_id, user_id, starts_at, ends_at) VALUES ($1, $2, $3, $4) RETURNING id
`

type PostNewAppointmentParams struct {
	TrainerID int32     `json:"trainer_id"`
	UserID    int32     `json:"user_id"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
}

func (q *Queries) PostNewAppointment(ctx context.Context, arg PostNewAppointmentParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, postNewAppointment,
		arg.TrainerID,
		arg.UserID,
		arg.StartsAt,
		arg.EndsAt,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}