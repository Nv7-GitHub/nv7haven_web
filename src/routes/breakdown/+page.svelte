<script lang="ts">
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import { url } from "$lib/consts";

  type Result = {
    Part: string,
    Words: string,
  }
  let results: Result[] = [];

  let input = "";
  let loading = false;

  async function getResults() {
    loading = true;
    let res = await fetch(url + "/breakdown/" + encodeURIComponent(input));
    results = await res.json();
    input = "";
    loading = false;
  }
</script>

<h1>Breakdown</h1>
<p class="lead">Use Natural Language Processing to breakdown a sentence! Uses a Go library called <a href="https://github.com/jdkato/prose">prose</a>.</p>

<textarea bind:value={input} class="form-control"></textarea>

{#if results.length > 0}
  <table class="table table-striped mt-3">
    <thead>
      <tr>
        <th scope="col">Words</th>
        <th scope="col">Part of Speech</th>
      </tr>
    </thead>
    <tbody>
      {#each results as result}
        <tr>
          <td>{result.Words}</td>
          <td>{result.Part}</td>
        </tr>
      {/each}
    </tbody>
  </table>
{/if}

<LoadingButton style={"btn-primary" + (results.length == 0 ? " mt-3" : "")} loading={loading} on:click={getResults}>BreakDown</LoadingButton>
