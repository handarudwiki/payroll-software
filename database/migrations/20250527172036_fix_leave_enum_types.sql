-- +goose Up
ALTER TABLE leaves RENAME COLUMN type TO type_old;
ALTER TABLE leaves RENAME COLUMN status TO status_old;

-- Corrected types
ALTER TABLE leaves ADD COLUMN type type_status NOT NULL DEFAULT 'annual';
ALTER TABLE leaves ADD COLUMN status leave_status NOT NULL DEFAULT 'pending';

-- Update with mapping only if your old data can be mapped safely
-- This part needs to be correct based on your actual previous usage.
-- But if the old data is invalid or mismatched, just skip the transfer and start clean.

-- If you're resetting values:
UPDATE leaves SET type = 'annual';
UPDATE leaves SET status = 'pending';

-- Drop the wrongly typed columns
ALTER TABLE leaves DROP COLUMN type_old;
ALTER TABLE leaves DROP COLUMN status_old;

-- +goose Down
ALTER TABLE leaves ADD COLUMN type_old leave_status NOT NULL DEFAULT 'pending';
ALTER TABLE leaves ADD COLUMN status_old type_status NOT NULL DEFAULT 'annual';

UPDATE leaves SET type_old = 'pending';
UPDATE leaves SET status_old = 'annual';

ALTER TABLE leaves DROP COLUMN type;
ALTER TABLE leaves DROP COLUMN status;

ALTER TABLE leaves RENAME COLUMN type_old TO type;
ALTER TABLE leaves RENAME COLUMN status_old TO status;
