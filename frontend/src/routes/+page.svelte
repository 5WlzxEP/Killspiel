<script lang="ts">
	import {Step, Stepper} from "@skeletonlabs/skeleton"

	function next(e: {detail: {state: {current: number, total: number}, step: number}}): void {
		console.log(e)
		e.detail.state.current = 3

	}
</script>

<svelte:head>
	<title>Killspiel</title>
</svelte:head>

<div class="container h-full mx-auto m-2 justify-center items-center">
	<div class="card p-3">
		<Stepper on:next={next}>
			<Step></Step>
			<Step>
				<p>(content)</p>
			</Step>

			<Step >asfsfad</Step>
			<Step>afgdsgdsdgsgfsddfghs</Step>
		</Stepper>
	</div>
</div>
