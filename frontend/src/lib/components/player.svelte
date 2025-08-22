<script lang="ts">
  import * as Card from "$lib/components/ui/card/index.js";
  import { command, roomstate, token } from "$lib/room/room.svelte";
  import Badge from "./ui/badge/badge.svelte";
  import Button from "./ui/button/button.svelte";
  let { player, gameplayer, self, playerIndex } = $props();
</script>

<Card.Root class="w-[233px] {self ? 'border-green-900 border-2' : ''}">
  <Card.Header>
    <Card.Title class="text-center"
      >{player.name} {self ? "(you)" : ""}</Card.Title
    >
  </Card.Header>
  {#if gameplayer}
    <Card.Content>
      <div class="flex flex-wrap gap-2 mb-4">
        {#if roomstate.room.game.current_player === playerIndex}
          <Badge variant="outline">Leader</Badge>
        {/if}
        {#if gameplayer.is_in_composition}
          <Badge variant="secondary">In Team</Badge>
        {/if}
        {#if gameplayer.role_card.side === "red"}
          <Badge variant="destructive">Red</Badge>
        {/if}
        {#if gameplayer.role_card.role === "assasin"}
          <Badge variant="destructive">Assasin</Badge>
        {/if}
      </div>
      {#if roomstate.room.game.player_order[roomstate.room.game.current_player] === token.id}
        {#if !gameplayer.is_in_composition}
          <Button
            variant="secondary"
            onclick={() => {
              command("add", player.id);
            }}>Add</Button
          >
        {:else}
          <Button
            variant="secondary"
            onclick={() => {
              command("add", player.id);
            }}>Remove</Button
          >
        {/if}
      {/if}
    </Card.Content>
  {/if}
</Card.Root>
