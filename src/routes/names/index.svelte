<script lang="ts">
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";

  let query = "";
  let loading = false;
  let results: string[] = [];

  async function search() {
    loading = true;
    let res = await fetch(url + "/search_names/" + encodeURIComponent(query + "%"));
    results = await res.json();
    loading = false;
  }

  search();
</script>

<h1>Names</h1>
<p class="lead">
  Get statistics about your name! The data is from how many people have registered for Social Security since 78 years ago. Learn more <a href="https://www.ssa.gov/oact/babynames/limits.html">here</a>!
</p>

<form class="form input-group">
  <input type="search" class="form-control" placeholder="Search..." bind:value={query}/>
  <button class="btn btn-outline-secondary" on:click|preventDefault={search}>
    <i class="bi bi-search"></i>
  </button>
</form>

{#if loading}
<Spinner></Spinner>
{:else}
<ul class="list-group mt-3">
  {#each results as name}
    <a class="list-group-item list-group-item-action" href={"/names/" + name}>{name}</a>
  {/each}
</ul>
{/if}

