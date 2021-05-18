-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Berlin";

-- Create books table
CREATE TABLE signal_verifications (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    moodle_token VARCHAR (32) NOT NULL UNIQUE,
    phone_number VARCHAR (16) NOT NULL,
    verification_code INTEGER NOT NULL,
    valid_until TIMESTAMP NOT NULL
);

CREATE TABLE assignments (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    moodle_token VARCHAR (32) NOT NULL UNIQUE,
    phone_number VARCHAR (16) NOT NULL,
    assignments json DEFAULT '[]'::json
);