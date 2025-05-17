-- +goose Up
-- +goose StatementBegin

CREATE TYPE attendance_status AS ENUM ('present', 'absent', 'sick', 'late', 'on_leave');

CREATE TABLE attendances (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES employees(id),
    date DATE NOT NULL,
    status attendance_status NOT NULL DEFAULT 'present',
    working_hours INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE attendances;
DROP TYPE attendance_status;
-- +goose StatementEnd
