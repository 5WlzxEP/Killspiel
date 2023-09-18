package RiotApi

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"slices"
)

type LoL struct {
	ApiKey       string `json:"apiKey"`
	ProfileIcon  int    `json:"profileIcon"`
	Kategorie    string `json:"kategorie"`
	Server       string `json:"server"`
	SummonerName string `json:"summonerName"`
	region       string
}

var lolKategorien = []string{"AllInPings", "AssistMePings", "Assists", "BaitPings", "BaronKills", "BasicPings", "BountyLevel", "AssistStreakCount", "AbilityUses", "AcesBefore15Minutes", "AlliedJungleMonsterKills", "BaronTakedowns", "BlastConeOppositeOpponentCount", "BountyGold", "BuffsStolen", "CompleteSupportQuestInTime", "ControlWardsPlaced", "DamagePerMinute", "DamageTakenOnTeamPercentage", "DancedWithRiftHerald", "DeathsByEnemyChamps", "DodgeSkillShotsSmallWindow", "DoubleAces", "DragonTakedowns", "EarlyLaningPhaseGoldExpAdvantage", "EffectiveHealAndShielding", "ElderDragonKillsWithOpposingSoul", "ElderDragonMultikills", "EnemyChampionImmobilizations", "EnemyJungleMonsterKills", "EpicMonsterKillsNearEnemyJungler", "EpicMonsterKillsWithin30SecondsOfSpawn", "EpicMonsterSteals", "EpicMonsterStolenWithoutSmite", "FirstTurretKilled", "FlawlessAces", "FullTeamTakedown", "GameLength", "GetTakedownsInAllLanesEarlyJungleAsLaner", "GoldPerMinute", "HadOpenNexus", "ImmobilizeAndKillWithAlly", "InitialBuffCount", "InitialCrabCount", "JungleCsBefore10Minutes", "JunglerTakedownsNearDamagedEpicMonster", "KTurretsDestroyedBeforePlatesFall", "Kda", "KillAfterHiddenWithAlly", "KillParticipation", "KilledChampTookFullTeamDamageSurvived", "KillingSprees", "KillsNearEnemyTurret", "KillsOnOtherLanesEarlyJungleAsLaner", "KillsOnRecentlyHealedByAramPack", "KillsUnderOwnTurret", "KillsWithHelpFromEpicMonster", "KnockEnemyIntoTeamAndKill", "LandSkillShotsEarlyGame", "LaneMinionsFirst10Minutes", "LaningPhaseGoldExpAdvantage", "LegendaryCount", "LostAnInhibitor", "MaxCsAdvantageOnLaneOpponent", "MaxKillDeficit", "MaxLevelLeadLaneOpponent", "MejaisFullStackInTime", "MoreEnemyJungleThanOpponent", "MultiKillOneSpell", "MultiTurretRiftHeraldCount", "Multikills", "MultikillsAfterAggressiveFlash", "OuterTurretExecutesBefore10Minutes", "OutnumberedKills", "OutnumberedNexusKill", "PerfectDragonSoulsTaken", "PerfectGame", "PickKillWithAlly", "PlayedChampSelectPosition", "PoroExplosions", "QuickCleanse", "QuickFirstTurret", "QuickSoloKills", "RiftHeraldTakedowns", "SaveAllyFromDeath", "ScuttleCrabKills", "SkillshotsDodged", "SkillshotsHit", "SnowballsHit", "SoloBaronKills", "SoloKills", "SoloTurretsLategame", "StealthWardsPlaced", "SurvivedSingleDigitHpCount", "SurvivedThreeImmobilizesInFight", "TakedownOnFirstTurret", "Takedowns", "TakedownsAfterGainingLevelAdvantage", "TakedownsBeforeJungleMinionSpawn", "TakedownsFirstXMinutes", "TakedownsInAlcove", "TakedownsInEnemyFountain", "TeamBaronKills", "TeamDamagePercentage", "TeamElderDragonKills", "TeamRiftHeraldKills", "TookLargeDamageSurvived", "TurretPlatesTaken", "TurretTakedowns", "TurretsTakenWithRiftHerald", "TwentyMinionsIn3SecondsCount", "TwoWardsOneSweeperCount", "UnseenRecalls", "VisionScoreAdvantageLaneOpponent", "VisionScorePerMinute", "WardTakedowns", "WardTakedownsBefore20M", "WardsGuarded", "ControlWardTimeCoverageInRiverOrEnemyHalf", "HighestWardKills", "JunglerKillsEarlyJungle", "KillsOnLanersEarlyJungleAsJungler", "MythicItemUsed", "FirstTurretKilledTime", "EarliestDragonTakedown", "FastestLegendary", "HighestChampionDamage", "FasterSupportQuestCompletion", "HighestCrowdControlScore", "ChampExperience", "ChampLevel", "CommandPings", "ConsumablesPurchased", "DamageDealtToBuildings", "DamageDealtToObjectives", "DamageDealtToTurrets", "DamageSelfMitigated", "DangerPings", "Deaths", "DetectorWardsPlaced", "DoubleKills", "DragonKills", "EligibleForProgression", "EnemyMissingPings", "EnemyVisionPings", "FirstBloodAssist", "FirstBloodKill", "FirstTowerAssist", "FirstTowerKill", "GameEndedInEarlySurrender", "GameEndedInSurrender", "GetBackPings", "GoldEarned", "GoldSpent", "HoldPings", "InhibitorKills", "InhibitorTakedowns", "InhibitorsLost", "ItemsPurchased", "KillingSprees", "Kills", "LargestCriticalStrike", "LargestKillingSpree", "LargestMultiKill", "LongestTimeSpentLiving", "MagicDamageDealt", "MagicDamageDealtToChampions", "MagicDamageTaken", "NeedVisionPings", "NeutralMinionsKilled", "NexusKills", "NexusLost", "NexusTakedowns", "ObjectivesStolen", "ObjectivesStolenAssists", "OnMyWayPings", "ParticipantId", "PentaKills", "PhysicalDamageDealt", "PhysicalDamageDealtToChampions", "PhysicalDamageTaken", "Placement", "PlayerAugment1", "PlayerAugment2", "PlayerAugment3", "PlayerAugment4", "PlayerSubteamId", "ProfileIcon", "PushPings", "QuadraKills", "SightWardsBoughtInGame", "Spell1Casts", "Spell2Casts", "Spell3Casts", "Spell4Casts", "SubteamPlacement", "Summoner1Casts", "Summoner1Id", "Summoner2Casts", "Summoner2Id", "SummonerLevel", "TeamEarlySurrendered", "TeamId", "TimeCCingOthers", "TimePlayed", "TotalAllyJungleMinionsKilled", "TotalDamageDealt", "TotalDamageDealtToChampions", "TotalDamageShieldedOnTeammates", "TotalDamageTaken", "TotalEnemyJungleMinionsKilled", "TotalHeal", "TotalHealsOnTeammates", "TotalMinionsKilled", "TotalTimeCCDealt", "TotalTimeSpentDead", "TotalUnitsHealed", "TripleKills", "TrueDamageDealt", "TrueDamageDealtToChampions", "TrueDamageTaken", "TurretKills", "TurretTakedowns", "TurretsLost", "UnrealKills", "VisionClearedPings", "VisionScore", "VisionWardsBoughtInGame", "WardsKilled", "WardsPlaced", "Win"}

func (a *Api) postLoL(ctx *fiber.Ctx) error {
	var lol LoL
	err := ctx.BodyParser(&lol)
	if err != nil {
		return err
	}

	if !isServer(lol.Server) {
		return ctx.Status(http.StatusBadRequest).SendString("Server ist nicht valide.")
	}
	if !validKey(lol.ApiKey, a.client) {
		return ctx.Status(http.StatusBadRequest).SendString("Api Key nicht valide.")
	}
	summoner, err := getLoLSummonerByName(lol.SummonerName, lol.Server, lol.ApiKey, a.client)
	if err != nil || summoner == nil {
		return ctx.Status(http.StatusBadRequest).SendString("Beschwörername nicht valide.")
	}

	a.LoL = lol
	a.setRegion()

	a.summoner = summoner
	a.ready = true
	a.save()

	return ctx.SendStatus(http.StatusNoContent)
}

// isLoLKategorie returns if something (first value) is a valid lolKategorie. If no value is given checks if LoL.Kategorie is a lolKategorie
func (a *Api) isLoLKategorie(kategorie ...string) bool {
	if len(kategorie) > 0 {
		return slices.Contains(lolKategorien, kategorie[0])
	}
	return slices.Contains(lolKategorien, a.LoL.Kategorie)
}

func (a *Api) getValueFloat(game *MatchDto) float64 {
	v := a.getValue(game)
	switch v.(type) {
	case int:
		return float64(v.(int))
	case float64:
		return v.(float64)
	case bool:
		if v.(bool) {
			return 1
		}
		return 0
	}
	return 0
}

func (a *Api) getValue(game *MatchDto) any {
	idx := slices.Index(game.Metadata.Participants, a.currentSummoner.Puuid)
	if idx == -1 {
		return nil
	}

	player := game.Info.Participants[idx]

	switch a.LoL.Kategorie {
	case "AllInPings":
		return player.AllInPings
	case "AssistMePings":
		return player.AssistMePings
	case "Assists":
		return player.Assists
	case "BaitPings":
		return player.BaitPings
	case "BaronKills":
		return player.BaronKills
	case "BasicPings":
		return player.BasicPings
	case "BountyLevel":
		return player.BountyLevel
	case "AssistStreakCount":
		return player.Challenges.AssistStreakCount
	case "AbilityUses":
		return player.Challenges.AbilityUses
	case "AcesBefore15Minutes":
		return player.Challenges.AcesBefore15Minutes
	case "AlliedJungleMonsterKills":
		return player.Challenges.AlliedJungleMonsterKills
	case "BaronTakedowns":
		return player.Challenges.BaronTakedowns
	case "BlastConeOppositeOpponentCount":
		return player.Challenges.BlastConeOppositeOpponentCount
	case "BountyGold":
		return player.Challenges.BountyGold
	case "BuffsStolen":
		return player.Challenges.BuffsStolen
	case "CompleteSupportQuestInTime":
		return player.Challenges.CompleteSupportQuestInTime
	case "ControlWardsPlaced":
		return player.Challenges.ControlWardsPlaced
	case "DamagePerMinute":
		return player.Challenges.DamagePerMinute
	case "DamageTakenOnTeamPercentage":
		return player.Challenges.DamageTakenOnTeamPercentage
	case "DancedWithRiftHerald":
		return player.Challenges.DancedWithRiftHerald
	case "DeathsByEnemyChamps":
		return player.Challenges.DeathsByEnemyChamps
	case "DodgeSkillShotsSmallWindow":
		return player.Challenges.DodgeSkillShotsSmallWindow
	case "DoubleAces":
		return player.Challenges.DoubleAces
	case "DragonTakedowns":
		return player.Challenges.DragonTakedowns
	case "EarlyLaningPhaseGoldExpAdvantage":
		return player.Challenges.EarlyLaningPhaseGoldExpAdvantage
	case "EffectiveHealAndShielding":
		return player.Challenges.EffectiveHealAndShielding
	case "ElderDragonKillsWithOpposingSoul":
		return player.Challenges.ElderDragonKillsWithOpposingSoul
	case "ElderDragonMultikills":
		return player.Challenges.ElderDragonMultikills
	case "EnemyChampionImmobilizations":
		return player.Challenges.EnemyChampionImmobilizations
	case "EnemyJungleMonsterKills":
		return player.Challenges.EnemyJungleMonsterKills
	case "EpicMonsterKillsNearEnemyJungler":
		return player.Challenges.EpicMonsterKillsNearEnemyJungler
	case "EpicMonsterKillsWithin30SecondsOfSpawn":
		return player.Challenges.EpicMonsterKillsWithin30SecondsOfSpawn
	case "EpicMonsterSteals":
		return player.Challenges.EpicMonsterSteals
	case "EpicMonsterStolenWithoutSmite":
		return player.Challenges.EpicMonsterStolenWithoutSmite
	case "FirstTurretKilled":
		return player.Challenges.FirstTurretKilled
	case "FlawlessAces":
		return player.Challenges.FlawlessAces
	case "FullTeamTakedown":
		return player.Challenges.FullTeamTakedown
	case "GameLength":
		return player.Challenges.GameLength
	case "GetTakedownsInAllLanesEarlyJungleAsLaner":
		return player.Challenges.GetTakedownsInAllLanesEarlyJungleAsLaner
	case "GoldPerMinute":
		return player.Challenges.GoldPerMinute
	case "HadOpenNexus":
		return player.Challenges.HadOpenNexus
	case "ImmobilizeAndKillWithAlly":
		return player.Challenges.ImmobilizeAndKillWithAlly
	case "InitialBuffCount":
		return player.Challenges.InitialBuffCount
	case "InitialCrabCount":
		return player.Challenges.InitialCrabCount
	case "JungleCsBefore10Minutes":
		return player.Challenges.JungleCsBefore10Minutes
	case "JunglerTakedownsNearDamagedEpicMonster":
		return player.Challenges.JunglerTakedownsNearDamagedEpicMonster
	case "KTurretsDestroyedBeforePlatesFall":
		return player.Challenges.KTurretsDestroyedBeforePlatesFall
	case "Kda":
		return player.Challenges.Kda
	case "KillAfterHiddenWithAlly":
		return player.Challenges.KillAfterHiddenWithAlly
	case "KillParticipation":
		return player.Challenges.KillParticipation
	case "KilledChampTookFullTeamDamageSurvived":
		return player.Challenges.KilledChampTookFullTeamDamageSurvived
	case "KillsNearEnemyTurret":
		return player.Challenges.KillsNearEnemyTurret
	case "KillsOnOtherLanesEarlyJungleAsLaner":
		return player.Challenges.KillsOnOtherLanesEarlyJungleAsLaner
	case "KillsOnRecentlyHealedByAramPack":
		return player.Challenges.KillsOnRecentlyHealedByAramPack
	case "KillsUnderOwnTurret":
		return player.Challenges.KillsUnderOwnTurret
	case "KillsWithHelpFromEpicMonster":
		return player.Challenges.KillsWithHelpFromEpicMonster
	case "KnockEnemyIntoTeamAndKill":
		return player.Challenges.KnockEnemyIntoTeamAndKill
	case "LandSkillShotsEarlyGame":
		return player.Challenges.LandSkillShotsEarlyGame
	case "LaneMinionsFirst10Minutes":
		return player.Challenges.LaneMinionsFirst10Minutes
	case "LaningPhaseGoldExpAdvantage":
		return player.Challenges.LaningPhaseGoldExpAdvantage
	case "LegendaryCount":
		return player.Challenges.LegendaryCount
	case "LostAnInhibitor":
		return player.Challenges.LostAnInhibitor
	case "MaxCsAdvantageOnLaneOpponent":
		return player.Challenges.MaxCsAdvantageOnLaneOpponent
	case "MaxKillDeficit":
		return player.Challenges.MaxKillDeficit
	case "MaxLevelLeadLaneOpponent":
		return player.Challenges.MaxLevelLeadLaneOpponent
	case "MejaisFullStackInTime":
		return player.Challenges.MejaisFullStackInTime
	case "MoreEnemyJungleThanOpponent":
		return player.Challenges.MoreEnemyJungleThanOpponent
	case "MultiKillOneSpell":
		return player.Challenges.MultiKillOneSpell
	case "MultiTurretRiftHeraldCount":
		return player.Challenges.MultiTurretRiftHeraldCount
	case "Multikills":
		return player.Challenges.Multikills
	case "MultikillsAfterAggressiveFlash":
		return player.Challenges.MultikillsAfterAggressiveFlash
	case "OuterTurretExecutesBefore10Minutes":
		return player.Challenges.OuterTurretExecutesBefore10Minutes
	case "OutnumberedKills":
		return player.Challenges.OutnumberedKills
	case "OutnumberedNexusKill":
		return player.Challenges.OutnumberedNexusKill
	case "PerfectDragonSoulsTaken":
		return player.Challenges.PerfectDragonSoulsTaken
	case "PerfectGame":
		return player.Challenges.PerfectGame
	case "PickKillWithAlly":
		return player.Challenges.PickKillWithAlly
	case "PlayedChampSelectPosition":
		return player.Challenges.PlayedChampSelectPosition
	case "PoroExplosions":
		return player.Challenges.PoroExplosions
	case "QuickCleanse":
		return player.Challenges.QuickCleanse
	case "QuickFirstTurret":
		return player.Challenges.QuickFirstTurret
	case "QuickSoloKills":
		return player.Challenges.QuickSoloKills
	case "RiftHeraldTakedowns":
		return player.Challenges.RiftHeraldTakedowns
	case "SaveAllyFromDeath":
		return player.Challenges.SaveAllyFromDeath
	case "ScuttleCrabKills":
		return player.Challenges.ScuttleCrabKills
	case "SkillshotsDodged":
		return player.Challenges.SkillshotsDodged
	case "SkillshotsHit":
		return player.Challenges.SkillshotsHit
	case "SnowballsHit":
		return player.Challenges.SnowballsHit
	case "SoloBaronKills":
		return player.Challenges.SoloBaronKills
	case "SoloKills":
		return player.Challenges.SoloKills
	case "SoloTurretsLategame":
		return player.Challenges.SoloTurretsLategame
	case "StealthWardsPlaced":
		return player.Challenges.StealthWardsPlaced
	case "SurvivedSingleDigitHpCount":
		return player.Challenges.SurvivedSingleDigitHpCount
	case "SurvivedThreeImmobilizesInFight":
		return player.Challenges.SurvivedThreeImmobilizesInFight
	case "TakedownOnFirstTurret":
		return player.Challenges.TakedownOnFirstTurret
	case "Takedowns":
		return player.Challenges.Takedowns
	case "TakedownsAfterGainingLevelAdvantage":
		return player.Challenges.TakedownsAfterGainingLevelAdvantage
	case "TakedownsBeforeJungleMinionSpawn":
		return player.Challenges.TakedownsBeforeJungleMinionSpawn
	case "TakedownsFirstXMinutes":
		return player.Challenges.TakedownsFirstXMinutes
	case "TakedownsInAlcove":
		return player.Challenges.TakedownsInAlcove
	case "TakedownsInEnemyFountain":
		return player.Challenges.TakedownsInEnemyFountain
	case "TeamBaronKills":
		return player.Challenges.TeamBaronKills
	case "TeamDamagePercentage":
		return player.Challenges.TeamDamagePercentage
	case "TeamElderDragonKills":
		return player.Challenges.TeamElderDragonKills
	case "TeamRiftHeraldKills":
		return player.Challenges.TeamRiftHeraldKills
	case "TookLargeDamageSurvived":
		return player.Challenges.TookLargeDamageSurvived
	case "TurretPlatesTaken":
		return player.Challenges.TurretPlatesTaken
	case "TurretsTakenWithRiftHerald":
		return player.Challenges.TurretsTakenWithRiftHerald
	case "TwentyMinionsIn3SecondsCount":
		return player.Challenges.TwentyMinionsIn3SecondsCount
	case "TwoWardsOneSweeperCount":
		return player.Challenges.TwoWardsOneSweeperCount
	case "UnseenRecalls":
		return player.Challenges.UnseenRecalls
	case "VisionScoreAdvantageLaneOpponent":
		return player.Challenges.VisionScoreAdvantageLaneOpponent
	case "VisionScorePerMinute":
		return player.Challenges.VisionScorePerMinute
	case "WardTakedowns":
		return player.Challenges.WardTakedowns
	case "WardTakedownsBefore20M":
		return player.Challenges.WardTakedownsBefore20M
	case "WardsGuarded":
		return player.Challenges.WardsGuarded
	case "ControlWardTimeCoverageInRiverOrEnemyHalf":
		return player.Challenges.ControlWardTimeCoverageInRiverOrEnemyHalf
	case "HighestWardKills":
		return player.Challenges.HighestWardKills
	case "JunglerKillsEarlyJungle":
		return player.Challenges.JunglerKillsEarlyJungle
	case "KillsOnLanersEarlyJungleAsJungler":
		return player.Challenges.KillsOnLanersEarlyJungleAsJungler
	case "MythicItemUsed":
		return player.Challenges.MythicItemUsed
	case "FirstTurretKilledTime":
		return player.Challenges.FirstTurretKilledTime
	case "EarliestDragonTakedown":
		return player.Challenges.EarliestDragonTakedown
	case "FastestLegendary":
		return player.Challenges.FastestLegendary
	case "HighestChampionDamage":
		return player.Challenges.HighestChampionDamage
	case "FasterSupportQuestCompletion":
		return player.Challenges.FasterSupportQuestCompletion
	case "HighestCrowdControlScore":
		return player.Challenges.HighestCrowdControlScore
	case "ChampExperience":
		return player.ChampExperience
	case "ChampLevel":
		return player.ChampLevel
	case "CommandPings":
		return player.CommandPings
	case "ConsumablesPurchased":
		return player.ConsumablesPurchased
	case "DamageDealtToBuildings":
		return player.DamageDealtToBuildings
	case "DamageDealtToObjectives":
		return player.DamageDealtToObjectives
	case "DamageDealtToTurrets":
		return player.DamageDealtToTurrets
	case "DamageSelfMitigated":
		return player.DamageSelfMitigated
	case "DangerPings":
		return player.DangerPings
	case "Deaths":
		return player.Deaths
	case "DetectorWardsPlaced":
		return player.DetectorWardsPlaced
	case "DoubleKills":
		return player.DoubleKills
	case "DragonKills":
		return player.DragonKills
	case "EligibleForProgression":
		return player.EligibleForProgression
	case "EnemyMissingPings":
		return player.EnemyMissingPings
	case "EnemyVisionPings":
		return player.EnemyVisionPings
	case "FirstBloodAssist":
		return player.FirstBloodAssist
	case "FirstBloodKill":
		return player.FirstBloodKill
	case "FirstTowerAssist":
		return player.FirstTowerAssist
	case "FirstTowerKill":
		return player.FirstTowerKill
	case "GameEndedInEarlySurrender":
		return player.GameEndedInEarlySurrender
	case "GameEndedInSurrender":
		return player.GameEndedInSurrender
	case "GetBackPings":
		return player.GetBackPings
	case "GoldEarned":
		return player.GoldEarned
	case "GoldSpent":
		return player.GoldSpent
	case "HoldPings":
		return player.HoldPings
	case "InhibitorKills":
		return player.InhibitorKills
	case "InhibitorTakedowns":
		return player.InhibitorTakedowns
	case "InhibitorsLost":
		return player.InhibitorsLost
	case "ItemsPurchased":
		return player.ItemsPurchased
	case "KillingSprees":
		return player.KillingSprees
	case "Kills":
		return player.Kills
	case "LargestCriticalStrike":
		return player.LargestCriticalStrike
	case "LargestKillingSpree":
		return player.LargestKillingSpree
	case "LargestMultiKill":
		return player.LargestMultiKill
	case "LongestTimeSpentLiving":
		return player.LongestTimeSpentLiving
	case "MagicDamageDealt":
		return player.MagicDamageDealt
	case "MagicDamageDealtToChampions":
		return player.MagicDamageDealtToChampions
	case "MagicDamageTaken":
		return player.MagicDamageTaken
	case "NeedVisionPings":
		return player.NeedVisionPings
	case "NeutralMinionsKilled":
		return player.NeutralMinionsKilled
	case "NexusKills":
		return player.NexusKills
	case "NexusLost":
		return player.NexusLost
	case "NexusTakedowns":
		return player.NexusTakedowns
	case "ObjectivesStolen":
		return player.ObjectivesStolen
	case "ObjectivesStolenAssists":
		return player.ObjectivesStolenAssists
	case "OnMyWayPings":
		return player.OnMyWayPings
	case "ParticipantId":
		return player.ParticipantId
	case "PentaKills":
		return player.PentaKills
	case "PhysicalDamageDealt":
		return player.PhysicalDamageDealt
	case "PhysicalDamageDealtToChampions":
		return player.PhysicalDamageDealtToChampions
	case "PhysicalDamageTaken":
		return player.PhysicalDamageTaken
	case "Placement":
		return player.Placement
	case "PlayerAugment1":
		return player.PlayerAugment1
	case "PlayerAugment2":
		return player.PlayerAugment2
	case "PlayerAugment3":
		return player.PlayerAugment3
	case "PlayerAugment4":
		return player.PlayerAugment4
	case "PlayerSubteamId":
		return player.PlayerSubteamId
	case "ProfileIcon":
		return player.ProfileIcon
	case "PushPings":
		return player.PushPings
	case "QuadraKills":
		return player.QuadraKills
	case "SightWardsBoughtInGame":
		return player.SightWardsBoughtInGame
	case "Spell1Casts":
		return player.Spell1Casts
	case "Spell2Casts":
		return player.Spell2Casts
	case "Spell3Casts":
		return player.Spell3Casts
	case "Spell4Casts":
		return player.Spell4Casts
	case "SubteamPlacement":
		return player.SubteamPlacement
	case "Summoner1Casts":
		return player.Summoner1Casts
	case "Summoner1Id":
		return player.Summoner1Id
	case "Summoner2Casts":
		return player.Summoner2Casts
	case "Summoner2Id":
		return player.Summoner2Id
	case "SummonerLevel":
		return player.SummonerLevel
	case "TeamEarlySurrendered":
		return player.TeamEarlySurrendered
	case "TeamId":
		return player.TeamId
	case "TimeCCingOthers":
		return player.TimeCCingOthers
	case "TimePlayed":
		return player.TimePlayed
	case "TotalAllyJungleMinionsKilled":
		return player.TotalAllyJungleMinionsKilled
	case "TotalDamageDealt":
		return player.TotalDamageDealt
	case "TotalDamageDealtToChampions":
		return player.TotalDamageDealtToChampions
	case "TotalDamageShieldedOnTeammates":
		return player.TotalDamageShieldedOnTeammates
	case "TotalDamageTaken":
		return player.TotalDamageTaken
	case "TotalEnemyJungleMinionsKilled":
		return player.TotalEnemyJungleMinionsKilled
	case "TotalHeal":
		return player.TotalHeal
	case "TotalHealsOnTeammates":
		return player.TotalHealsOnTeammates
	case "TotalMinionsKilled":
		return player.TotalMinionsKilled
	case "TotalTimeCCDealt":
		return player.TotalTimeCCDealt
	case "TotalTimeSpentDead":
		return player.TotalTimeSpentDead
	case "TotalUnitsHealed":
		return player.TotalUnitsHealed
	case "TripleKills":
		return player.TripleKills
	case "TrueDamageDealt":
		return player.TrueDamageDealt
	case "TrueDamageDealtToChampions":
		return player.TrueDamageDealtToChampions
	case "TrueDamageTaken":
		return player.TrueDamageTaken
	case "TurretKills":
		return player.TurretKills
	case "TurretTakedowns":
		return player.TurretTakedowns
	case "TurretsLost":
		return player.TurretsLost
	case "UnrealKills":
		return player.UnrealKills
	case "VisionClearedPings":
		return player.VisionClearedPings
	case "VisionScore":
		return player.VisionScore
	case "VisionWardsBoughtInGame":
		return player.VisionWardsBoughtInGame
	case "WardsKilled":
		return player.WardsKilled
	case "WardsPlaced":
		return player.WardsPlaced
	case "Win":
		return player.Win
	}

	log.Println("No valid LoLKategorie chosen.")

	return 0
}
