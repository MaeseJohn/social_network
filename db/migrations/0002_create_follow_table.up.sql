CREATE TABLE follows (
    follower_id INT NOT NULL,
    followed_id INT NOT NULL,
    follow_date DATE NOT NULL,
    PRIMARY KEY (follower_id, followed_id),
    FOREIGN KEY (follower_id) REFERENCES users(id),
    FOREIGN KEY (followed_id) REFERENCES users(id)
);