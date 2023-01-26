package riotclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	APIKey string
	region Region
	client *http.Client
}

func New(key string, r Region) *Client {
	return &Client{
		APIKey: key,
		region: r,
		client: http.DefaultClient,
	}
}

func (c *Client) SummonerPUUID(name string) (string, error) {
	req, _ := http.NewRequest("GET", "https://"+c.region.String()+".api.riotgames.com/lol/summoner/v4/summoners/by-name/"+name, nil)
	req.Header.Set("X-Riot-Token", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	summoner := SummonerByNameResponse{}
	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	json.Unmarshal(bytes, &summoner)
	return summoner.PUUID, nil
}

func (c *Client) MatchesBySummonerName(name string) ([]MatchDTO, error) {
	puuid, err := c.SummonerPUUID(name)
	if err != nil {
		return nil, err
	}

	matchIDs, err := c.matchesByPUUID(puuid)
	if err != nil {
		return nil, err
	}

	var matches []MatchDTO
	for _, matchID := range matchIDs {
		match, err := c.matchByID(matchID)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}

	return matches, nil
}

func (c *Client) matchesByPUUID(puuid string) (MatchesByPUUIDResponse, error) {
	req, _ := http.NewRequest("GET", "https://"+c.region.Routing()+".api.riotgames.com/lol/match/v5/matches/by-puuid/"+puuid+"/ids", nil)
	req.Header.Set("X-Riot-Token", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matchIDs := make(MatchesByPUUIDResponse, DefaultResponseCount)
	json.Unmarshal(bytes, &matchIDs)
	return matchIDs, nil
}

func (c *Client) matchByID(matchID string) (MatchDTO, error) {
	req, _ := http.NewRequest("GET", "https://"+c.region.Routing()+".api.riotgames.com/lol/match/v5/matches/"+matchID, nil)
	req.Header.Set("X-Riot-Token", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return MatchDTO{}, err
	}

	defer resp.Body.Close()

	match := MatchDTO{}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return MatchDTO{}, err
	}

	json.Unmarshal(bytes, &match)
	return match, nil
}
