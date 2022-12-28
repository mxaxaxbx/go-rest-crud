DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id 		   VARCHAR(32) PRIMARY KEY  NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    email 	   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id 		   VARCHAR(32) PRIMARY KEY  NOT NULL,
    post_content VARCHAR(32) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id VARCHAR(32) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
