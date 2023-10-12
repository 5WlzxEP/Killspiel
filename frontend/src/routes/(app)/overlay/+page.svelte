<script lang="ts">
	import { onDestroy } from "svelte"

	type style = { color: string; name: string; dark: number }

	const v = [
		{ color: "212 22 60", name: "crimson", dark: 0 },
		{ color: "212 22 60", name: "crimson", dark: 1 },
		{ color: "116 74 161", name: "gold-nouveau", dark: 0 },
		{ color: "230 200 51", name: "gold-nouveau", dark: 1 },
		{ color: "168 190 241", name: "hamlindigo", dark: 0 },
		{ color: "168 190 241", name: "hamlindigo", dark: 1 },
		{ color: "236 72 153", name: "modern", dark: 0 },
		{ color: "236 72 153", name: "modern", dark: 1 },
		{ color: "6 182 212", name: "rocket", dark: 0 },
		{ color: "6 182 212", name: "rocket", dark: 1 },
		{ color: "236 170 54", name: "sahara", dark: 0 },
		{ color: "236 170 54", name: "sahara", dark: 1 },
		{ color: "134 208 203", name: "seafoam", dark: 0 },
		{ color: "134 208 203", name: "seafoam", dark: 1 },
		{ color: "15 186 129", name: "skeleton", dark: 0 },
		{ color: "15 186 129", name: "skeleton", dark: 1 },
		{ color: "234 134 26", name: "vintage", dark: 0 },
		{ color: "234 134 26", name: "vintage", dark: 1 },
		{ color: "59 130 246", name: "wintry", dark: 0 },
		{ color: "59 130 246", name: "wintry", dark: 1 }
	]
	let value: style

	function getRandomV() {
		value = v[Math.floor(Math.random() * v.length)]
	}

	getRandomV()

	let intervall = setInterval(getRandomV, 10000)

	onDestroy(() => clearInterval(intervall))
</script>

<div class="card w-full p-2 container mx-auto text-center mb-2">
	Themes:
	{#each v.filter((v) => v.dark === 1) as val, i}
		{#if i === 0}
			<x style="color: rgb({val.color})">{val.name}</x>
		{:else}
			,
			<x style="color: rgb({val.color})">{val.name}</x>
		{/if}
	{/each}
</div>

<div class="grid mx-auto container">
	<div>
		<div class="underline">
			{window.location}games?theme=<x style="color: rgb({value.color})">{value.name}</x
			>&dark={value.dark}
		</div>
		Width: 376, Height: 614
		<iframe
			width="376"
			height="614"
			src="games?theme={value.name}&dark={value.dark}"
			title="game overlay"
		/>
	</div>
</div>

<style>
</style>
