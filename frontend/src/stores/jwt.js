import { writable } from "svelte/store";

const stored = localStorage.getItem('authkey')

export const jwt = writable(stored)

jwt.subscribe((value) => localStorage.authKey = value)
