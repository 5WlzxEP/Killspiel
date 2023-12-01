<script lang="ts">
	import {
		Autocomplete,
		type AutocompleteOption,
		type PopupSettings,
		popup,
		getModalStore,
		type ModalSettings
	} from "@skeletonlabs/skeleton"
	import { IconHelp } from "@tabler/icons-svelte"

	const modalStore = getModalStore()

	const allValues = [
		"AllInPings",
		"AssistMePings",
		"Assists",
		"BaitPings",
		"BaronKills",
		"BasicPings",
		"BountyLevel",
		"AssistStreakCount",
		"AbilityUses",
		"AcesBefore15Minutes",
		"AlliedJungleMonsterKills",
		"BaronTakedowns",
		"BlastConeOppositeOpponentCount",
		"BountyGold",
		"BuffsStolen",
		"CompleteSupportQuestInTime",
		"ControlWardsPlaced",
		"DamagePerMinute",
		"DamageTakenOnTeamPercentage",
		"DancedWithRiftHerald",
		"DeathsByEnemyChamps",
		"DodgeSkillShotsSmallWindow",
		"DoubleAces",
		"DragonTakedowns",
		"EarlyLaningPhaseGoldExpAdvantage",
		"EffectiveHealAndShielding",
		"ElderDragonKillsWithOpposingSoul",
		"ElderDragonMultikills",
		"EnemyChampionImmobilizations",
		"EnemyJungleMonsterKills",
		"EpicMonsterKillsNearEnemyJungler",
		"EpicMonsterKillsWithin30SecondsOfSpawn",
		"EpicMonsterSteals",
		"EpicMonsterStolenWithoutSmite",
		"FirstTurretKilled",
		"FlawlessAces",
		"FullTeamTakedown",
		"GameLength",
		"GetTakedownsInAllLanesEarlyJungleAsLaner",
		"GoldPerMinute",
		"HadOpenNexus",
		"ImmobilizeAndKillWithAlly",
		"InitialBuffCount",
		"InitialCrabCount",
		"JungleCsBefore10Minutes",
		"JunglerTakedownsNearDamagedEpicMonster",
		"KTurretsDestroyedBeforePlatesFall",
		"Kda",
		"KillAfterHiddenWithAlly",
		"KillParticipation",
		"KilledChampTookFullTeamDamageSurvived",
		"KillingSprees",
		"KillsNearEnemyTurret",
		"KillsOnOtherLanesEarlyJungleAsLaner",
		"KillsOnRecentlyHealedByAramPack",
		"KillsUnderOwnTurret",
		"KillsWithHelpFromEpicMonster",
		"KnockEnemyIntoTeamAndKill",
		"LandSkillShotsEarlyGame",
		"LaneMinionsFirst10Minutes",
		"LaningPhaseGoldExpAdvantage",
		"LegendaryCount",
		"LostAnInhibitor",
		"MaxCsAdvantageOnLaneOpponent",
		"MaxKillDeficit",
		"MaxLevelLeadLaneOpponent",
		"MejaisFullStackInTime",
		"MoreEnemyJungleThanOpponent",
		"MultiKillOneSpell",
		"MultiTurretRiftHeraldCount",
		"Multikills",
		"MultikillsAfterAggressiveFlash",
		"OuterTurretExecutesBefore10Minutes",
		"OutnumberedKills",
		"OutnumberedNexusKill",
		"PerfectDragonSoulsTaken",
		"PerfectGame",
		"PickKillWithAlly",
		"PlayedChampSelectPosition",
		"PoroExplosions",
		"QuickCleanse",
		"QuickFirstTurret",
		"QuickSoloKills",
		"RiftHeraldTakedowns",
		"SaveAllyFromDeath",
		"ScuttleCrabKills",
		"SkillshotsDodged",
		"SkillshotsHit",
		"SnowballsHit",
		"SoloBaronKills",
		"SoloKills",
		"SoloTurretsLategame",
		"StealthWardsPlaced",
		"SurvivedSingleDigitHpCount",
		"SurvivedThreeImmobilizesInFight",
		"TakedownOnFirstTurret",
		"Takedowns",
		"TakedownsAfterGainingLevelAdvantage",
		"TakedownsBeforeJungleMinionSpawn",
		"TakedownsFirstXMinutes",
		"TakedownsInAlcove",
		"TakedownsInEnemyFountain",
		"TeamBaronKills",
		"TeamDamagePercentage",
		"TeamElderDragonKills",
		"TeamRiftHeraldKills",
		"TookLargeDamageSurvived",
		"TurretPlatesTaken",
		"TurretTakedowns",
		"TurretsTakenWithRiftHerald",
		"TwentyMinionsIn3SecondsCount",
		"TwoWardsOneSweeperCount",
		"UnseenRecalls",
		"VisionScoreAdvantageLaneOpponent",
		"VisionScorePerMinute",
		"WardTakedowns",
		"WardTakedownsBefore20M",
		"WardsGuarded",
		"ControlWardTimeCoverageInRiverOrEnemyHalf",
		"HighestWardKills",
		"JunglerKillsEarlyJungle",
		"KillsOnLanersEarlyJungleAsJungler",
		"MythicItemUsed",
		"FirstTurretKilledTime",
		"EarliestDragonTakedown",
		"FastestLegendary",
		"HighestChampionDamage",
		"FasterSupportQuestCompletion",
		"HighestCrowdControlScore",
		"ChampExperience",
		"ChampLevel",
		"CommandPings",
		"ConsumablesPurchased",
		"DamageDealtToBuildings",
		"DamageDealtToObjectives",
		"DamageDealtToTurrets",
		"DamageSelfMitigated",
		"DangerPings",
		"Deaths",
		"DetectorWardsPlaced",
		"DoubleKills",
		"DragonKills",
		"EligibleForProgression",
		"EnemyMissingPings",
		"EnemyVisionPings",
		"FirstBloodAssist",
		"FirstBloodKill",
		"FirstTowerAssist",
		"FirstTowerKill",
		"GameEndedInEarlySurrender",
		"GameEndedInSurrender",
		"GetBackPings",
		"GoldEarned",
		"GoldSpent",
		"HoldPings",
		"InhibitorKills",
		"InhibitorTakedowns",
		"InhibitorsLost",
		"ItemsPurchased",
		"Kills",
		"LargestCriticalStrike",
		"LargestKillingSpree",
		"LargestMultiKill",
		"LongestTimeSpentLiving",
		"MagicDamageDealt",
		"MagicDamageDealtToChampions",
		"MagicDamageTaken",
		"NeedVisionPings",
		"NeutralMinionsKilled",
		"NexusKills",
		"NexusLost",
		"NexusTakedowns",
		"ObjectivesStolen",
		"ObjectivesStolenAssists",
		"OnMyWayPings",
		"ParticipantId",
		"PentaKills",
		"PhysicalDamageDealt",
		"PhysicalDamageDealtToChampions",
		"PhysicalDamageTaken",
		"Placement",
		"PlayerAugment1",
		"PlayerAugment2",
		"PlayerAugment3",
		"PlayerAugment4",
		"PlayerSubteamId",
		"ProfileIcon",
		"PushPings",
		"QuadraKills",
		"SightWardsBoughtInGame",
		"Spell1Casts",
		"Spell2Casts",
		"Spell3Casts",
		"Spell4Casts",
		"SubteamPlacement",
		"Summoner1Casts",
		"Summoner1Id",
		"Summoner2Casts",
		"Summoner2Id",
		"SummonerLevel",
		"TeamEarlySurrendered",
		"TeamId",
		"TimeCCingOthers",
		"TimePlayed",
		"TotalAllyJungleMinionsKilled",
		"TotalDamageDealt",
		"TotalDamageDealtToChampions",
		"TotalDamageShieldedOnTeammates",
		"TotalDamageTaken",
		"TotalEnemyJungleMinionsKilled",
		"TotalHeal",
		"TotalHealsOnTeammates",
		"TotalMinionsKilled",
		"TotalTimeCCDealt",
		"TotalTimeSpentDead",
		"TotalUnitsHealed",
		"TripleKills",
		"TrueDamageDealt",
		"TrueDamageDealtToChampions",
		"TrueDamageTaken",
		"TurretKills",
		"TurretsLost",
		"UnrealKills",
		"VisionClearedPings",
		"VisionScore",
		"VisionWardsBoughtInGame",
		"WardsKilled",
		"WardsPlaced",
		"Win"
	]
	const allOptions = allValues.map((v) => {
		return { label: v, value: v }
	})

	const basicOptions = ["Assists", "Deaths", "Kda", "Kills"].map((v) => {
		return { label: v, value: v }
	})

	let basic = true

	function loadAllOptions() {
		basic = false
		options = allOptions
	}

	function onSelection(event: CustomEvent<AutocompleteOption>): void {
		if (allValues.includes(event.detail.label)) value = event.detail.label
	}

	const modal: ModalSettings = {
		type: "alert",
		title: "Kategorie",
		body: "Wählt aus, welcher Wert genutzt werden soll.",
		buttonTextCancel: "Schließen"
	}

	let options: AutocompleteOption[] = basicOptions

	let popupSettings: PopupSettings = {
		event: "focus-click",
		target: "popupAutocomplete",
		placement: "bottom"
	}

	export let value: string
</script>

<label class="label">
	<span
		>Kategorie zum Auswerten
		<div
			role="button"
			tabindex="-1"
			on:keypress
			on:click={() => {
				modalStore.trigger(modal)
			}}
			class="p-1 cursor-pointer inline-block"
		>
			<IconHelp class="inline-block" />
		</div>
	</span>
	<div class="input-group input-group-divider grid-cols-[1fr_auto]">
		<input
			class="autocomplete"
			type="search"
			bind:value
			placeholder="Kills"
			use:popup={popupSettings}
		/>
		<div class="input-group-shim">*</div>
	</div>
	<div data-popup="popupAutocomplete" class="card w-[35%]">
		<div class="card w-full max-h-52 p-4 overflow-y-auto" tabindex="-1">
			<Autocomplete bind:input={value} {options} on:selection={onSelection} />
			{#if basic}
				<hr class="m-1" />
				<div class="autocomplete">
					<nav class="autocomplete-nav">
						<ul class="autocomplete-list list-nav">
							<li class="autocomplete-item">
								<button
									class="autocomplete-button w-full anchor"
									type="button"
									on:click|preventDefault={loadAllOptions}
								>
									Lade alle Optionen
								</button>
							</li>
						</ul>
					</nav>
				</div>
			{/if}
		</div>
	</div>
</label>
