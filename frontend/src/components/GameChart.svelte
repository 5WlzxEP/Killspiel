<script lang="ts">
	import { Bar, getElementAtEvent } from "svelte-chartjs"

	import {
		Chart as ChartJS,
		Title,
		Legend,
		Tooltip,
		BarElement,
		LinearScale,
		CategoryScale
	} from "chart.js"
	import { goto } from "$app/navigation"
	import type { Writable } from "svelte/store"

	ChartJS.register(Title, Legend, Tooltip, BarElement, LinearScale, CategoryScale)
	export let verteilung: Writable<never>
	export let gameid: string | number

	let data: {
		labels: Array<string>
		datasets: Array<{
			type: string
			label: string
			backgroundColor: string
			borderWidth: number
			data: Array<string>
		}>
	}
	$: data = {
		labels: Object.keys($verteilung),
		datasets: [
			{
				type: "bar",
				label: "Verteilung der Votes",
				backgroundColor: color(),
				borderWidth: 1,
				data: Object.values($verteilung)
			}
		]
	}

	let chart: never

	function color() {
		const c = getComputedStyle(document.body)
			.getPropertyValue("--color-primary-500")
			.split(" ")
			.map((v) => toHex(v))
		return "#" + c.join("")
	}

	function toHex(i: string): string {
		let str = Number(i).toString(16)
		return str.length == 1 ? "0" + str : str
	}

	function onClick(e: PointerEvent) {
		goto(`/game/${gameid}/` + data.labels[getElementAtEvent(chart, e)[0].index].toString())
	}
</script>

<div class="text-primary-500">
	<Bar {data} bind:chart options={{ responsive: true }} on:click={onClick} />
</div>
