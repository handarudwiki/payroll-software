-- +goose Up
-- +goose StatementBegin
CREATE TABLE employee_components (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    salary_component_id INT REFERENCES salary_components(id),
    amount DECIMAL(10, 2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    custom_override BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employee_components;
-- +goose StatementEnd
