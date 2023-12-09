import { localStorageStore } from "@skeletonlabs/skeleton"
import type { Writable } from "svelte/store"

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL

export type Server =
	| "br1"
	| "eun1"
	| "euw1"
	| "jp1"
	| "kr"
	| "la1"
	| "la2"
	| "na1"
	| "oc1"
	| "ph2"
	| "ru"
	| "sg2"
	| "th2"
	| "tr1"
	| "tw2"
	| "vn2"

export type LoL = {
	name: string
	tag: string
	profileIcon: number
	kategorie: string
	server: Server
	apiKey: string
}

export type General = {
	intervall: number
}

export type Resp = {
	lol: LoL
	general: General
}

export const _ready: Writable<boolean> = localStorageStore("Riot-Ready", false)

export async function _isReady() {
	const url = `${BACKEND_URL}/api/data/riot/ready`
	const res = await fetch(url)
	_ready.set((await res.text()) === "true")
}
