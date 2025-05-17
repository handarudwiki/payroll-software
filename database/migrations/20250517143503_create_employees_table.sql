-- +goose Up
-- +goose StatementBegin

CREATE TYPE employee_status AS ENUM ('active', 'inactive', 'on_leave');

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    nik VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    position_id INT REFERENCES positions(id),
    department_id INT REFERENCES departments(id),
    hire_date DATE NOT NULL,
    user_id INT REFERENCES users(id),
    status employee_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE employees;

DROP TYPE employee_status;

-- +goose StatementEnd
