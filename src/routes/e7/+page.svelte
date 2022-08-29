<script lang="ts">
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";
  import { onMount } from "svelte";

  let loading = true;
  let ascending = false;
  let page = 0;
  enum Mode {
    Player = "player",
    Color = "color",
  }
  let mode = Mode.Player;

  type Item = {
    Title: string,
    Value: number,
  }
  type Result = {
    PageLength: number,
    Items: Item[],
  }
  let result: Result;

  async function load() {
    let res = await fetch(url + "/ldb_query/" + (ascending ? "1" : "0" + "/") + mode + "/" + page);
    result = await res.json();
    loading = false;
  }

  onMount(load);
</script>

<h1>Nv7's Elemental Leaderboard</h1>
<p class="lead mb-3">A leaderboard for element count in <a href="https://elem7.nv7haven.com">Nv7's Elemental</a></p>

<ul class="nav nav-pills float-start mb-1">
  <li class="nav-item">
    <span class="nav-link disabled">Sort By</span>
  </li>
  <li class="nav-item">
    <button class="nav-link" class:active={!ascending} on:click={() => {ascending = false; load()}}>Descending</button>
  </li>
  <li class="nav-item">
    <button class="nav-link" class:active={ascending} on:click={() => {ascending = true; load()}}>Ascending</button>
  </li>
</ul>

<ul class="nav nav-pills float-end mb-1">
  <li class="nav-item">
    <span class="nav-link disabled">Leaderboard Type</span>
  </li>
  <li class="nav-item">
    <button class="nav-link" class:active={mode == Mode.Player} on:click={() => {mode = Mode.Player; load()}}>Player</button>
  </li>
  <li class="nav-item">
    <button class="nav-link" class:active={mode == Mode.Color} on:click={() => {mode = Mode.Color; load()}}>Color</button>
  </li>
</ul>

{#if loading}
  <Spinner></Spinner>
{:else}
  <table class="table mt5">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Name</th>
        <th scope="col">Score</th>
      </tr>
    </thead>
    <tbody>
      {#each result.Items as item, i}
        <tr>
          <th scope="row">{page*result.PageLength+i+1}</th>
          <td>{item.Title}</td>
          <td>{item.Value}</td>
        </tr>
      {/each}
    </tbody>
  </table>

  <div class="btn-group mb-5">
    <button class="btn btn-secondary" on:click={() => {page--; load()}} disabled={page == 0}><i class="bi bi-arrow-left"></i></button>
    <button class="btn btn-secondary" on:click={() => {page++; load()}} disabled={result.Items.length < result.PageLength}><i class="bi bi-arrow-right"></i></button>
  </div>
{/if}