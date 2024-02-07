import type { PageLoad } from "../../../../.svelte-kit/types/src/routes/(app)/game/$types"
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
}

export const load: PageLoad = async ({ fetch: fetch, url }) => {
	try {
		const limit = parseInt(url.searchParams.get("limit") ?? "50")
		const page = Math.max(parseInt(url.searchParams.get("page") ?? "0") - 1, 0)
		const url2 = `${BACKEND_URL}/api/game/?offset=${page * limit}&limit=${limit}`
		const res = await fetch(url2)
		return { d: await res.json(), page: page, limit: limit }
	} catch (e) {
		console.error(e)
		throw error(404, "not found")
	}
}
