-- 1. Create database
CREATE DATABASE albumdb;

-- 2. Connect to database
\c albumdb;

-- 3. Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 4. Create albums table
CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    artist VARCHAR(100) NOT NULL,
    price NUMERIC(10,2) NOT NULL CHECK (price > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
