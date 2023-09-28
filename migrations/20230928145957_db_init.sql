-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    id SERIAL PRIMARY KEY NOT NULL,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    salt TEXT NOT NULL,
    email TEXT NOT NULL,
    last_login TIMESTAMP NULL,
    is_superuser BOOLEAN NULL DEFAULT FALSE,
    is_staff BOOLEAN NULL DEFAULT FALSE,
    info TEXT NULL DEFAULT 'Not specified'
);

CREATE TABLE DeviceGroups (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    user_id INT NOT NULL,
    status TEXT NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_DeviceGroups_Users FOREIGN KEY (user_id)
    REFERENCES Users(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE Devices (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    device_group_id INT NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    authentication_token TEXT NULL,
    status TEXT NOT NULL,
    last_activity_time TIMESTAMP NULL,
    location_name TEXT NOT NULL,
    location_latitude DECIMAL NULL,
    location_longitude DECIMAL NULL,
    location_timestamp TIMESTAMP NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_Devices_DeviceGroups FOREIGN KEY (device_group_id)
    REFERENCES DeviceGroups(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE Parameters (
    id SERIAL PRIMARY KEY NOT NULL,
    parameter_name TEXT NOT NULL,
    parameter_symbol TEXT NOT NULL,
    device_id INT NOT NULL,


    CONSTRAINT FK_Parameters_Devices FOREIGN KEY (device_id)
    REFERENCES Devices(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE DataFrames (
    id SERIAL PRIMARY KEY NOT NULL,
    parameter_id INT NOT NULL,
    event_time TIMESTAMP NOT NULL,
    result TEXT NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_DataFrames_Parameters FOREIGN KEY (parameter_id)
    REFERENCES Parameters(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);

CREATE TABLE SentData (
    id SERIAL PRIMARY KEY NOT NULL,
    parameter_id INT NOT NULL,
    event_time TIMESTAMP NOT NULL,
    value TEXT NOT NULL,
    description TEXT NULL DEFAULT 'Not specified',


    CONSTRAINT FK_SentData_Parameters FOREIGN KEY (parameter_id)
    REFERENCES Parameters(id)
    ON DELETE CASCADE
    ON UPDATE RESTRICT
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE SentData;
DROP TABLE DataFrames;
DROP TABLE Parameters;
DROP TABLE Devices;
DROP TABLE DeviceGroups;
DROP TABLE Users;
-- +goose StatementEnd