<script lang="ts">
  import { runBsp, load } from "$lib/bsp/bsp";
  import Spinner from "$lib/components/Spinner.svelte";
  import { onMount } from "svelte";

  let code = `[PRINT "Hello, World!"]`;
  let loading = true;

  let output = "";

  function run() {
    output = runBsp(code);
    if (output == "") {
      output = "⚠️: No Output";
    }
  }

  onMount(async () => {
    await load();
    loading = false;
  });
</script>

<h1>B#</h1>
<p class="lead">Run B# code here! Find the source and documentation for B# <a href="https://github.com/Nv7-Github/bsharp">here</a>!</p>

{#if loading}
<Spinner></Spinner>
{:else}
<h2>Code</h2>
<textarea class="form-control" placeholder='[PRINT "Hello, World!"]' bind:value={code}></textarea>
<button class="btn btn-primary mt-3 mb-3" on:click={run}>Run!</button>
{/if}

{#if output != ""}
<h2>Output</h2>
<pre class="block mt-3 text-start"><code>{output}</code></pre>
{/if}

<style>
  .block {
    background-color: #f5f5f5;
    border-radius: 10px;
    padding: 1em;
    max-height: 75vh;
    overflow: scroll;
  }
</style>