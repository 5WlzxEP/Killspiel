<script lang="ts">
	import { goto } from "$app/navigation"
	import { onMount } from "svelte"
	import { state } from "@stores/state"

	const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

	type Game = {
		id: number
		correct: number
		time: string
		correctCount: number
		userCount: number
		precision: number
	}

	let data: Promise<Array<Game>> = new Promise((_, r) => setTimeout(r, 2000))

	const update = async () => {
		const url = `${BACKEND_URL}/api/game/?limit=10`
		const res = await fetch(url)
		data = res.json()
	}

	onMount(update)

	$: {
		if ($state === 0) update()
	}
</script>

<div class="container mx-auto">
	<div class="table-container w-full">
		<table class="table table-hover table-cell-fit mx-auto w-full m-1">
			<thead>
				<tr>
					<th>Id</th>
					<th class="text-right">Ergebnis</th>
					<th class="text-right">Anzahl Gewinner</th>
					<th class="text-right">Teilnehmer</th>
					<th class="text-right">% ri. Schätzungen</th>
					<th class="text-right">Präzision</th>
				</tr>
			</thead>
			<tbody>
				{#await data}
					<tr>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
						<td>
							<div class="placeholder animate-pulse" />
						</td>
					</tr>
				{:then d}
					{#each d as row}
						<tr class="cursor-pointer" on:click={() => goto("game/" + row.id.toString())}>
							<td>{row.id}</td>
							<td class="text-right">{row.correct}</td>
							<td class="text-right">{row.correctCount}</td>
							<td class="text-right">{row.userCount}</td>
							<td class="text-right"
								>{((row.userCount > 0 ? row.correctCount / row.userCount : 0) * 100).toFixed(2)}%
							</td>
							<td class="text-right">{row.precision.toFixed(2)}</td>
						</tr>
					{/each}
				{:catch error}
					<tr>
						<td class="text-center" colspan="10"
							>Keine Daten oder keine Verbindung
							{error ? ": " + error : ""}
						</td>
					</tr>
				{/await}
			</tbody>
		</table>
	</div>
</div>
