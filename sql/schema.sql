ALTER DATABASE futuredb SET timezone TO 'America/Los_Angeles';

-- For the purpose of brevity, tables 'user' and 'trainer' have the minimally required data
CREATE TABLE member (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE trainer (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE appointments (
    id SERIAL NOT NULL,
    user_id UUID REFERENCES member(id) NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL CONSTRAINT opening CHECK (starts_at::TIME >= '08:00:00'),
    ends_at TIMESTAMPTZ NOT NULL CONSTRAINT closing CHECK (ends_at::TIME < '17:00:00'),
    trainer_id UUID REFERENCES trainer(id) NOT NULL,
    CONSTRAINT duration CHECK (AGE(ends_at, starts_at) = interval '30 minutes'),
    CONSTRAINT weekday CHECK (EXTRACT(ISODOW FROM starts_at) NOT IN (6, 7))
);

CREATE VIEW appointment_slot AS
    SELECT starts_at FROM appointments;

CREATE FUNCTION list_available_appointments(_trainer_id UUID, _starts_at TIMESTAMPTZ, _ends_at TIMESTAMPTZ) RETURNS SETOF appointment_slot AS $$
DECLARE
    SATURDAY INTEGER := 6;
    SUNDAY INTEGER := 7;
BEGIN
    IF NOT EXISTS(SELECT FROM trainer WHERE id = _trainer_id) THEN
        RAISE EXCEPTION 'Trainer UUID does not exist.';
    END IF;
    RETURN QUERY SELECT starts_at FROM generate_series(_starts_at, _ends_at - interval '1 minute', INTERVAL '30 minutes') starts_at WHERE 
    starts_at::TIME >= '08:00:00' AND starts_at::TIME < '17:00:00' AND EXTRACT(ISODOW FROM starts_at) NOT IN (SATURDAY, SUNDAY)
    EXCEPT SELECT starts_at FROM appointments WHERE trainer_id = _trainer_id ORDER BY starts_at ASC;
    RETURN;
END;
$$ LANGUAGE plpgsql;