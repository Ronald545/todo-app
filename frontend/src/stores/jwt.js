import { writable, readable } from "svelte/store";

const stored = localStorage.getItem('authkey')

export const jwt = writable(stored),jwt_read = readable(stored)

jwt.subscribe((value) => localStorage.setItem('authKey', value))
