<script lang="ts">
  import { goto } from "$app/navigation";
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import { url } from "$lib/consts";
  import { passwordStore } from "../password";

  let name = "";
  let password = "";
  let loading = false;

  async function create() {
    loading = true;
    await fetch(url + "/new_note/" + encodeURIComponent(name) + "/" + encodeURIComponent(password));
    passwordStore.set(password);
    goto("/notes/" + encodeURIComponent(name));
    loading = false;
  }
</script>

<h1>New Note</h1>
<form class="text-start">
  <div class="mb-3">
    <label for="name" class="form-label">Note Name</label>
    <input type="text" class="form-control" id="name" placeholder="Note name..." bind:value={name}>
  </div>

  <div>
    <label for="password" class="form-label">Password</label>
    <input type="password" class="form-control" id="password" placeholder="Password..." bind:value={password}>
    <div class="form-text">Leave blank for no password.</div>
  </div>

  <div class="text-center">
    <LoadingButton type="submit" style="btn-success" on:click={create} loading={loading}>Create</LoadingButton>
  </div>
</form>