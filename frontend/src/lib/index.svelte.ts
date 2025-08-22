export const s = $state({
  view: new URLSearchParams(location.search).get("room") === "1" ? "room" : "lobby"
})

export function decodeJWT(jwt: string) {
  const parts = jwt.split('.');
  const payload = parts[1];
  const decodedPayload = JSON.parse(atob(payload.replace(/-/g, '+').replace(/_/g, '/')));
  return decodedPayload;
}



