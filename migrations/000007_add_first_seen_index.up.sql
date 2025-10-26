-- add first_seen index to speed up various queries
CREATE INDEX IF NOT EXISTS idx_aircraft_data_first_seen ON aircraft_data (first_seen DESC);
