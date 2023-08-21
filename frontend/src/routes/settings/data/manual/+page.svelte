<script lang="ts">
    import {type ToastSettings, toastStore, Toast} from '@skeletonlabs/skeleton'
    import {onMount} from "svelte";
    import InputText from "@components/InputText.svelte";
    import InputArea from "@components/InputArea.svelte";
    import {writable} from "svelte/store";

    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

    let result = 0
    let state = 0
    let start: HTMLButtonElement
    let dissolve: HTMLButtonElement

    function change(e: MouseEvent & { currentTarget: (EventTarget & HTMLButtonElement) }) {
        state = ++state % 3
        console.log(state)
        start.disabled = (state != 0) // none
        dissolve.disabled = (state != 2) || (result == undefined) // running


        if (state === 1) {
            setTimeout(change, 5000)
        }
    }

    async function submit(e: Event) {
        console.log(e)
        try {
            const res = await fetch(`${BACKEND_URL}/api/collector/chat/`, {
                headers: {
                    'Content-Type': 'application/json'
                },
                method: "POST",
                body: JSON.stringify("")
            })
            if (res.ok) {
                const t: ToastSettings = {
                    message: 'Einstellungen erfolgreich gespeichert.',
                    timeout: 5000,
                    background: 'variant-filled-success',
                }
                toastStore.trigger(t)
            }
        } catch (e) {
            console.error(e)
            const t: ToastSettings = {
                message: 'Es ist ein Fehler beim Abfragen der aktualisieren Einstellungen aufgetreten.',
                timeout: 5000,
                background: 'variant-filled-error',
            }
            toastStore.trigger(t)
        }
    }

    onMount(async () => {
        const url = `${BACKEND_URL}/api/data/manual/`
        try {
            const res = await fetch(url)
            const val = await res.json()
            state = val.state
        } catch (e) {
            const t: ToastSettings = {
                message: 'Es ist ein Fehler beim Abfragen der aktuellen Einstellungen aufgetreten.',
                timeout: 5000,
                background: 'variant-filled-error',
            }
            toastStore.trigger(t)
        }
        start.disabled = (state != 0) // none
         // running

        console.log(result)
    })

</script>

<div class="container mx-auto">
    <div class="card w-full">
        <h1 class="text-center text-2xl p-3">Manuell</h1>
        <hr>
        <div class="grid gap-10 w-full m-2 grid-cols-[7fr_1fr_3fr_1fr_7fr] p-2">
            <div class="p-2 grid gap-2 h-3">
                <button class="btn variant-ghost" type="button" bind:this={start} on:click={e => change(e)}>
                    Starten
                </button>
            </div>

            <span class="divider-vertical h-full"/>

            <div class="p-2 grid gap-2 h-3">
                Current State:
                {#if state === 0}
                    Ready
                {:else if (state === 1)}
                    Sammele Schätzungen
                {:else if (state === 2)}
                    Warte auf Auflösung
                {/if}
            </div>

            <span class="divider-vertical h-full"/>

            <form class="p-2 gap-2 grid">
                <input class="input" bind:value={result} on:change={() => dissolve.disabled = (state !== 2) || (result === undefined)} placeholder="7.2" type="number" required>
                <button class="btn variant-ghost" bind:this={dissolve} on:click={e => change(e)} type="submit">
                    Ergebnis speichern
                </button>
            </form>
        </div>
    </div>
</div>