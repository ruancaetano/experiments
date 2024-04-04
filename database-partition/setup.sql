CREATE TABLE log_events (
    event_id SERIAL PRIMARY KEY,
    event_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    event_type VARCHAR(50) NOT NULL,
    event_application VARCHAR(50) NOT NULL,
    event_description TEXT
);

CREATE INDEX idx_event_type ON log_events (event_type);
CREATE INDEX idx_event_application ON log_events (event_application);


