package riotclient

type SummonerByNameResponse struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	SummonerLevel int    `json:"summonerLevel"`
}

type MatchesByPUUIDResponse []string

type MatchDTO struct {
	Metadata MetadataDTO `json:"metadata"`
	Info     InfoDTO     `json:"info"`
}

type MetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}

type InfoDTO struct {
	GameCreation       uint64        `json:"gameCreation"`
	GameDuration       uint64        `json:"gameDuration"`
	GameEndTimestamp   uint64        `json:"gameEndTimestamp"`
	GameID             uint64        `json:"gameId"`
	GameMode           string        `json:"gameMode"`
	GameName           string        `json:"gameName"`
	GameStartTimestamp uint64        `json:"gameStartTimestamp"`
	GameType           string        `json:"gameType"`
	GameVersion        string        `json:"gameVersion"`
	MapId              int           `json:"mapId"`
	Participants       []Participant `json:"participants"`
	PlatformID         string        `json:"platformId"`
	QueueID            int           `json:"queueId"`
	Teams              []Team        `json:"teams"`
	TournamentCode     string        `json:"tornamentCode"`
}

// type Participant struct {
// 	// TODO
// }

// Participant allows ANY json to be unmarshalled
// Participant should be replaced with a struct when the time comes to use it
type Participant map[string]interface{}

type Team struct {
	Bans       []BanDTO      `json:"bans"`
	Objectives ObjectivesDTO `json:"objectives"`
	TeamID     int           `json:"teamId"`
	Win        bool          `json:"win"`
}

type BanDTO struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}

type ObjectivesDTO struct {
	Baron      ObjectiveDTO `json:"baron"`
	Champion   ObjectiveDTO `json:"champion"`
	Dragon     ObjectiveDTO `json:"dragon"`
	Inhibitor  ObjectiveDTO `json:"inhibitor"`
	RiftHerald ObjectiveDTO `json:"riftHerald"`
	Tower      ObjectiveDTO `json:"tower"`
}

type ObjectiveDTO struct {
	First bool `json:"first"`
	Kills int  `json:"kills"`
}

type ChampionMasteryDTO struct {
	ChampionPointsUntilNextLevel uint64 `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionID                   uint64 `json:"championId"`
	LastPlayedTime               uint64 `json:"lastPlayedTime"`
	ChampionLevel                int    `json:"championLevel"`
	EncryptedSummonerId          string `json:"summonerId"`
	ChampionPoints               uint64 `json:"championPoints"`
	ChampionPointsSinceLastLevel uint64 `json:"championPointsSinceLastLevel"`
	TokensEarned                 uint64 `json:"tokensEarned"`
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	ProfileIconID string `json:"profileIconId"`
	RevisionDate  string `json:"revisionDate"`
	Name          string `json:"name"`
	PUUID         string `json:"puuid"`
	SummonerLevel string `json:"summonerLevel"`
}
