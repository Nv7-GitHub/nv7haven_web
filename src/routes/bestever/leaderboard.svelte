<script lang="ts">
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";
  import { onMount, onDestroy } from "svelte";
  import { count } from "./data";

  let ready = false;
  let items: string[] = [];
  let loading = false;

  async function loadLdb() {
    loading = true;
    let res = await fetch(url + "/bestever_get_ldb/" + $count.toString());
    items = await res.json();
    loading = false;
  }

  let ev: EventSource;

  async function getReady() {
    await loadLdb();
    ev = new EventSource("https://nv7haven.firebaseio.com/data.json");
    ev.addEventListener("put", () => {
      loadLdb()
    })
    ev.onopen = () => {
      ready = true;
    }
    onDestroy(ev.close);
  }

  onMount(getReady);
</script>

<p class="lead">These are the best things ever!</p>

{#if ready}
<ul class="list-group">
  {#each items as item}
    <li class="list-group-item">{item}</li>
  {/each}
</ul>
{:else}
<Spinner></Spinner>
{/if}

<p class="lead mt-3">Want to see more?</p>
<form class="input-group mb-5">
  <input type="number" bind:value={$count} class="form-control"/>
  <LoadingButton style="btn-primary" type="submit" loading={loading} on:click={loadLdb}>Load</LoadingButton>
</form>
