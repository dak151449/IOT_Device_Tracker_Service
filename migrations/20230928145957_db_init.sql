-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    salt TEXT NOT NULL,
    email TEXT NOT NULL,
    info TEXT NULL DEFAULT 'Not specified'
);

CREATE TABLE device_groups (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    user_id INT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_device_groups_users FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE devices (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    device_group_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    authentication_token TEXT NULL,
    status TEXT NOT NULL,
    last_activity_time TIMESTAMP NULL,
    location_name TEXT NOT NULL,
    location_latitude DECIMAL NULL,
    location_longitude DECIMAL NULL,
    location_timestamp TIMESTAMP NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_devices_device_groups FOREIGN KEY (device_group_id)
    REFERENCES device_groups(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE parameters (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    device_id INT NOT NULL,


    CONSTRAINT FK_parameters_devices FOREIGN KEY (device_id)
    REFERENCES devices(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE data_frames (
    id SERIAL PRIMARY KEY NOT NULL,
    parameter_id INT NOT NULL,
    event_time TIMESTAMP NOT NULL,
    result TEXT NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_data_frames_parameters FOREIGN KEY (parameter_id)
    REFERENCES Parameters(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE sent_data (
    id SERIAL PRIMARY KEY NOT NULL,
    parameter_id INT NOT NULL,
    event_time TIMESTAMP NOT NULL,
    value TEXT NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_sent_data_parameters FOREIGN KEY (parameter_id)
    REFERENCES parameters(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE sent_data;
DROP TABLE data_frames;
DROP TABLE parameters;
DROP TABLE devices;
DROP TABLE device_groups;
DROP TABLE users;
-- +goose StatementEnd