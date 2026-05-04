CREATE INDEX idx_aircraft_data_last_seen_distance
ON aircraft_data (last_seen DESC, last_seen_distance);
