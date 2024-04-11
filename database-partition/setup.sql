CREATE or replace FUNCTION extract_month(ts timestamp) RETURNS integer AS $$
SELECT  EXTRACT(MONTH FROM $1);
$$ LANGUAGE SQL IMMUTABLE;


CREATE TABLE log_events (
                            event_id SERIAL,
                            event_time timestamp NOT NULL,
                            event_type VARCHAR(50) NOT NULL,
                            event_application VARCHAR(50) NOT NULL,
                            event_description TEXT
) PARTITION BY list(extract_month(event_time));

-- create partitions
CREATE TABLE log_events_month_1 PARTITION OF log_events
    FOR VALUES IN (1,4,7,10);


CREATE TABLE log_events_month_2 PARTITION OF log_events
    FOR VALUES IN (2,5,8,11);

CREATE TABLE log_events_month_3 PARTITION OF log_events
    FOR VALUES IN (3,6,9,12);

CREATE INDEX idx_event_type ON log_events (event_type);
CREATE INDEX idx_event_application ON log_events (event_application);




