package riotclient

import (
	"encoding/json"
	"io"
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

func (c *Client) Summoner(name string) (Summoner, error) {
	summoner := Summoner{}
	err := c.makeFullReq("GET", "https://"+c.region.String()+".api.riotgames.com/lol/summoner/v4/summoners/by-name/"+name, nil, &summoner)
	if err != nil {
		return Summoner{}, nil
	}

	return summoner, nil
}

func (c *Client) MatchesBySummonerName(name string) ([]MatchDTO, error) {
	summoner, err := c.Summoner(name)
	if err != nil {
		return nil, err
	}

	matchIDs, err := c.matchesByPUUID(summoner.PUUID)
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

func (c *Client) ChampionMasteries(summonerID string) ([]ChampionMasteryDTO, error) {
	masteries := make([]ChampionMasteryDTO, DefaultResponseCount)
	err := c.makeFullReq("GET", "https://"+c.region.String()+".api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/"+summonerID, nil, &masteries)
	if err != nil {
		return nil, err
	}

	return masteries, nil
}

func (c *Client) matchesByPUUID(puuid string) (MatchesByPUUIDResponse, error) {
	matchIDs := make(MatchesByPUUIDResponse, DefaultResponseCount)
	err := c.makeFullReq("GET", "https://"+c.region.Routing()+".api.riotgames.com/lol/match/v5/matches/by-puuid/"+puuid+"/ids", nil, &matchIDs)
	if err != nil {
		return nil, err
	}

	return matchIDs, nil
}

func (c *Client) matchByID(matchID string) (MatchDTO, error) {
	match := MatchDTO{}
	err := c.makeFullReq("GET", "https://"+c.region.Routing()+".api.riotgames.com/lol/match/v5/matches/"+matchID, nil, &match)
	if err != nil {
		return MatchDTO{}, err
	}

	return match, nil
}

func (c *Client) makeFullReq(method string, url string, body io.Reader, out interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("X-Riot-Token", c.APIKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(bytes, &out)
	return nil
}
