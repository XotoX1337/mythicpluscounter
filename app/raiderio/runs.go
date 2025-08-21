package raiderio

import (
	"net/http"
	"time"
)

type RunsService service

type Run struct {
	Summary struct {
		Status  string `json:"status"`
		Dungeon struct {
			Id        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"short_name"`
			Slug      string `json:"slug"`
		} `json:"dungeon"`
		MythicLevel     int       `json:"mythic_level"`
		ClearTimeMs     int       `json:"clear_time_ms"`
		KeystoneTimeMs  int       `json:"keystone_time_ms"`
		TimeRemainingMs int       `json:"time_remaining_ms"`
		CompletedAt     time.Time `json:"completed_at"`
		NumChests       int       `json:"num_chests"`
		Role            string    `json:"role"`
	} `json:"summary"`
	Score float64 `json:"score"`
}

type Runs struct {
	Runs *[]Run `json:"runs"`
}

type RunsListOptions struct {
	Season      string `url:"season,omitempty"`
	CharacterId string `url:"characterId,omitempty"`
	ListOptions
}

func (s *RunsService) List(opts *RunsListOptions) (*[]Run, error) {
	var queryURL string
	var err error
	var runs *Runs

	queryURL, err = addOptions("characters/mythic-plus-runs", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, queryURL)
	if err != nil {
		return nil, err
	}

	err = s.client.Do(req, &runs)
	if err != nil {
		return nil, err
	}

	return runs.Runs, nil
}
