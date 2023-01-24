package riotclient

import (
	"encoding/json"
	"fmt"
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

func (c *Client) MatchesForSummoner(name string) error {
	// TODO

	return nil
}

func (c *Client) MatchesByPUUID(puuid string) (MatchesByPUUIDResponse, error) {
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
	fmt.Println("found string:", string(bytes))
	fmt.Println("found matches:", matchIDs)
	return matchIDs, nil
}
