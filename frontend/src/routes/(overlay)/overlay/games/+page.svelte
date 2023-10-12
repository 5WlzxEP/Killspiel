<script lang="ts">
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

		// forced update every 5 minutes since last update
		clearTimeout(timeout)
		timeout = setTimeout(update, 5 * 60 * 1000)
	}

	$: {
		if ($state === 0 || $state === 2) update()
	}

	// eslint-disable-next-line no-undef
	let timeout: NodeJS.Timeout
</script>

<h1 class="text-primary-500 text-center text-lg card rounded-none">Letzte Spiele</h1>
<div class="table-container w-full rounded-none">
	<table class="table table-hover table-cell-fit mx-auto w-full rounded-none overflow-x-hidden">
		<thead>
			<tr class="overflow-x-hidden">
				<th class="text-right">Erg.</th>
				<th class="text-right">Gew.</th>
				<th class="text-right">Teiln.</th>
				<th class="text-right">% ri. Schä.</th>
				<th class="text-right">Prä.</th>
			</tr>
		</thead>
		<tbody>
			{#await data}
				{#each { length: 10 } as _}
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
					</tr>
				{/each}
			{:then d}
				{#each d as row}
					<tr>
						<td class="text-right text-xl">{row.correct}</td>
						<td class="text-right text-xl">{row.correctCount}</td>
						<td class="text-right text-xl">{row.userCount}</td>
						<td class="text-right text-xl"
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
