<script lang="ts">
	import GameChart from "@components/GameChart.svelte"
	import { onMount } from "svelte"
	import type { Game } from "./game/[id]/+page"
	import { state } from "@stores/state"
	import { writable } from "svelte/store"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	let data: Game

	const update = async () => {
		const url = `${BACKEND_URL}/api/game/latest/`
		const res = await fetch(url)
		data = await res.json()
		$verteilung = JSON.parse(data.verteilung)
	}

	onMount(update)

	$: {
		if ($state === 2) setTimeout(update, 1000)
	}

	let verteilung = writable()

	$: {
		if (data?.verteilung) $verteilung = JSON.parse(data.verteilung)
	}
</script>

{#if data?.verteilung}
	<h1 class="text-lg text-center anchor">
		<a href="/game/{data.id}">Aktuelles Spiel: {data.id}</a>
	</h1>
	<div class="card p-2">
		<GameChart bind:verteilung gameid={data.id} />
	</div>
{/if}
