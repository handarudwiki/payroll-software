-- +goose Up
-- +goose StatementBegin
CREATE TABLE positions(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    base_salary NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE positions;
-- +goose StatementEnd
