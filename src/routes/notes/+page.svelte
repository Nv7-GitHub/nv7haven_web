<script lang="ts">
  import { goto } from "$app/navigation";
  import { url } from "$lib/consts";
  import { passwordStore } from "./password";

  let loading = false;
  let query = "";
  let pages: string[] = [];

  async function search() {
    loading = true;
    let res = await fetch(url + "/note_search/" + encodeURIComponent(query + "%"))
    pages = await res.json();
    loading = false;
  }

  async function open(note: string) {
    let res = await fetch(url + "/has_password/" + encodeURIComponent(note));
    let hasPassword = await res.text() == "1";
    if (hasPassword) {
      let password = prompt("Enter password for " + note);
      if (password == null) {
        return;
      }
      res = await fetch(url + "/get_note/" + encodeURIComponent(note) + "/" + encodeURIComponent(password));
      if (res.status == 200) {
        passwordStore.set(password);
        goto("/notes/" + encodeURIComponent(note));
        return;
      }
    } else {
      passwordStore.set("");
      goto("/notes/" + encodeURIComponent(note));
    }
  }

  search();
</script>

<h1>QwikNotes</h1>
<p class="lead">Have you ever been switching computers or users but wanted to keep some text quickly and easily? Well now you can! QwikNotes is a note-taking app that uses your IP Address to identify you, so that as long as you are on the same WiFi network you keep your notes! You can optionally password-protect your notes in case you don't want people on your wifi network seeing your notes.</p>

<form class="input-group">
  <button class="btn btn-outline-secondary" on:click|preventDefault={() => {goto("/notes/create")}} type="button">
    <i class="bi bi-plus-lg"></i>
  </button>

  <input type="search" class="form-control" placeholder="Search..." bind:value={query}/>
  <button class="btn btn-outline-secondary" type="submit" disabled={loading} on:click|preventDefault={search}>
    {#if loading}
      <span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
    {:else}
      <i class="bi bi-search"></i>
    {/if}
  </button>
</form>

<ul class="list-group mt-3">
  {#each pages as page}
    <button class="list-group-item list-group-item-action" on:click={() => {open(page)}}>{page}</button>
  {/each}
</ul>
