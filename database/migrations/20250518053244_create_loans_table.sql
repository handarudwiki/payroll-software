-- +goose Up
-- +goose StatementBegin
CREATE TYPE loan_type AS ENUM ('active', 'paid', 'cancelled');

CREATE TABLE loans (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    total_amount DECIMAL(10, 2) NOT NULL,
    monthly_installment DECIMAL(10, 2) NOT NULL,
    remaining_amount DECIMAL(10, 2) NOT NULL,
    start_date DATE NOT NULL,
    status loan_type NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE loans;
DROP TYPE loan_type;
-- +goose StatementEnd
