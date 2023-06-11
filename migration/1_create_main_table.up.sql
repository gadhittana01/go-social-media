CREATE TABLE IF NOT EXISTS users(
   id SERIAL PRIMARY KEY,
   fullname VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts(
   id SERIAL PRIMARY KEY,
   userID INT NOT NULL REFERENCES users(id),
   title VARCHAR NOT NULL,
   description VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tags(
   id SERIAL PRIMARY KEY,
   tagName VARCHAR NOT NULL,
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS post_tags(
   id SERIAL PRIMARY KEY,
   postID INT NOT NULL REFERENCES posts(id),
   tagID INT NOT NULL REFERENCES tags(id),
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);