<script lang="ts">
import { goto } from "$app/navigation";

  import { page } from "$app/stores";
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";
  import { passwordStore } from "./password";

  let name = decodeURIComponent($page.params["name"]);
  let text = "";
  let original = "";
  let loading = false;

  let saving = false;
  let deleting = false;

  async function load() {
    loading = true;
    let res = await fetch(url + "/get_note/" + encodeURIComponent(name) + "/" + encodeURIComponent($passwordStore));
    text = await res.text();
    original = text;
    loading = false;
  }

  load();

  async function save() {
    saving = true;
    await fetch(url + "/change_note/" + encodeURIComponent(name), {
      method: "POST",
      body: text
    });
    original = text;
    saving = false;
  }

  async function del() {
    deleting = true;
    await fetch(url + "/delete_note/" + encodeURIComponent(name) + "/" + encodeURIComponent($passwordStore));
    goto("/notes");
    passwordStore.set("");
    deleting = false;
  }
</script>

<h1>{name}</h1>

{#if loading}
<Spinner></Spinner>
{:else}
<textarea bind:value={text} class="form-control"></textarea>

<div class="btn-group mt-3">
  <LoadingButton style="btn-success" type="submit" disabled={original == text} loading={saving} on:click={save}>Save</LoadingButton>
  <LoadingButton style="btn-danger" type="button" loading={deleting} on:click={del}>Delete</LoadingButton>
</div>
{/if}