<script lang="ts">
  import Player from "$lib/components/player.svelte";
  import Button from "$lib/components/ui/button/button.svelte";
  import Missions from "./Missions.svelte";
  import RoleReveal from "./RoleReveal.svelte";
  import { command, roomstate, token } from "./room.svelte";
</script>

<div class="pt-4">
  {#if !roomstate.room.game}
    <p class="w-fit mx-auto mb-6">Waiting for game...</p>
  {/if}

  {#if roomstate.room.game}
    <Missions />
  {/if}

  {#if roomstate.room.game?.state === "voting"}
    <div class="justify-center flex gap-2 mb-4">
      <Button
        onclick={() => {
          command("vote", "approve");
        }}
      >
        Approve
      </Button>
      <Button
        onclick={() => {
          command("vote", "reject");
        }}
      >
        Reject
      </Button>
    </div>
  {/if}

  {#if roomstate.room.game?.state === "role_reveal"}
    <RoleReveal />
  {:else}
    <a-players class="w-[933px] flex flex-wrap mx-auto gap-2">
      {#if !roomstate.room.game}
        {#each Object.values(roomstate.room.players || {}) as player}
          <Player {player} self={player.id === token.id} />
        {/each}
      {:else}
        {#each roomstate.room.game.player_order as playerID, i}
          <Player
            player={roomstate.room.players[playerID]}
            self={playerID === token.id}
            gameplayer={roomstate.room.game.players[playerID]}
            playerIndex={i}
          />
        {/each}
      {/if}
    </a-players>
  {/if}

  <div class="flex w-[700px] mx-auto justify-center mt-12 gap-4">
    {#if token.id === roomstate.room?.host_id}
      {#if !roomstate.room.game}
        <Button
          onclick={() => {
            command("start", null);
          }}
        >
          Start
        </Button>
      {:else}
        <Button
          onclick={() => {
            command("continue", null);
          }}
        >
          Continue
        </Button>
        <Button
          onclick={() => {
            command("next", null);
          }}
        >
          Next Mission
        </Button>
      {/if}
    {/if}
  </div>
</div>
