CREATE TABLE IF NOT EXISTS users(
   id SERIAL PRIMARY KEY,
   fullname VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);