package lol

type Summoner struct {
	AccountId                   int64  `json:"accountId"`
	DisplayName                 string `json:"displayName"`
	GameName                    string `json:"gameName"`
	InternalName                string `json:"internalName"`
	NameChangeFlag              bool   `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int    `json:"percentCompleteForNextLevel"`
	Privacy                     string `json:"privacy"`
	ProfileIconId               int    `json:"profileIconId"`
	Puuid                       string `json:"puuid"`
	RerollPoints                struct {
		CurrentPoints    int `json:"currentPoints"`
		MaxRolls         int `json:"maxRolls"`
		NumberOfRolls    int `json:"numberOfRolls"`
		PointsCostToRoll int `json:"pointsCostToRoll"`
		PointsToReroll   int `json:"pointsToReroll"`
	} `json:"rerollPoints"`
	SummonerId       int64  `json:"summonerId"`
	SummonerLevel    int    `json:"summonerLevel"`
	TagLine          string `json:"tagLine"`
	Unnamed          bool   `json:"unnamed"`
	XpSinceLastLevel int    `json:"xpSinceLastLevel"`
	XpUntilNextLevel int    `json:"xpUntilNextLevel"`
}

type SummonerViaSgp struct {
	Id                int64  `json:"id"`
	Puuid             string `json:"puuid"`
	AccountId         int64  `json:"accountId"`
	Name              string `json:"name"`
	InternalName      string `json:"internalName"`
	ProfileIconId     int    `json:"profileIconId"`
	Level             int    `json:"level"`
	ExpPoints         int    `json:"expPoints"`
	LevelAndXpVersion int    `json:"levelAndXpVersion"`
	RevisionId        int    `json:"revisionId"`
	RevisionDate      int64  `json:"revisionDate"`
	LastGameDate      int64  `json:"lastGameDate"`
	NameChangeFlag    bool   `json:"nameChangeFlag"`
	Unnamed           bool   `json:"unnamed"`
	Privacy           string `json:"privacy"`
	ExpToNextLevel    int    `json:"expToNextLevel"`
}

type GamingInfo struct {
	ReconnectDelay int    `json:"reconnectDelay"`
	GameName       string `json:"gameName"`
	Game           struct {
		Id                int64  `json:"id"`
		GameState         string `json:"gameState"`
		QueueTypeName     string `json:"queueTypeName"`
		Name              string `json:"name"`
		PickTurn          int    `json:"pickTurn"`
		MapId             int    `json:"mapId"`
		GameMode          string `json:"gameMode"`
		MaxNumPlayers     int    `json:"maxNumPlayers"`
		GameType          string `json:"gameType"`
		GameQueueConfigId int    `json:"gameQueueConfigId"`
		SpectatorDelay    int    `json:"spectatorDelay"`
		GameVersion       string `json:"gameVersion"`
		TeamOne           []struct {
			Puuid                 string `json:"puuid"`
			SummonerId            int64  `json:"summonerId"`
			LastSelectedSkinIndex int    `json:"lastSelectedSkinIndex"`
			TeamOwner             bool   `json:"teamOwner"`
			ProfileIconId         int    `json:"profileIconId"`
			TeamParticipantId     int    `json:"teamParticipantId"`
			ChampionId            int    `json:"championId"`
			SelectedRole          string `json:"selectedRole"`
			SelectedPosition      string `json:"selectedPosition"`
			SummonerName          string `json:"summonerName"`
			SummonerInternalName  string `json:"summonerInternalName"`
		} `json:"teamOne"`
		TeamTwo []struct {
			Puuid                 string `json:"puuid"`
			SummonerId            int64  `json:"summonerId"`
			LastSelectedSkinIndex int    `json:"lastSelectedSkinIndex"`
			TeamOwner             bool   `json:"teamOwner"`
			ProfileIconId         int    `json:"profileIconId"`
			TeamParticipantId     int    `json:"teamParticipantId"`
			ChampionId            int    `json:"championId"`
			SelectedRole          string `json:"selectedRole"`
			SelectedPosition      string `json:"selectedPosition"`
			SummonerName          string `json:"summonerName"`
			SummonerInternalName  string `json:"summonerInternalName"`
		} `json:"teamTwo"`
		PlayerChampionSelections []struct {
			SummonerInternalName string `json:"summonerInternalName"`
			ChampionId           int    `json:"championId"`
			SelectedSkinIndex    int    `json:"selectedSkinIndex"`
			Spell1Id             int    `json:"spell1Id"`
			Spell2Id             int    `json:"spell2Id"`
		} `json:"playerChampionSelections"`
		BannedChampions []interface{} `json:"bannedChampions"`
		Observers       []interface{} `json:"observers"`
	} `json:"game"`
	PlayerCredentials struct {
		GameId                int64  `json:"gameId"`
		QueueId               int    `json:"queueId"`
		PlayerId              int64  `json:"playerId"`
		Puuid                 string `json:"puuid"`
		ServerPort            int    `json:"serverPort"`
		ChampionId            int    `json:"championId"`
		LastSelectedSkinIndex int    `json:"lastSelectedSkinIndex"`
		SummonerId            int64  `json:"summonerId"`
		Observer              bool   `json:"observer"`
		GameVersion           string `json:"gameVersion"`
		GameMode              string `json:"gameMode"`
		ObserverEncryptionKey string `json:"observerEncryptionKey"`
		ObserverServerIp      string `json:"observerServerIp"`
		ObserverServerPort    int    `json:"observerServerPort"`
		QueueType             string `json:"queueType"`
		GameCreateDate        int64  `json:"gameCreateDate"`
	} `json:"playerCredentials"`
}
