CREATE TABLE follow_requests (
    sender_id uuid NOT NULL,
    receiver_id uuid NOT NULL,
    status TEXT NOT NULL,
    request_date DATE NOT NULL,
    procesed_date DATE,
    FOREIGN KEY (sender_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_id) REFERENCES users(user_id)
);