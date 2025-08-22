<script lang="ts">
  import { roomstate } from "./room.svelte";
</script>

<div class="flex justify-center gap-4 mb-8">
  {#each Object.values(roomstate.room.game.missions || {}) as mission}
    <div
      class="bg-secondary p-4 px-8 flex flex-col items-center {mission.fails_required !==
      1
        ? 'border-2 border-green-500'
        : ''} {roomstate.room.game.round === mission.index
        ? 'bg-white text-primary-foreground'
        : ''}"
    >
      <div>
        {mission.participant_count}
      </div>
      <div>
        {#if mission.result}
          {mission.result}
        {:else}
          ------
        {/if}
      </div>
    </div>
  {/each}
</div>
