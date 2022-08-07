<script lang="ts">
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";
  import { onMount } from "svelte";

  let left = 0;
  let right = 0;
  let data: Record<number, string>;
  let waiting = true;

  async function update_lr() {
    waiting = true;

    let keys: number[] = [];
    while (keys.length != 2 || has_done(keys)) {
      keys = await fetch_lr();
    }
    
    left = keys[0];
    right = keys[1];
    localStorage.setItem("bestever_" + left.toString() + "," + right.toString(), "1");
    waiting = false;
  }

  function has_done(keys: number[]) {
    return localStorage.getItem("bestever_" + keys[0].toString() + "," + keys[1].toString()) == "1";
  }

  async function fetch_lr(): Promise<number[]> {
    let res = await fetch(url + "/bestever_get_suggest");
    data = await res.json();
    return Object.keys(data).map(k => parseInt(k)).sort();
  }

  async function vote(rightVal: boolean) {
    waiting = true;
    let val = rightVal ? left : right;
    await fetch(url + "/bestever_vote/" + val.toString());
    await update_lr();
  }

  let ready = false;
  async function getReady() {
    await update_lr();
    ready = true;
  }

  onMount(getReady)
</script>

<p class="lead">What is better?</p>

{#if ready}
<div class="btn-group">
  <button class="btn btn-primary" disabled={waiting} on:click={() => {vote(false)}}>{data[left]}</button>
  <button class="btn btn-info" disabled={waiting} on:click={() => {vote(true)}}>{data[right]}</button>
  <button class="btn btn-danger" on:click={update_lr} disabled={waiting}>Skip</button>
</div>
{:else}
<Spinner></Spinner>
{/if}
