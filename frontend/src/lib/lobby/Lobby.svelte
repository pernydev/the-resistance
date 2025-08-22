<script>
  import Button from "$lib/components/ui/button/button.svelte";
  import Input from "$lib/components/ui/input/input.svelte";
  import { s } from "$lib/index.svelte";

  let username = $state("");
  let code = $state("");

  async function submit() {
    const url = `${import.meta.env.VITE_API_URL}/rooms/${code === "" ? "new" : code}`;
    console.log(url);
    const resp = await fetch(url, {
      method: "POST",
      body: JSON.stringify({
        name: username,
      }),
    });

    const data = await resp.json();
    if (code === "") {
      alert("Room created, code is " + data.room_id);
    }

    localStorage.setItem("token", data.token);
    location.href = "/?room=1";
  }
</script>

<div class="h-screen flex justify-center items-center flex-col gap-6">
  <h1 class="text-3xl">The Resistance</h1>
  <Input placeholder="Username" bind:value={username} class="w-128" />
  <Input
    placeholder="Room code (Leave blank to create new)"
    bind:value={code}
    class="w-128"
  />
  <Button onclick={submit}>Join</Button>
</div>
