CREATE TABLE IF NOT EXISTS tags(
   id SERIAL PRIMARY KEY,
   tagName VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);