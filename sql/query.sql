-- name: ListAvailableAppointments :many
SELECT * FROM list_available_appointments(@trainer_id, @starts_at, @ends_at);

-- name: PostNewAppointment :one
INSERT INTO appointments(trainer_id, user_id, starts_at, ends_at) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: ListTrainerScheduledAppointments :many
SELECT * FROM appointments WHERE trainer_id = $1;