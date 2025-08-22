<script lang="ts">
  import { Button } from "$lib/components/ui/button/index.js";
  import Lobby from "$lib/lobby/Lobby.svelte";
  import Room from "$lib/room/Room.svelte";
  import { s } from "$lib/index.svelte";
</script>

{#if s.view === "lobby"}
  <Lobby />
{:else if s.view === "room"}
  <Room />
{/if}
