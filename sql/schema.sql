ALTER DATABASE futuredb SET timezone TO 'America/Los_Angeles';

CREATE TABLE appointments (
    id SERIAL NOT NULL,
    user_id INT NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL,
    trainer_id INT NOT NULL
);

CREATE VIEW appointment_slot AS
    SELECT starts_at FROM appointments;

CREATE FUNCTION list_available_appointments(_trainer_id INT, _starts_at TIMESTAMPTZ, _ends_at TIMESTAMPTZ) RETURNS SETOF appointment_slot AS $$
DECLARE
    SATURDAY INTEGER := 6;
    SUNDAY INTEGER := 7;
BEGIN
    RETURN QUERY SELECT starts_at FROM generate_series(_starts_at, _ends_at - interval '1 minute', INTERVAL '30 minutes') starts_at WHERE 
    starts_at::TIME >= '08:00:00' AND starts_at::TIME < '17:00:00' AND EXTRACT(ISODOW FROM starts_at) NOT IN (SATURDAY, SUNDAY)
    EXCEPT SELECT starts_at FROM appointments ORDER BY starts_at ASC;
    RETURN;
END;
$$ LANGUAGE plpgsql;