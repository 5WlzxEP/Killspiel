<script lang="ts">
	import { onDestroy } from "svelte"

	type style = { color: string; name: string; dark: boolean }

	const v: Array<style> = [
		{ color: "15 186 129", name: "skeleton", dark: false },
		{ color: "15 186 129", name: "skeleton", dark: true },
		{ color: "236 72 153", name: "modern", dark: false },
		{ color: "236 72 153", name: "modern", dark: true },
		{ color: "6 182 212", name: "rocket", dark: false },
		{ color: "6 182 212", name: "rocket", dark: true },
		{ color: "134 208 203", name: "seafoam", dark: false },
		{ color: "134 208 203", name: "seafoam", dark: true },
		{ color: "234 134 26", name: "vintage", dark: false },
		{ color: "234 134 26", name: "vintage", dark: true },
		{ color: "236 170 54", name: "sahara", dark: false },
		{ color: "236 170 54", name: "sahara", dark: true },
		{ color: "168 190 241", name: "hamlindigo", dark: false },
		{ color: "168 190 241", name: "hamlindigo", dark: true },
		{ color: "116 74 161", name: "gold-nouveau", dark: false },
		{ color: "230 200 51", name: "gold-nouveau", dark: true },
		{ color: "212 22 60", name: "crimson", dark: false },
		{ color: "212 22 60", name: "crimson", dark: true },
		{ color: "59 130 246", name: "wintry", dark: false },
		{ color: "59 130 246", name: "wintry", dark: true }
	]
	let value: style

	function getRandomV() {
		value = v[Math.floor(Math.random() * v.length)]
	}

	getRandomV()

	let intervall = setInterval(getRandomV, 10000)

	onDestroy(() => clearInterval(intervall))

	function setTheme(theme: style) {
		clearInterval(intervall)

		value = theme
	}
</script>

<div class="w-full p-4 container mx-auto mb-2 [&>*]:cursor-pointer">
	{#each v as val, i}
		{#if i === 0}
			<button style="color: rgb({val.color})" on:click={() => setTheme(val)}>{val.name}</button>
		{:else}
			,
			<button style="color: rgb({val.color})" on:click={() => setTheme(val)}
				>{val.name} {val.dark ? "(dark)" : ""}</button
			>
		{/if}
	{/each}

	<table class="mt-1 card p-2" style="border-spacing: 10px 0; border-collapse: separate;">
		<thead>
			<tr>
				<th colspan="5">Parameter</th>
			</tr>
			<tr>
				<th class="p-2">Name</th>
				<th class="p-2">Bedeutung</th>
				<th class="p-2">Werte</th>
				<th class="p-2">Standardwert</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>theme</td>
				<td>Setzt das Theme.</td>
				<td
					>skeleton, modern, rocket, seafoam, vintage, sahara, hamlindigo, gold-nouveau, crimson,
					wintry</td
				>
				<td><i>skeleton</i></td>
			</tr>
			<tr>
				<td>dark</td>
				<td>Setzt den Dark-mode</td>
				<td><i>true</i> &#8793; Dark-mode, <i>false</i> &#8793; White-mode</td>
				<td><i>true</i></td>
			</tr>
		</tbody>
	</table>
</div>

<div class="grid-cols-1 grid mx-auto container gap-2">
	<div>
		<h1 class="text-2xl text-center">Letzte Spiele</h1>
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
	<div>
		<h1 class="text-2xl text-center">Leaderboard</h1>

		<table class="card p-2" style="border-spacing: 10px 0; border-collapse: separate;">
			<thead>
				<tr>
					<th colspan="5">Parameter</th>
				</tr>
				<tr>
					<th class="p-2">Name</th>
					<th class="p-2">Bedeutung</th>
					<th class="p-2">Werte</th>
					<th class="p-2">Standardwert</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>order</td>
					<td>Bestimmt, ob aufsteigend oder absteigend sortiert wird.</td>
					<td><i>asc</i> &#8793; Aufsteigend, <i>desc</i> &#8793; Absteigend</td>
					<td><i>desc</i></td>
				</tr>
				<tr>
					<td>sorting</td>
					<td>Setzt den Wert, nach dem sortiert wird</td>
					<td
						><i>name</i> &#8793; Name, <i>points</i> &#8793; Punkte, <i>teilnahmen</i> &#8793; Teilnahmen</td
					>
					<td><i>points</i></td>
				</tr>
				<tr>
					<td>limit</td>
					<td>bestimmt die Anzahl an Einträgen</td>
					<td>Nummer von 1-100</td>
					<td><i>10</i></td>
				</tr>
				<tr>
					<td>latest</td>
					<td>Gibt an. ob <i>Letzten 8 Versuche</i> angezeigt wird</td>
					<td>true &#8793; es wird angezeigt, false &#8793; es wird nicht angezeigt</td>
					<td><i>false</i></td>
				</tr>
			</tbody>
		</table>
		<div class="grid grid-cols-2">
			<div>
				<div class="underline">
					{window.location}leaderboard?theme=<x style="color: rgb({value.color})">{value.name}</x
					>&dark={value.dark}&order=asc&sorting=name&limit=10&latest=true
				</div>
				Width: 767, Height: 655
				<iframe
					width="767"
					height="655"
					src="leaderboard?theme={value.name}&dark={value.dark}&order=asc&sorting=name&limit=10&latest=true"
					title="game overlay"
					style="scale: calc(614/655); transform-origin: 0 0;"
				/>
			</div>

			<div>
				<div class="underline">
					{window.location}leaderboard?theme=<x style="color: rgb({value.color})">{value.name}</x
					>&dark={value.dark}
				</div>
				Width: 595, Height: 615
				<iframe
					width="595"
					height="615"
					src="leaderboard?theme={value.name}&dark={value.dark}"
					title="game overlay"
				/>
			</div>
		</div>
	</div>
</div>
