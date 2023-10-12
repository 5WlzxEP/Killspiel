<script lang="ts">
	import type { Game } from "./+page"
	import { goto } from "$app/navigation"

	export let data: { d: Array<Game> }
</script>

<div class="container mx-auto">
	<div class="table-container w-full m-1">
		<table class="table table-hover table-cell-fit mx-auto w-full m-1">
			<thead>
				<tr>
					<th>Rank</th>
					<th class="text-right">Richtiges Ergebnis</th>
					<th class="text-right">Richtige Schätzungen</th>
					<th class="text-right">Teilnehmer</th>
					<th class="text-right">% Richtige Schätzungen</th>
					<th class="text-right">Präzision</th>
					<th class="text-right">Zeit</th>
				</tr>
			</thead>
			<tbody>
				{#each data.d as row}
					<tr class="cursor-pointer" on:click={() => goto(row.id.toString())}>
						<td>{row.id}</td>
						<td class="text-right">{row.correct}</td>
						<td class="text-right">{row.correctCount}</td>
						<td class="text-right">{row.userCount}</td>
						<td class="text-right"
							>{((row.userCount > 0 ? row.correctCount / row.userCount : 0) * 100).toFixed(2)}%</td
						>
						<td class="text-right">{row.precision.toFixed(2)}</td>
						<td class="text-right">
							{new Date(row.time).toLocaleString()}
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
