package RiotApi

import (
	"fmt"
	"strings"
)

type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func (a *Api) getSummonerByName(name string) (*Summoner, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", a.Server, name)

	s, err := get[Summoner](a, url)
	return s, err
}

type CurrentGameInfo struct {
	GameId            int64  `json:"gameId"`
	MapId             int    `json:"mapId"`
	GameMode          string `json:"gameMode"`
	GameType          string `json:"gameType"`
	GameQueueConfigId int    `json:"gameQueueConfigId"`
	Participants      []struct {
		TeamId                   int           `json:"teamId"`
		Spell1Id                 int           `json:"spell1Id"`
		Spell2Id                 int           `json:"spell2Id"`
		ChampionId               int           `json:"championId"`
		ProfileIconId            int           `json:"profileIconId"`
		SummonerName             string        `json:"summonerName"`
		Bot                      bool          `json:"bot"`
		SummonerId               string        `json:"summonerId"`
		GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
	} `json:"participants"`
	PlatformId    string `json:"platformId"`
	GameStartTime int64  `json:"gameStartTime"`
	GameLength    int    `json:"gameLength"`
}

func (a *Api) getActiveGameById(encryptedSummonerId string) (*CurrentGameInfo, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/%s", a.Server, encryptedSummonerId)

	return get[CurrentGameInfo](a, url)
}

func (a *Api) getMatchById(matchId string) (*MatchDto, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s_%d", strings.ToLower(a.region), a.Server, matchId)

	return get[MatchDto](a, url)
}

type MatchDto struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchId      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		GameCreation       int64  `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int64  `json:"gameEndTimestamp"`
		GameId             int64  `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int64  `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapId              int    `json:"mapId"`
		Participants       []struct {
			AllInPings    int `json:"allInPings"`
			AssistMePings int `json:"assistMePings"`
			Assists       int `json:"assists"`
			BaitPings     int `json:"baitPings"`
			BaronKills    int `json:"baronKills"`
			BasicPings    int `json:"basicPings"`
			BountyLevel   int `json:"bountyLevel"`
			Challenges    struct {
				AssistStreakCount                         int     `json:"12AssistStreakCount"`
				AbilityUses                               int     `json:"abilityUses"`
				AcesBefore15Minutes                       int     `json:"acesBefore15Minutes"`
				AlliedJungleMonsterKills                  int     `json:"alliedJungleMonsterKills"`
				BaronTakedowns                            int     `json:"baronTakedowns"`
				BlastConeOppositeOpponentCount            int     `json:"blastConeOppositeOpponentCount"`
				BountyGold                                int     `json:"bountyGold"`
				BuffsStolen                               int     `json:"buffsStolen"`
				CompleteSupportQuestInTime                int     `json:"completeSupportQuestInTime"`
				ControlWardsPlaced                        int     `json:"controlWardsPlaced"`
				DamagePerMinute                           float64 `json:"damagePerMinute"`
				DamageTakenOnTeamPercentage               float64 `json:"damageTakenOnTeamPercentage"`
				DancedWithRiftHerald                      int     `json:"dancedWithRiftHerald"`
				DeathsByEnemyChamps                       int     `json:"deathsByEnemyChamps"`
				DodgeSkillShotsSmallWindow                int     `json:"dodgeSkillShotsSmallWindow"`
				DoubleAces                                int     `json:"doubleAces"`
				DragonTakedowns                           int     `json:"dragonTakedowns"`
				EarlyLaningPhaseGoldExpAdvantage          int     `json:"earlyLaningPhaseGoldExpAdvantage"`
				EffectiveHealAndShielding                 float64 `json:"effectiveHealAndShielding"`
				ElderDragonKillsWithOpposingSoul          int     `json:"elderDragonKillsWithOpposingSoul"`
				ElderDragonMultikills                     int     `json:"elderDragonMultikills"`
				EnemyChampionImmobilizations              int     `json:"enemyChampionImmobilizations"`
				EnemyJungleMonsterKills                   int     `json:"enemyJungleMonsterKills"`
				EpicMonsterKillsNearEnemyJungler          int     `json:"epicMonsterKillsNearEnemyJungler"`
				EpicMonsterKillsWithin30SecondsOfSpawn    int     `json:"epicMonsterKillsWithin30SecondsOfSpawn"`
				EpicMonsterSteals                         int     `json:"epicMonsterSteals"`
				EpicMonsterStolenWithoutSmite             int     `json:"epicMonsterStolenWithoutSmite"`
				FirstTurretKilled                         int     `json:"firstTurretKilled"`
				FlawlessAces                              int     `json:"flawlessAces"`
				FullTeamTakedown                          int     `json:"fullTeamTakedown"`
				GameLength                                float64 `json:"gameLength"`
				GetTakedownsInAllLanesEarlyJungleAsLaner  int     `json:"getTakedownsInAllLanesEarlyJungleAsLaner,omitempty"`
				GoldPerMinute                             float64 `json:"goldPerMinute"`
				HadOpenNexus                              int     `json:"hadOpenNexus"`
				ImmobilizeAndKillWithAlly                 int     `json:"immobilizeAndKillWithAlly"`
				InitialBuffCount                          int     `json:"initialBuffCount"`
				InitialCrabCount                          int     `json:"initialCrabCount"`
				JungleCsBefore10Minutes                   float64 `json:"jungleCsBefore10Minutes"`
				JunglerTakedownsNearDamagedEpicMonster    int     `json:"junglerTakedownsNearDamagedEpicMonster"`
				KTurretsDestroyedBeforePlatesFall         int     `json:"kTurretsDestroyedBeforePlatesFall"`
				Kda                                       float64 `json:"kda"`
				KillAfterHiddenWithAlly                   int     `json:"killAfterHiddenWithAlly"`
				KillParticipation                         float64 `json:"killParticipation"`
				KilledChampTookFullTeamDamageSurvived     int     `json:"killedChampTookFullTeamDamageSurvived"`
				KillingSprees                             int     `json:"killingSprees"`
				KillsNearEnemyTurret                      int     `json:"killsNearEnemyTurret"`
				KillsOnOtherLanesEarlyJungleAsLaner       int     `json:"killsOnOtherLanesEarlyJungleAsLaner,omitempty"`
				KillsOnRecentlyHealedByAramPack           int     `json:"killsOnRecentlyHealedByAramPack"`
				KillsUnderOwnTurret                       int     `json:"killsUnderOwnTurret"`
				KillsWithHelpFromEpicMonster              int     `json:"killsWithHelpFromEpicMonster"`
				KnockEnemyIntoTeamAndKill                 int     `json:"knockEnemyIntoTeamAndKill"`
				LandSkillShotsEarlyGame                   int     `json:"landSkillShotsEarlyGame"`
				LaneMinionsFirst10Minutes                 int     `json:"laneMinionsFirst10Minutes"`
				LaningPhaseGoldExpAdvantage               int     `json:"laningPhaseGoldExpAdvantage"`
				LegendaryCount                            int     `json:"legendaryCount"`
				LostAnInhibitor                           int     `json:"lostAnInhibitor"`
				MaxCsAdvantageOnLaneOpponent              float64 `json:"maxCsAdvantageOnLaneOpponent"`
				MaxKillDeficit                            int     `json:"maxKillDeficit"`
				MaxLevelLeadLaneOpponent                  int     `json:"maxLevelLeadLaneOpponent"`
				MejaisFullStackInTime                     int     `json:"mejaisFullStackInTime"`
				MoreEnemyJungleThanOpponent               float64 `json:"moreEnemyJungleThanOpponent"`
				MultiKillOneSpell                         int     `json:"multiKillOneSpell"`
				MultiTurretRiftHeraldCount                int     `json:"multiTurretRiftHeraldCount"`
				Multikills                                int     `json:"multikills"`
				MultikillsAfterAggressiveFlash            int     `json:"multikillsAfterAggressiveFlash"`
				OuterTurretExecutesBefore10Minutes        int     `json:"outerTurretExecutesBefore10Minutes"`
				OutnumberedKills                          int     `json:"outnumberedKills"`
				OutnumberedNexusKill                      int     `json:"outnumberedNexusKill"`
				PerfectDragonSoulsTaken                   int     `json:"perfectDragonSoulsTaken"`
				PerfectGame                               int     `json:"perfectGame"`
				PickKillWithAlly                          int     `json:"pickKillWithAlly"`
				PlayedChampSelectPosition                 int     `json:"playedChampSelectPosition"`
				PoroExplosions                            int     `json:"poroExplosions"`
				QuickCleanse                              int     `json:"quickCleanse"`
				QuickFirstTurret                          int     `json:"quickFirstTurret"`
				QuickSoloKills                            int     `json:"quickSoloKills"`
				RiftHeraldTakedowns                       int     `json:"riftHeraldTakedowns"`
				SaveAllyFromDeath                         int     `json:"saveAllyFromDeath"`
				ScuttleCrabKills                          int     `json:"scuttleCrabKills"`
				SkillshotsDodged                          int     `json:"skillshotsDodged"`
				SkillshotsHit                             int     `json:"skillshotsHit"`
				SnowballsHit                              int     `json:"snowballsHit"`
				SoloBaronKills                            int     `json:"soloBaronKills"`
				SoloKills                                 int     `json:"soloKills"`
				SoloTurretsLategame                       int     `json:"soloTurretsLategame,omitempty"`
				StealthWardsPlaced                        int     `json:"stealthWardsPlaced"`
				SurvivedSingleDigitHpCount                int     `json:"survivedSingleDigitHpCount"`
				SurvivedThreeImmobilizesInFight           int     `json:"survivedThreeImmobilizesInFight"`
				TakedownOnFirstTurret                     int     `json:"takedownOnFirstTurret"`
				Takedowns                                 int     `json:"takedowns"`
				TakedownsAfterGainingLevelAdvantage       int     `json:"takedownsAfterGainingLevelAdvantage"`
				TakedownsBeforeJungleMinionSpawn          int     `json:"takedownsBeforeJungleMinionSpawn"`
				TakedownsFirstXMinutes                    int     `json:"takedownsFirstXMinutes"`
				TakedownsInAlcove                         int     `json:"takedownsInAlcove"`
				TakedownsInEnemyFountain                  int     `json:"takedownsInEnemyFountain"`
				TeamBaronKills                            int     `json:"teamBaronKills"`
				TeamDamagePercentage                      float64 `json:"teamDamagePercentage"`
				TeamElderDragonKills                      int     `json:"teamElderDragonKills"`
				TeamRiftHeraldKills                       int     `json:"teamRiftHeraldKills"`
				TookLargeDamageSurvived                   int     `json:"tookLargeDamageSurvived"`
				TurretPlatesTaken                         int     `json:"turretPlatesTaken"`
				TurretTakedowns                           int     `json:"turretTakedowns"`
				TurretsTakenWithRiftHerald                int     `json:"turretsTakenWithRiftHerald"`
				TwentyMinionsIn3SecondsCount              int     `json:"twentyMinionsIn3SecondsCount"`
				TwoWardsOneSweeperCount                   int     `json:"twoWardsOneSweeperCount"`
				UnseenRecalls                             int     `json:"unseenRecalls"`
				VisionScoreAdvantageLaneOpponent          float64 `json:"visionScoreAdvantageLaneOpponent"`
				VisionScorePerMinute                      float64 `json:"visionScorePerMinute"`
				WardTakedowns                             int     `json:"wardTakedowns"`
				WardTakedownsBefore20M                    int     `json:"wardTakedownsBefore20M"`
				WardsGuarded                              int     `json:"wardsGuarded"`
				ControlWardTimeCoverageInRiverOrEnemyHalf float64 `json:"controlWardTimeCoverageInRiverOrEnemyHalf,omitempty"`
				HighestWardKills                          int     `json:"highestWardKills,omitempty"`
				JunglerKillsEarlyJungle                   int     `json:"junglerKillsEarlyJungle,omitempty"`
				KillsOnLanersEarlyJungleAsJungler         int     `json:"killsOnLanersEarlyJungleAsJungler,omitempty"`
				MythicItemUsed                            int     `json:"mythicItemUsed,omitempty"`
				FirstTurretKilledTime                     float64 `json:"firstTurretKilledTime,omitempty"`
				EarliestDragonTakedown                    float64 `json:"earliestDragonTakedown,omitempty"`
				FastestLegendary                          float64 `json:"fastestLegendary,omitempty"`
				HighestChampionDamage                     int     `json:"highestChampionDamage,omitempty"`
				FasterSupportQuestCompletion              int     `json:"fasterSupportQuestCompletion,omitempty"`
				HighestCrowdControlScore                  int     `json:"highestCrowdControlScore,omitempty"`
			} `json:"challenges"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionId                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			CommandPings                int    `json:"commandPings"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			DangerPings                 int    `json:"dangerPings"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			EnemyMissingPings           int    `json:"enemyMissingPings"`
			EnemyVisionPings            int    `json:"enemyVisionPings"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GetBackPings                int    `json:"getBackPings"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			HoldPings                   int    `json:"holdPings"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			NeedVisionPings             int    `json:"needVisionPings"`
			NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
			NexusKills                  int    `json:"nexusKills"`
			NexusLost                   int    `json:"nexusLost"`
			NexusTakedowns              int    `json:"nexusTakedowns"`
			ObjectivesStolen            int    `json:"objectivesStolen"`
			ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
			OnMyWayPings                int    `json:"onMyWayPings"`
			ParticipantId               int    `json:"participantId"`
			PentaKills                  int    `json:"pentaKills"`
			Perks                       struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			Placement                      int    `json:"placement"`
			PlayerAugment1                 int    `json:"playerAugment1"`
			PlayerAugment2                 int    `json:"playerAugment2"`
			PlayerAugment3                 int    `json:"playerAugment3"`
			PlayerAugment4                 int    `json:"playerAugment4"`
			PlayerSubteamId                int    `json:"playerSubteamId"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RiotIdName                     string `json:"riotIdName"`
			RiotIdTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			SubteamPlacement               int    `json:"subteamPlacement"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1Id                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2Id                    int    `json:"summoner2Id"`
			SummonerId                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamId                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalAllyJungleMinionsKilled   int    `json:"totalAllyJungleMinionsKilled"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalEnemyJungleMinionsKilled  int    `json:"totalEnemyJungleMinionsKilled"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionClearedPings             int    `json:"visionClearedPings"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
		} `json:"participants"`
		PlatformId string `json:"platformId"`
		QueueId    int    `json:"queueId"`
		Teams      []struct {
			Bans []struct {
				ChampionId int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Objectives struct {
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamId int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}
