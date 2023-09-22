<script lang="ts">
	import { getModalStore, type ModalSettings } from "@skeletonlabs/skeleton"
	import { IconHelp } from "@tabler/icons-svelte"

	let ModalSett: ModalSettings

	const modalStore = getModalStore()

	export let value: string
	export let placeholder = ""
	export let label: string
	export let modal:
		| {
				title: string
				body: string
		  }
		| undefined = undefined

	$: {
		if (modal) {
			ModalSett = {
				type: "alert",
				title: modal.title,
				body: modal.body,
				buttonTextCancel: "Schließen"
			}
		}
	}

	function trigger() {
		modalStore.trigger(ModalSett)
	}
</script>

<label class="label">
	<span>
		{label}
		{#if modal}
			<div on:keypress={trigger} on:click={trigger} class="inline-block p-1 cursor-pointer">
				<IconHelp class="inline-block" />
			</div>
		{/if}
	</span>

	<textarea class="textarea" rows="4" bind:value {placeholder} />
</label>
