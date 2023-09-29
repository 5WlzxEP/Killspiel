<script lang="ts">
	import { afterUpdate, SvelteComponentTyped } from "svelte"
	import {
		IconBadges,
		IconBrandTwitch,
		IconHome,
		IconSettings,
		IconUser,
		IconDeviceGamepad2
	} from "@tabler/icons-svelte"
	import { page } from "$app/stores"

	type crumb = {
		path: string
		name: string
		icon: SvelteComponentTyped | undefined
	}

	let crumbs: Array<crumb> = []
	afterUpdate(() => {
		let location = $page.url.pathname
		crumbs = new Array<crumb>()
		let cs = location.split("/").slice(1)
		let base = ""

		for (let crumb of cs) {
			if (crumb === "") {
				continue
			}
			base = `${base}/${crumb}`
			const name = crumb.charAt(0).toUpperCase() + crumb.slice(1)
			let icon: SvelteComponentTyped | undefined = undefined
			switch (name) {
				case "Leaderboard":
					icon = IconBadges
					break
				case "Settings":
					icon = IconSettings
					break
				case "Twitchchat":
					icon = IconBrandTwitch
					break
				case "User":
					icon = IconUser
					break
				case "Game":
					icon = IconDeviceGamepad2
			}

			crumbs.push({
				icon: icon,
				path: base,
				name: name
			})
		}
	})
</script>

<div class="mx-auto container justify-center items-center p-2">
	<ol class="breadcrumb btn">
		<li class="crumb">
			<a class="anchor" title="Home" href="/">
				<IconHome />
			</a>
		</li>
		{#each crumbs as crumb}
			<li class="crumb-separator" aria-hidden="true">&rsaquo;</li>

			{#if crumb.icon}
				<li class="crumb">
					<a class="anchor" title={crumb.name} href={crumb.path}>
						<svelte:component this={crumb.icon} />
					</a>
				</li>
			{:else}
				<li class="crumb"><a class="anchor" href={crumb.path}>{crumb.name}</a></li>
			{/if}
		{/each}
	</ol>
</div>
