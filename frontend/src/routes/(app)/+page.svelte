<script lang="ts">
	import Leaderboard from "./Leaderboard.svelte"
	import Games from "./Games.svelte"
	import Game from "./Game.svelte"
	import Controls from "./Controls.svelte"
</script>

<svelte:head>
	<title>Killspiel</title>
</svelte:head>

<div class="container mx-auto mt-2 grid grid-cols-2 gap-3">
	<div>
		<Controls />
	</div>
	<div>
		<Game />
	</div>
	<div>
		<h1 class="text-lg text-center anchor"><a href="/leaderboard">Leaderboard</a></h1>
		<Leaderboard />
	</div>
	<div>
		<h1 class="text-lg text-center anchor"><a href="/game">Letzte Spiele</a></h1>
		<Games />
	</div>
</div>
