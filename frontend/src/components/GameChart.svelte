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

	ChartJS.register(Title, Legend, Tooltip, BarElement, LinearScale, CategoryScale)
	export let verteilung

	let data = {
		labels: Object.keys(verteilung),
		datasets: [
			{
				type: "bar",
				label: "Verteilung der Votes",
				borderColor: "#5B0888",
				backgroundColor: "#713ABE",
				borderWidth: 1,
				data: Object.values(verteilung)
			}
		]
	}

	let chart

	function onClick(e: PointerEvent) {
		goto(data.labels[getElementAtEvent(chart, e)[0].index].toString())
	}
</script>

<div class="text-primary-500">
	<Bar {data} bind:chart options={{ responsive: true }} on:click={onClick} />
</div>
