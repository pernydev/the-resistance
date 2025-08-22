<script lang="ts">
  import { roomstate, token } from "./room.svelte";

  let roleCard = roomstate.room.game.players[token.id].role_card;

  const explanations = {
    "blue:normal":
      "Your goal is to help your fellow blue team members succeed in missions.",
    "blue:commander":
      "You know the members of the red team, your goal is to help your team while not revealing your identity to the Assasin.",
    "blue:bodyguard":
      "You know the Commander, who knows the members of the red team. Your goal is to protect the Commander from being caught by the Assasin.",
    "red:normal":
      "Your goal is to, together with the red team, sabotage and fail missions.",
    "red:assasin":
      "Your goal is to, togetrher with the red team, sabotage and fail missions. After the game, if the blue team wins, you need to shoot one of them who you suspect is the Assasin. The Assasin knows all red members.",
  };
</script>

<h1 class="text-center">
  You are on the <b>{roleCard.side}</b> team, and are a <b>{roleCard.role}</b> player.
</h1>

<p class="max-w-[300px] mx-auto mt-8 text-center">
  {explanations[roleCard.side + ":" + roleCard.role]}
</p>

{#if roleCard.side === "red"}
  <div class="flex items-center flex-col mt-8">
    <span class="mb-2 font-bold text-lg">Red team members:</span>
    {#each Object.values(roomstate.room.game.players) as player}
      {#if player.role_card.side === "red"}
        {roomstate.room.players[player.id].name}
      {/if}
    {/each}
  </div>
{/if}
