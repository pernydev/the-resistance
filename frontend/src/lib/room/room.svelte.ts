import { decodeJWT } from "$lib/index.svelte"

export const ws = new WebSocket(`${import.meta.env.VITE_API_URL}/ws`)
ws.onopen = () => {
  console.log("open")
  ws.send(localStorage.getItem("token")!!)
}

ws.onmessage = (event) => {
  roomstate.room = JSON.parse(event.data);
  console.log(roomstate)
}

export const roomstate: any = $state({
  room: {}
});

export async function command(command: string, data: any) {
  ws.send(JSON.stringify({
    command,
    data,
  }))
}


export const token = decodeJWT(localStorage.getItem("token")!!);
