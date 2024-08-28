CREATE TABLE posts (
   id INT AUTO_INCREMENT PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   user_id INT NOT NULL,
   forum_id INT NOT NULL,
   text TEXT,
   created_at DATETIME,
   updated_at DATETIME,
   FOREIGN KEY (user_id) REFERENCES users(id),
   FOREIGN KEY (forum_id) REFERENCES forums(id)
);