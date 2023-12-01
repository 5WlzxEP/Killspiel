import type {PageLoad} from "../../../../../.svelte-kit/types/src/routes/(app)/user/[id]/$types"
import {error} from "@sveltejs/kit"
import type {SvelteComponent} from "svelte"

export const prerender = false

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

export type User = {
	id: number
	name: string
	guesses: number
	points: number
	latest: number
	history: Array<{
		id: number
		guess: number
		correct: number
		precision: number
		time: string
		icon?: SvelteComponent
	}>
}

export const load: PageLoad = async ({ fetch: fetch, params }) => {
	try {
		const url = `${BACKEND_URL}/api/user/${params.id}/`
		const res = await fetch(url)
		return await res.json()
	} catch (e) {
		console.error(e)

		throw error(404, "not found")
	}
}
