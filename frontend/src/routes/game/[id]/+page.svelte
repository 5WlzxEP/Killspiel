<script lang="ts">
	import type { Game } from "./+page"
	import GameChart from "@components/GameChart.svelte"

	/** @type {import("./$types").PageData} */
	export let data: Game
</script>

<!--<Modal buttonPositive="variant-ghost-error" />-->
<svelte:head>
	<title>Game {data.id} | Killspiel</title>
</svelte:head>

<div class="container mx-auto p-2">
	<h1 class="text-6xl text-center">Spiel {data.id}</h1>
	<hr class="p-2 mt-3" />

	<div class="grid grid-cols-3 gap-4">
		<div class="card p-2 text-xl">
			Teilnehmer:
			<div class="text-5xl text-center">{data.userCount}</div>
		</div>
		<div class="card p-2 text-xl">
			Korrekte Schätzungen:
			<div class="text-5xl text-center">{data.correctCount}</div>
		</div>
		<div class="card p-2 text-xl">
			Trefferrate:
			<div class="text-5xl text-center">
				{data.userCount !== 0 ? ((data.correctCount / data.userCount) * 100).toFixed(2) : 0}%
			</div>
		</div>
		<div class="card p-2 text-xl">
			Richtiges Ergebnis:
			<div class="text-5xl text-center">{data.correct}</div>
		</div>
		<div class="card p-2 text-xl">
			Präzision:
			<div class="text-5xl text-center">{data.precision}</div>
		</div>
		<div class="card p-2 text-xl">
			Erstellt:
			<div class="text-5xl text-center">
				{new Date(Date.parse(data.time)).toLocaleString()}
			</div>
		</div>
	</div>

	<div class="p-2 card mt-2">
		<GameChart verteilung={JSON.parse(data.verteilung)} />
	</div>
</div>
