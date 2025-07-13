package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"meme-api/internal/meme-api/models"
)

const (
	memeAPIURL = "https://meme-api.com/gimme"
)

type MemeService struct {
	client *http.Client
}

func NewMemeService() *MemeService {
	return &MemeService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *MemeService) GetRandomMeme() (*models.Meme, error) {
	resp, err := s.client.Get(memeAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch meme: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var meme models.Meme
	if err := json.NewDecoder(resp.Body).Decode(&meme); err != nil {
		return nil, fmt.Errorf("failed to decode meme response: %w", err)
	}

	return &meme, nil
}

func (s *MemeService) GetMemeBySubreddit(subreddit string) (*models.Meme, error) {
	url := fmt.Sprintf("%s/%s", memeAPIURL, subreddit)
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch meme from subreddit %s: %w", subreddit, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for subreddit %s: %d", subreddit, resp.StatusCode)
	}

	var meme models.Meme
	if err := json.NewDecoder(resp.Body).Decode(&meme); err != nil {
		return nil, fmt.Errorf("failed to decode meme response for subreddit %s: %w", subreddit, err)
	}

	return &meme, nil
}
