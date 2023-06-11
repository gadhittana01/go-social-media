CREATE TABLE IF NOT EXISTS post_tags(
   id SERIAL PRIMARY KEY,
   postID INT NOT NULL REFERENCES posts(id),
   tagID INT NOT NULL REFERENCES tags(id),
   created_at TIMESTAMP,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);