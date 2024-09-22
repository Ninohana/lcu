package lcu

import "time"

// SgpToken 鉴权信息
type SgpToken struct {
	AccessToken string `json:"accessToken"`
	Issuer      string `json:"issuer"`
	Subject     string `json:"subject"`
	Token       string `json:"token"`
}

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

type ReplaysConfigurationV1 struct {
	GameVersion                      string `json:"gameVersion"`
	IsInTournament                   bool   `json:"isInTournament"`
	IsLoggedIn                       bool   `json:"isLoggedIn"`
	IsPatching                       bool   `json:"isPatching"`
	IsPlayingGame                    bool   `json:"isPlayingGame"`
	IsPlayingReplay                  bool   `json:"isPlayingReplay"`
	IsReplaysEnabled                 bool   `json:"isReplaysEnabled"`
	IsReplaysForEndOfGameEnabled     bool   `json:"isReplaysForEndOfGameEnabled"`
	IsReplaysForMatchHistoryEnabled  bool   `json:"isReplaysForMatchHistoryEnabled"`
	MinServerVersion                 string `json:"minServerVersion"`
	MinutesUntilReplayConsideredLost int    `json:"minutesUntilReplayConsideredLost"`
}

type GameInfo struct {
	EndOfGameResult       string    `json:"endOfGameResult"`
	GameCreation          int64     `json:"gameCreation"`
	GameCreationDate      time.Time `json:"gameCreationDate"`
	GameDuration          int       `json:"gameDuration"`
	GameId                int64     `json:"gameId"`
	GameMode              string    `json:"gameMode"`
	GameType              string    `json:"gameType"`
	GameVersion           string    `json:"gameVersion"`
	MapId                 int       `json:"mapId"`
	ParticipantIdentities []struct {
		ParticipantId int `json:"participantId"`
		Player        struct {
			AccountId         int    `json:"accountId"`
			CurrentAccountId  int    `json:"currentAccountId"`
			CurrentPlatformId string `json:"currentPlatformId"`
			GameName          string `json:"gameName"`
			MatchHistoryUri   string `json:"matchHistoryUri"`
			PlatformId        string `json:"platformId"`
			ProfileIcon       int    `json:"profileIcon"`
			Puuid             string `json:"puuid"`
			SummonerId        int64  `json:"summonerId"`
			SummonerName      string `json:"summonerName"`
			TagLine           string `json:"tagLine"`
		} `json:"player"`
	} `json:"participantIdentities"`
	Participants []struct {
		ChampionId                int    `json:"championId"`
		HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
		ParticipantId             int    `json:"participantId"`
		Spell1Id                  int    `json:"spell1Id"`
		Spell2Id                  int    `json:"spell2Id"`
		Stats                     struct {
			Assists                         int  `json:"assists"`
			CausedEarlySurrender            bool `json:"causedEarlySurrender"`
			ChampLevel                      int  `json:"champLevel"`
			CombatPlayerScore               int  `json:"combatPlayerScore"`
			DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
			DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
			DamageSelfMitigated             int  `json:"damageSelfMitigated"`
			Deaths                          int  `json:"deaths"`
			DoubleKills                     int  `json:"doubleKills"`
			EarlySurrenderAccomplice        bool `json:"earlySurrenderAccomplice"`
			FirstBloodAssist                bool `json:"firstBloodAssist"`
			FirstBloodKill                  bool `json:"firstBloodKill"`
			FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
			FirstInhibitorKill              bool `json:"firstInhibitorKill"`
			FirstTowerAssist                bool `json:"firstTowerAssist"`
			FirstTowerKill                  bool `json:"firstTowerKill"`
			GameEndedInEarlySurrender       bool `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender            bool `json:"gameEndedInSurrender"`
			GoldEarned                      int  `json:"goldEarned"`
			GoldSpent                       int  `json:"goldSpent"`
			InhibitorKills                  int  `json:"inhibitorKills"`
			Item0                           int  `json:"item0"`
			Item1                           int  `json:"item1"`
			Item2                           int  `json:"item2"`
			Item3                           int  `json:"item3"`
			Item4                           int  `json:"item4"`
			Item5                           int  `json:"item5"`
			Item6                           int  `json:"item6"`
			KillingSprees                   int  `json:"killingSprees"`
			Kills                           int  `json:"kills"`
			LargestCriticalStrike           int  `json:"largestCriticalStrike"`
			LargestKillingSpree             int  `json:"largestKillingSpree"`
			LargestMultiKill                int  `json:"largestMultiKill"`
			LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
			MagicDamageDealt                int  `json:"magicDamageDealt"`
			MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
			MagicalDamageTaken              int  `json:"magicalDamageTaken"`
			NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
			NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
			NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
			ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
			ParticipantId                   int  `json:"participantId"`
			PentaKills                      int  `json:"pentaKills"`
			Perk0                           int  `json:"perk0"`
			Perk0Var1                       int  `json:"perk0Var1"`
			Perk0Var2                       int  `json:"perk0Var2"`
			Perk0Var3                       int  `json:"perk0Var3"`
			Perk1                           int  `json:"perk1"`
			Perk1Var1                       int  `json:"perk1Var1"`
			Perk1Var2                       int  `json:"perk1Var2"`
			Perk1Var3                       int  `json:"perk1Var3"`
			Perk2                           int  `json:"perk2"`
			Perk2Var1                       int  `json:"perk2Var1"`
			Perk2Var2                       int  `json:"perk2Var2"`
			Perk2Var3                       int  `json:"perk2Var3"`
			Perk3                           int  `json:"perk3"`
			Perk3Var1                       int  `json:"perk3Var1"`
			Perk3Var2                       int  `json:"perk3Var2"`
			Perk3Var3                       int  `json:"perk3Var3"`
			Perk4                           int  `json:"perk4"`
			Perk4Var1                       int  `json:"perk4Var1"`
			Perk4Var2                       int  `json:"perk4Var2"`
			Perk4Var3                       int  `json:"perk4Var3"`
			Perk5                           int  `json:"perk5"`
			Perk5Var1                       int  `json:"perk5Var1"`
			Perk5Var2                       int  `json:"perk5Var2"`
			Perk5Var3                       int  `json:"perk5Var3"`
			PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
			PerkSubStyle                    int  `json:"perkSubStyle"`
			PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
			PlayerAugment1                  int  `json:"playerAugment1"`
			PlayerAugment2                  int  `json:"playerAugment2"`
			PlayerAugment3                  int  `json:"playerAugment3"`
			PlayerAugment4                  int  `json:"playerAugment4"`
			PlayerAugment5                  int  `json:"playerAugment5"`
			PlayerAugment6                  int  `json:"playerAugment6"`
			PlayerScore0                    int  `json:"playerScore0"`
			PlayerScore1                    int  `json:"playerScore1"`
			PlayerScore2                    int  `json:"playerScore2"`
			PlayerScore3                    int  `json:"playerScore3"`
			PlayerScore4                    int  `json:"playerScore4"`
			PlayerScore5                    int  `json:"playerScore5"`
			PlayerScore6                    int  `json:"playerScore6"`
			PlayerScore7                    int  `json:"playerScore7"`
			PlayerScore8                    int  `json:"playerScore8"`
			PlayerScore9                    int  `json:"playerScore9"`
			PlayerSubteamId                 int  `json:"playerSubteamId"`
			QuadraKills                     int  `json:"quadraKills"`
			SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
			SubteamPlacement                int  `json:"subteamPlacement"`
			TeamEarlySurrendered            bool `json:"teamEarlySurrendered"`
			TimeCCingOthers                 int  `json:"timeCCingOthers"`
			TotalDamageDealt                int  `json:"totalDamageDealt"`
			TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
			TotalDamageTaken                int  `json:"totalDamageTaken"`
			TotalHeal                       int  `json:"totalHeal"`
			TotalMinionsKilled              int  `json:"totalMinionsKilled"`
			TotalPlayerScore                int  `json:"totalPlayerScore"`
			TotalScoreRank                  int  `json:"totalScoreRank"`
			TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
			TotalUnitsHealed                int  `json:"totalUnitsHealed"`
			TripleKills                     int  `json:"tripleKills"`
			TrueDamageDealt                 int  `json:"trueDamageDealt"`
			TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                 int  `json:"trueDamageTaken"`
			TurretKills                     int  `json:"turretKills"`
			UnrealKills                     int  `json:"unrealKills"`
			VisionScore                     int  `json:"visionScore"`
			VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
			WardsKilled                     int  `json:"wardsKilled"`
			WardsPlaced                     int  `json:"wardsPlaced"`
			Win                             bool `json:"win"`
		} `json:"stats"`
		TeamId   int `json:"teamId"`
		Timeline struct {
			CreepsPerMinDeltas struct {
			} `json:"creepsPerMinDeltas"`
			CsDiffPerMinDeltas struct {
			} `json:"csDiffPerMinDeltas"`
			DamageTakenDiffPerMinDeltas struct {
			} `json:"damageTakenDiffPerMinDeltas"`
			DamageTakenPerMinDeltas struct {
			} `json:"damageTakenPerMinDeltas"`
			GoldPerMinDeltas struct {
			} `json:"goldPerMinDeltas"`
			Lane               string `json:"lane"`
			ParticipantId      int    `json:"participantId"`
			Role               string `json:"role"`
			XpDiffPerMinDeltas struct {
			} `json:"xpDiffPerMinDeltas"`
			XpPerMinDeltas struct {
			} `json:"xpPerMinDeltas"`
		} `json:"timeline"`
	} `json:"participants"`
	PlatformId string `json:"platformId"`
	QueueId    int    `json:"queueId"`
	SeasonId   int    `json:"seasonId"`
	Teams      []struct {
		Bans []struct {
			ChampionId int `json:"championId"`
			PickTurn   int `json:"pickTurn"`
		} `json:"bans"`
		BaronKills           int    `json:"baronKills"`
		DominionVictoryScore int    `json:"dominionVictoryScore"`
		DragonKills          int    `json:"dragonKills"`
		FirstBaron           bool   `json:"firstBaron"`
		FirstBlood           bool   `json:"firstBlood"`
		FirstDargon          bool   `json:"firstDargon"`
		FirstInhibitor       bool   `json:"firstInhibitor"`
		FirstTower           bool   `json:"firstTower"`
		HordeKills           int    `json:"hordeKills"`
		InhibitorKills       int    `json:"inhibitorKills"`
		RiftHeraldKills      int    `json:"riftHeraldKills"`
		TeamId               int    `json:"teamId"`
		TowerKills           int    `json:"towerKills"`
		VilemawKills         int    `json:"vilemawKills"`
		Win                  string `json:"win"`
	} `json:"teams"`
}

type GamesInfo struct {
	AccountId int64 `json:"accountId"`
	Games     struct {
		GameBeginDate  string     `json:"gameBeginDate"`
		GameCount      int        `json:"gameCount"`
		GameEndDate    string     `json:"gameEndDate"`
		GameIndexBegin int        `json:"gameIndexBegin"`
		GameIndexEnd   int        `json:"gameIndexEnd"`
		Games          []GameInfo `json:"games"`
	} `json:"games"`
	PlatformId string `json:"platformId"`
}

type SummonerProfile struct {
	BackgroundSkinId int    `json:"backgroundSkinId"`
	Regalia          string `json:"regalia"`
}

type GameClient struct {
	ObserverServerIp   string `json:"observerServerIp"`
	ObserverServerPort int    `json:"observerServerPort"`
	Running            bool   `json:"running"`
	ServerIp           string `json:"serverIp"`
	ServerPort         int    `json:"serverPort"`
	Visible            bool   `json:"visible"`
}

type GameflowInfo struct {
	GameClient `json:"gameClient"`
	GameData   struct {
		GameId                   int           `json:"gameId"`
		GameName                 string        `json:"gameName"`
		IsCustomGame             bool          `json:"isCustomGame"`
		Password                 string        `json:"password"`
		PlayerChampionSelections []interface{} `json:"playerChampionSelections"`
		Queue                    struct {
			AllowablePremadeSizes   []interface{} `json:"allowablePremadeSizes"`
			AreFreeChampionsAllowed bool          `json:"areFreeChampionsAllowed"`
			AssetMutator            string        `json:"assetMutator"`
			Category                string        `json:"category"`
			ChampionsRequiredToPlay int           `json:"championsRequiredToPlay"`
			Description             string        `json:"description"`
			DetailedDescription     string        `json:"detailedDescription"`
			GameMode                string        `json:"gameMode"`
			GameTypeConfig          struct {
				AdvancedLearningQuests bool   `json:"advancedLearningQuests"`
				AllowTrades            bool   `json:"allowTrades"`
				BanMode                string `json:"banMode"`
				BanTimerDuration       int    `json:"banTimerDuration"`
				BattleBoost            bool   `json:"battleBoost"`
				CrossTeamChampionPool  bool   `json:"crossTeamChampionPool"`
				DeathMatch             bool   `json:"deathMatch"`
				DoNotRemove            bool   `json:"doNotRemove"`
				DuplicatePick          bool   `json:"duplicatePick"`
				ExclusivePick          bool   `json:"exclusivePick"`
				Id                     int    `json:"id"`
				LearningQuests         bool   `json:"learningQuests"`
				MainPickTimerDuration  int    `json:"mainPickTimerDuration"`
				MaxAllowableBans       int    `json:"maxAllowableBans"`
				Name                   string `json:"name"`
				OnboardCoopBeginner    bool   `json:"onboardCoopBeginner"`
				PickMode               string `json:"pickMode"`
				PostPickTimerDuration  int    `json:"postPickTimerDuration"`
				Reroll                 bool   `json:"reroll"`
				TeamChampionPool       bool   `json:"teamChampionPool"`
			} `json:"gameTypeConfig"`
			Id                         int    `json:"id"`
			IsRanked                   bool   `json:"isRanked"`
			IsTeamBuilderManaged       bool   `json:"isTeamBuilderManaged"`
			LastToggledOffTime         int    `json:"lastToggledOffTime"`
			LastToggledOnTime          int    `json:"lastToggledOnTime"`
			MapId                      int    `json:"mapId"`
			MaximumParticipantListSize int    `json:"maximumParticipantListSize"`
			MinLevel                   int    `json:"minLevel"`
			MinimumParticipantListSize int    `json:"minimumParticipantListSize"`
			Name                       string `json:"name"`
			NumPlayersPerTeam          int    `json:"numPlayersPerTeam"`
			QueueAvailability          string `json:"queueAvailability"`
			QueueRewards               struct {
				IsChampionPointsEnabled bool          `json:"isChampionPointsEnabled"`
				IsIpEnabled             bool          `json:"isIpEnabled"`
				IsXpEnabled             bool          `json:"isXpEnabled"`
				PartySizeIpRewards      []interface{} `json:"partySizeIpRewards"`
			} `json:"queueRewards"`
			RemovalFromGameAllowed      bool   `json:"removalFromGameAllowed"`
			RemovalFromGameDelayMinutes int    `json:"removalFromGameDelayMinutes"`
			ShortName                   string `json:"shortName"`
			ShowPositionSelector        bool   `json:"showPositionSelector"`
			SpectatorEnabled            bool   `json:"spectatorEnabled"`
			Type                        string `json:"type"`
		} `json:"queue"`
		SpectatorsAllowed bool          `json:"spectatorsAllowed"`
		TeamOne           []interface{} `json:"teamOne"`
		TeamTwo           []interface{} `json:"teamTwo"`
	} `json:"gameData"`
	GameDodge struct {
		DodgeIds []interface{} `json:"dodgeIds"`
		Phase    string        `json:"phase"`
		State    string        `json:"state"`
	} `json:"gameDodge"`
	Map struct {
		Assets struct {
			ChampSelectBackgroundSound  string `json:"champ-select-background-sound"`
			ChampSelectFlyoutBackground string `json:"champ-select-flyout-background"`
			ChampSelectPlanningIntro    string `json:"champ-select-planning-intro"`
			GameSelectIconActive        string `json:"game-select-icon-active"`
			GameSelectIconActiveVideo   string `json:"game-select-icon-active-video"`
			GameSelectIconDefault       string `json:"game-select-icon-default"`
			GameSelectIconDisabled      string `json:"game-select-icon-disabled"`
			GameSelectIconHover         string `json:"game-select-icon-hover"`
			GameSelectIconIntroVideo    string `json:"game-select-icon-intro-video"`
			GameflowBackground          string `json:"gameflow-background"`
			GameflowBackgroundDark      string `json:"gameflow-background-dark"`
			GameselectButtonHoverSound  string `json:"gameselect-button-hover-sound"`
			IconDefeat                  string `json:"icon-defeat"`
			IconDefeatV2                string `json:"icon-defeat-v2"`
			IconDefeatVideo             string `json:"icon-defeat-video"`
			IconEmpty                   string `json:"icon-empty"`
			IconHover                   string `json:"icon-hover"`
			IconLeaver                  string `json:"icon-leaver"`
			IconLeaverV2                string `json:"icon-leaver-v2"`
			IconLossForgivenV2          string `json:"icon-loss-forgiven-v2"`
			IconV2                      string `json:"icon-v2"`
			IconVictory                 string `json:"icon-victory"`
			IconVictoryVideo            string `json:"icon-victory-video"`
			MapNorth                    string `json:"map-north"`
			MapSouth                    string `json:"map-south"`
			MusicInqueueLoopSound       string `json:"music-inqueue-loop-sound"`
			PartiesBackground           string `json:"parties-background"`
			PostgameAmbienceLoopSound   string `json:"postgame-ambience-loop-sound"`
			ReadyCheckBackground        string `json:"ready-check-background"`
			ReadyCheckBackgroundSound   string `json:"ready-check-background-sound"`
			SfxAmbiencePregameLoopSound string `json:"sfx-ambience-pregame-loop-sound"`
			SocialIconLeaver            string `json:"social-icon-leaver"`
			SocialIconVictory           string `json:"social-icon-victory"`
		} `json:"assets"`
		CategorizedContentBundles struct {
		} `json:"categorizedContentBundles"`
		Description                         string `json:"description"`
		GameMode                            string `json:"gameMode"`
		GameModeName                        string `json:"gameModeName"`
		GameModeShortName                   string `json:"gameModeShortName"`
		GameMutator                         string `json:"gameMutator"`
		Id                                  int    `json:"id"`
		IsRGM                               bool   `json:"isRGM"`
		MapStringId                         string `json:"mapStringId"`
		Name                                string `json:"name"`
		PerPositionDisallowedSummonerSpells struct {
		} `json:"perPositionDisallowedSummonerSpells"`
		PerPositionRequiredSummonerSpells struct {
		} `json:"perPositionRequiredSummonerSpells"`
		PlatformId   string `json:"platformId"`
		PlatformName string `json:"platformName"`
		Properties   struct {
			SuppressRunesMasteriesPerks bool `json:"suppressRunesMasteriesPerks"`
		} `json:"properties"`
	} `json:"map"`
	Phase string `json:"phase"`
}

type SelectSession struct {
	Actions [][]struct {
		ActorCellId  int    `json:"actorCellId"`
		ChampionId   int    `json:"championId"`
		Completed    bool   `json:"completed"`
		Id           int    `json:"id"`
		IsAllyAction bool   `json:"isAllyAction"`
		IsInProgress bool   `json:"isInProgress"`
		PickTurn     int    `json:"pickTurn"`
		Type         string `json:"type"`
	} `json:"actions"`
	AllowBattleBoost    bool `json:"allowBattleBoost"`
	AllowDuplicatePicks bool `json:"allowDuplicatePicks"`
	AllowLockedEvents   bool `json:"allowLockedEvents"`
	AllowRerolling      bool `json:"allowRerolling"`
	AllowSkinSelection  bool `json:"allowSkinSelection"`
	Bans                struct {
		MyTeamBans    []interface{} `json:"myTeamBans"`
		NumBans       int           `json:"numBans"`
		TheirTeamBans []interface{} `json:"theirTeamBans"`
	} `json:"bans"`
	BenchChampions     []interface{} `json:"benchChampions"`
	BenchEnabled       bool          `json:"benchEnabled"`
	BoostableSkinCount int           `json:"boostableSkinCount"`
	ChatDetails        struct {
		MucJwtDto struct {
			ChannelClaim string `json:"channelClaim"`
			Domain       string `json:"domain"`
			Jwt          string `json:"jwt"`
			TargetRegion string `json:"targetRegion"`
		} `json:"mucJwtDto"`
		MultiUserChatId       string `json:"multiUserChatId"`
		MultiUserChatPassword string `json:"multiUserChatPassword"`
	} `json:"chatDetails"`
	Counter              int  `json:"counter"`
	GameId               int  `json:"gameId"`
	HasSimultaneousBans  bool `json:"hasSimultaneousBans"`
	HasSimultaneousPicks bool `json:"hasSimultaneousPicks"`
	IsCustomGame         bool `json:"isCustomGame"`
	IsSpectating         bool `json:"isSpectating"`
	LocalPlayerCellId    int  `json:"localPlayerCellId"`
	LockedEventIndex     int  `json:"lockedEventIndex"`
	MyTeam               []struct {
		AssignedPosition     string `json:"assignedPosition"`
		CellId               int    `json:"cellId"`
		ChampionId           int    `json:"championId"`
		ChampionPickIntent   int    `json:"championPickIntent"`
		NameVisibilityType   string `json:"nameVisibilityType"`
		ObfuscatedPuuid      string `json:"obfuscatedPuuid"`
		ObfuscatedSummonerId int    `json:"obfuscatedSummonerId"`
		Puuid                string `json:"puuid"`
		SelectedSkinId       int    `json:"selectedSkinId"`
		Spell1Id             int    `json:"spell1Id"`
		Spell2Id             int    `json:"spell2Id"`
		SummonerId           int64  `json:"summonerId"`
		Team                 int    `json:"team"`
		WardSkinId           int    `json:"wardSkinId"`
	} `json:"myTeam"`
	PickOrderSwaps     []interface{} `json:"pickOrderSwaps"`
	RecoveryCounter    int           `json:"recoveryCounter"`
	RerollsRemaining   int           `json:"rerollsRemaining"`
	SkipChampionSelect bool          `json:"skipChampionSelect"`
	TheirTeam          []interface{} `json:"theirTeam"`
	Timer              struct {
		AdjustedTimeLeftInPhase int    `json:"adjustedTimeLeftInPhase"`
		InternalNowInEpochMs    int64  `json:"internalNowInEpochMs"`
		IsInfinite              bool   `json:"isInfinite"`
		Phase                   string `json:"phase"`
		TotalTimeInPhase        int    `json:"totalTimeInPhase"`
	} `json:"timer"`
	Trades []interface{} `json:"trades"`
}
