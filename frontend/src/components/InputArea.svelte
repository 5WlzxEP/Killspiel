<script lang="ts">
    import { getModalStore, type ModalSettings} from "@skeletonlabs/skeleton";
    import {IconHelp} from "@tabler/icons-svelte";

    let ModalSett: ModalSettings

    const modalStore = getModalStore()

    export let value: string
    export let placeholder = ""
    export let label: string
    export let modal: {
        title: string,
        body: string
    } | undefined = undefined

    if (modal) {
        ModalSett = {
            type: 'alert',
            title: modal.title,
            body: modal.body,
            buttonTextCancel: "Schließen",
        }
    }
</script>

<label class="label">
    <span>{label}
        {#if modal}
            <button on:click={() => {modalStore.trigger(ModalSett)}} class="inline-block p-1"
                    type="button">
            <IconHelp class="inline-block"/>
            </button>
        {/if}
    </span>

        <textarea class="textarea" rows="4" bind:value={value} placeholder="{placeholder}" />
</label>