-- +goose Up
-- +goose StatementBegin
CREATE TABLE payrolls(
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    period DATE NOT NULL,
    total_allowances DECIMAL(10, 2) NOT NULL,
    total_deductions DECIMAL(10, 2) NOT NULL,
    net_salary DECIMAL(10, 2) NOT NULL,
    generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    base_salary DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payrolls;
-- +goose StatementEnd
