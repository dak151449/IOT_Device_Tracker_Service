-- +goose Up
INSERT INTO users (username, password, salt, email, info)
SELECT
        'User ' || generate_series,
        'password' || generate_series,
        'salt' || generate_series,
        'user' || generate_series || '@test.com',
        'User info ' || generate_series
FROM generate_series(1, 10);

INSERT INTO device_groups (name, user_id, status, created_at, description)
SELECT
        'Device Group ' || generate_series,
        (SELECT id FROM users ORDER BY random()*generate_series LIMIT 1),
        CASE WHEN random() < 0.5 THEN 'Active' ELSE 'Inactive' END,
        NOW() - (floor(random() * 365) || ' days')::INTERVAL,
        'Description for group ' || generate_series
FROM generate_series(1, 100);

INSERT INTO devices (name, device_group_id, created_at, authentication_token, status, location_name, location_latitude, location_longitude, location_timestamp, description)
SELECT
        'Device ' || generate_series,
        (SELECT id FROM device_groups ORDER BY random()*generate_series LIMIT 1),
        NOW() - (floor(random() * 365) || ' days')::INTERVAL,
        md5(random()::TEXT),
        CASE WHEN random() < 0.5 THEN 'Active' ELSE 'Inactive' END,
        'Location ' || generate_series,
        (random() * 180 - 90),
        (random() * 360 - 180),
        NOW() - (floor(random() * 365) || ' days')::INTERVAL,
        'Description for device ' || generate_series
FROM generate_series(1, 100);

-- +goose Down
DELETE FROM users;