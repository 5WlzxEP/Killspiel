import type { PageLoad } from "../../../../../../.svelte-kit/types/src/routes/(app)/game/[id]/[vote]/$types"
import { error } from "@sveltejs/kit"

export const prerender = false

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

export type User = {
	name: string
	id: number
	vote: number
}

export const load: PageLoad = async ({ fetch, params }) => {
	try {
		const url = `${BACKEND_URL}/api/game/${params.id}/${params.vote}`
		const res = await fetch(url)
		return { data: await res.json(), params: params }
	} catch (e) {
		console.error(e)

		throw error(404, "not found")
	}
}
