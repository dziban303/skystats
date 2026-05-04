package main

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

type CachedStatsService struct {
	pg *postgres
}

func NewCachedStatsService(pg *postgres) *CachedStatsService {
	return &CachedStatsService{pg: pg}
}

func (s *CachedStatsService) GetCachedStat(key string) (int, error) {
	var value int
	err := s.pg.db.QueryRow(context.Background(),
		"SELECT stat_value FROM cached_stats WHERE stat_key = $1", key).Scan(&value)
	return value, err
}

func (s *CachedStatsService) UpdateCachedStat(key string, value int) error {
	_, err := s.pg.db.Exec(context.Background(),
		`INSERT INTO cached_stats (stat_key, stat_value, last_updated)
		VALUES ($1, $2, NOW())
		ON CONFLICT (stat_key)
		DO UPDATE SET stat_value = $2, last_updated = NOW()`,
		key, value)
	return err
}

func (s *CachedStatsService) RefreshAllTotals() error {
	log.Info().Msg("Refreshing cached all-time statistics...")

	// Update total flights
	var totalFlights int
	err := s.pg.db.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM aircraft_data").Scan(&totalFlights)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count total flights")
		return err
	}
	err = s.UpdateCachedStat("total_flights", totalFlights)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update cached total_flights")
		return err
	}

	// Update total aircraft
	var totalAircraft int
	err = s.pg.db.QueryRow(context.Background(),
		"SELECT COUNT(DISTINCT hex) FROM aircraft_data").Scan(&totalAircraft)
	if err != nil {
		log.Error().Err(err).Msg("Failed to count total aircraft")
		return err
	}
	err = s.UpdateCachedStat("total_aircraft", totalAircraft)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update cached total_aircraft")
		return err
	}

	log.Info().
		Int("total_flights", totalFlights).
		Int("total_aircraft", totalAircraft).
		Msg("Successfully refreshed cached statistics")

	return nil
}

func (s *CachedStatsService) StartPeriodicRefresh(interval time.Duration) {
	go func() {
		// Initial refresh on startup
		if err := s.RefreshAllTotals(); err != nil {
			log.Error().Err(err).Msg("Initial cache refresh failed")
		}

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			if err := s.RefreshAllTotals(); err != nil {
				log.Error().Err(err).Msg("Periodic cache refresh failed")
			}
		}
	}()
}
