import type { PageLoad } from "../../../../../.svelte-kit/types/src/routes/(app)/game/[id]/$types"
import { error } from "@sveltejs/kit"

export const prerender = false

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

export type Game = {
	id: number
	correct: number
	time: string
	correctCount: number
	userCount: number
	precision: number
	verteilung: string
}

export const load: PageLoad = async ({ fetch: fetch, params }) => {
	try {
		const url = `${BACKEND_URL}/api/game/${params.id}/`
		const res = await fetch(url)
		return await res.json()
	} catch (e) {
		console.error(e)

		throw error(404, "not found")
	}
}
