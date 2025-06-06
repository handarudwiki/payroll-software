-- +goose Up
-- +goose StatementBegin
CREATE TABLE payslip_details(
    id SERIAL PRIMARY KEY,
    payroll_id INT REFERENCES payrolls(id),
    salary_component_id INT REFERENCES salary_components(id),
    component_type salary_component_type NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payslip_details;
-- +goose StatementEnd
