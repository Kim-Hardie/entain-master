-- Add the 'status' field to the 'races' table
ALTER TABLE races
    ADD COLUMN status VARCHAR(10) NOT NULL DEFAULT 'OPEN';

-- Update the 'status' field based on 'advertised_start_time'
UPDATE races
SET status = CASE
                 WHEN datetime(advertised_start_time) > datetime('now') THEN 'OPEN'
                 ELSE 'CLOSED'
    END
WHERE status IS NULL;
