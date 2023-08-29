<script lang="ts">
    import {getModalStore, type ModalSettings} from "@skeletonlabs/skeleton";
    import {IconHelp} from "@tabler/icons-svelte";
    import {onMount} from "svelte";

    const modalStore = getModalStore()

    let ModalSett: ModalSettings

    export let value: string
    export let prefix = ''
    export let required = false
    export let placeholder = ""
    export let label: string
    export let modal: {
        title: string,
        body: string
    }

    if (modal) {
        ModalSett = {
            type: 'alert',
            title: modal.title,
            body: modal.body,
            buttonTextCancel: "Schließen",
        }
    }

    let div: HTMLDivElement

    onMount(() => {
        div.classList.add(prefix ? "grid-cols-[auto_1fr_auto]" : "grid-cols-[1fr_auto]")
    })
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
    <div bind:this={div} class="input-group input-group-divider">
        {#if prefix}
            <div class="input-group-shim">{prefix}</div>
        {/if}
        <input type="text" bind:value={value} tabindex="0"
               placeholder="{placeholder}"
               {required}
        />
        {#if required}
            <div class="input-group-shim">*</div>
        {/if}
    </div>
</label>