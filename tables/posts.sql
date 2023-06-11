CREATE TABLE IF NOT EXISTS posts(
   id SERIAL PRIMARY KEY,
   userID INT NOT NULL,
   title VARCHAR NOT NULL,
   description VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);