CREATE TABLE follow_requests (
    follow_id uuid NOT NULL,
    sender_id uuid NOT NULL,
    receiver_id uuid NOT NULL,
    request_date DATE NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES users(user_id)
);