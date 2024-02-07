<script lang="ts">
	import { latestVotes } from "@stores/vote"
</script>

<div>
	<h1 class="text-center">Aktuelle Votes</h1>
	<hr />
	<div class="h-60 overflow-y-hidden w-full">
		{#each $latestVotes.d.slice(Math.max($latestVotes.d.length - 9, 0)).toReversed() as e}
			<div class="x">
				{e.time.toLocaleTimeString()}: <span style="color: {e.color}">{e.name}</span> voted {e.vote.toFixed(
					2
				)}
			</div>
		{/each}
	</div>
</div>

<style>
	.x:hover {
		background-color: rgba(100, 100, 100, 50%);
		border-radius: var(--theme-rounded-container);
	}
</style>
