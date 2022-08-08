<script lang="ts">
  import wordsRaw from "$lib/assets/words.txt?raw";
  import Spinner from "$lib/components/Spinner.svelte";
  import { load, solve } from "$lib/letterboxed/letterboxed";
  import { onMount, tick } from "svelte";

  let words = wordsRaw.split("\n");
  let ready = false;

  let first = "xie";
  let second = "mcn";
  let third = "ouy";
  let fourth = "lqr";

  let res: string[] = [];
  let solving = false;

  async function getReady() {
    await load(); // Load WASM
    ready = true;
  }

  async function solv() {
    solving = true;
    await tick();
    res = solve([first, second, third, fourth], words);
    solving = false;
  }

  function remove(val: string) {
    words = words.filter(v => v !== val);
    solv();
  }

  onMount(getReady)
</script>

<h1>LetterBoxed</h1>
<p class="lead">This is a program I wrote to solve <a href="https://www.nytimes.com/puzzles/letter-boxed">New York Times "Letter Boxed" puzzle</a>. You can find the source code <a href="https://github.com/Nv7-GitHub/FunRepository/blob/master/LetterBoxed/src/lib.rs">here</a>. The default values are from the Letter Boxed puzzle on July 19, 2022.</p>

{#if ready}
<form class="input-group">
  <input type="text" class="form-control" minlength="3" maxlength="3" bind:value={first}/>
  <input type="text" class="form-control" minlength="3" maxlength="3" bind:value={second}/>
  <input type="text" class="form-control" minlength="3" maxlength="3" bind:value={third}/>
  <input type="text" class="form-control" minlength="3" maxlength="3" bind:value={fourth}/>
  <button type="submit" on:click|preventDefault={solv} class="btn btn-primary" disabled={solving}>Solve</button>
</form>
{:else}
<Spinner></Spinner>
{/if}

{#if res.length > 0}
<h2 class="mt-3">Output</h2>
<ul class="list-group">
{#each res as v}
  <li class="list-group-item d-flex justify-content-between">
    <div class="p-0 m-0 flex-grow-1" style="transform: translateY(20%);">{v}</div> 
    <button class="btn btn-danger" on:click={() => {remove(v)}} disabled={solving}>
      <i class="bi bi-trash"></i>
    </button>
  </li>
{/each}
</ul>
{/if}

