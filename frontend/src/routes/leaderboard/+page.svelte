<script lang="ts">
    import type {PaginationSettings} from "@skeletonlabs/skeleton"
    const BACKEND_URL = import.meta.env.VITE_BACKEND_URL
    import Leaderboard from "@components/Leaderboard.svelte"

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

    export async function fetchLeaderboard(offset = 0, limit = 25) : Promise<result>  {
        const url= `${BACKEND_URL}/api/leaderboard`

        const resp = await fetch(url);
        const ob : result = await resp.json()
        ob.meta.amounts = [10, 25, 50, 100]
        ob.meta.offset = offset
        ob.meta.limit = limit
        return ob
    }

</script>

<svelte:head>
    <title>Leaderboard | Killspiel </title>
</svelte:head>
<div class="container mx-auto">
    <Leaderboard />
</div>