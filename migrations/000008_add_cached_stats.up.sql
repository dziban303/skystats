-- Create table for cached statistics
CREATE TABLE cached_stats (
    stat_key VARCHAR(50) PRIMARY KEY,
    stat_value BIGINT NOT NULL,
    last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Insert initial values for all-time totals
INSERT INTO cached_stats (stat_key, stat_value, last_updated)
VALUES
    ('total_aircraft', (SELECT COUNT(DISTINCT hex) FROM aircraft_data), NOW());
