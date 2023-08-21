<script>
    import {IconBrandTwitch} from "@tabler/icons-svelte";
</script>

<div class="container grid grid-cols-2 mx-auto gap-4">

    {#each {length: 5} as _, i}
        <a href="/settings/collector/twitchchat" class="card p-2">
           <div class="flex gap-2">
               <IconBrandTwitch />
            <h1>Twitchchat</h1>
           </div>
            Lass die Teilnehmer über den Twitchchat abstimmen.
        </a>
    {/each}
</div>
