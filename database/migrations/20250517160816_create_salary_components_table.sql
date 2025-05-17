-- +goose Up
-- +goose StatementBegin
CREATE TYPE salary_component_type AS ENUM ('allowance', 'deduction');

CREATE TABLE salary_components (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type salary_component_type NOT NULL DEFAULT 'allowance',
    is_recurring BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE salary_components;
DROP TYPE salary_component_type;
-- +goose StatementEnd
