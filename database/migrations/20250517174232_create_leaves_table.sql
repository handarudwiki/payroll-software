-- +goose Up
-- +goose StatementBegin
CREATE TYPE type_status AS ENUM ('annual', 'sick', 'unpaid');
CREATE TYPE leave_status AS ENUM ('pending', 'approved', 'rejected');

CREATE TABLE leaves (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    type leave_status NOT NULL DEFAULT 'pending',
    status type_status NOT NULL DEFAULT 'annual',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE leaves;
DROP TYPE leave_status;
DROP TYPE type_status;
-- +goose StatementEnd
