-- +goose Up
-- +goose StatementBegin
ALTER TABLE payrolls
  ALTER COLUMN net_salary TYPE NUMERIC(15,2),
  ALTER COLUMN base_salary TYPE NUMERIC(15,2),
  ALTER COLUMN total_allowances TYPE NUMERIC(15,2),
  ALTER COLUMN total_deductions TYPE NUMERIC(15,2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- ⚠️ WARNING: This reverts columns to NUMERIC(10,2). Adjust based on original types.
ALTER TABLE payrolls
  ALTER COLUMN net_salary TYPE NUMERIC(10,2),
  ALTER COLUMN base_salary TYPE NUMERIC(10,2),
  ALTER COLUMN total_allowances TYPE NUMERIC(10,2),
  ALTER COLUMN total_deductions TYPE NUMERIC(10,2);
-- +goose StatementEnd
