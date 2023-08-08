<script lang="ts">
    import {Paginator} from "@skeletonlabs/skeleton";
    import type {PaginationSettings} from "@skeletonlabs/skeleton";
    import {onMount} from "svelte";

    import { Toast, toastStore } from '@skeletonlabs/skeleton';
    import type { ToastSettings } from '@skeletonlabs/skeleton';

    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

    export let offset = 0
    export let limit = 25

    let meta: PaginationSettings = {
        offset: offset,
        limit: limit,
        size: 0,
        amounts: [10, 25, 50, 100],
    }

    type Leaderboard = {
        rank: number
        name: string
        points: number
        guesses: number
    }

    type result = {
        meta: PaginationSettings
        data: Array<Leaderboard>
    }

    async function fetchLeaderboard(): Promise<Array<Leaderboard>> {
        const url = `${BACKEND_URL}/api/leaderboard?o=${meta.offset * meta.limit}&l=${meta.limit}`
        try {
            const resp = await fetch(url);
            const ob: result = await resp.json()
            meta.size = ob.meta.size
            return ob.data
        } catch (e) {
            const t: ToastSettings = {
                message: 'Es ist ein Fehler beim Abfragen der Daten aufgetreten.',
                timeout: 5000,
                background: 'variant-filled-error',
            };
            toastStore.trigger(t)
        }
        return []
    }

    let data: Array<Leaderboard> = []

    async function update() {
        data = []
        data = await fetchLeaderboard()
    }

    onMount(async () => {
        await update()
    })
</script>

<Toast />
<div class="table-container mx-auto w-full m-1">
    <!-- Native Table Element -->
    <table class="table table-hover table-cell-fit mx-auto w-full m-1">
        <thead>
        <tr>
            <th class="table-sort-asc p-1 m-1">Rank </th>
            <th class="text-center">Name</th>
            <th class="text-center">Punkte</th>
            <th class="text-center">Teilnahmen</th>
            <th class="text-center">Rate</th>
        </tr>
        </thead>
        <tbody>
        {#each data as row}
            <tr class="">
                <td>{row.rank}</td>
                <td>{row.name}</td>
                <td class="text-right">{row.points}</td>
                <td class="text-right">{row.guesses}</td>
                <td class="text-right">{((row.guesses > 0 ? row.points / row.guesses : 0) * 100).toFixed(2)} %</td>
            </tr>
        {:else }
            <tr>
                <td>
<!--                    <ProgressRadial width="w-11" stroke={100} meter="stroke-primary-500" track="stroke-primary-500/30"/>-->
                    <div class="placeholder animate-pulse" />
                </td>
                <td>
                    <div class="placeholder animate-pulse" />
                </td>
                <td>
                    <div class="placeholder animate-pulse" />
                </td>
                <td>
                    <div class="placeholder animate-pulse" />
                </td>
                <td>
                    <div class="placeholder animate-pulse" />
                </td>
            </tr>
        {/each}
        </tbody>
    </table>
    <div class="mt-2">
                <Paginator
                        bind:settings={meta}
                        showFirstLastButtons="{true}"
                        showPreviousNextButtons="{true}"
                        on:page={update}
                        on:amount={update}
                />
    </div>
</div>
