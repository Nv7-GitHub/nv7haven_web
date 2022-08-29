<script lang="ts">
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import { url } from "$lib/consts";


  let text: string;

  let loading = false;
  let output = "";

  async function hellaify() {
    loading = true;
    let res = await fetch(url + "/hella/" + encodeURIComponent(text));
    output = await res.text();
    loading = false;
  }
</script>

<h1>Hella</h1>
<p class="lead">
  Hella is inspired by an episode of South Park in which Cartman replaces adds "hella-" to the start of all adjectives. Wondering if I could code that, I used a library called <a href="https://github.com/jdkato/prose">prose</a> to identify adjectives and add "hella-" to the start!
</p>

<form class="form">
  <textarea class="form-control" bind:value={text}/>

  <div class="mt-3">
    {#if output != ""}
      <p>{output}</p>
    {/if}
  </div>

  <LoadingButton type="submit" on:click={hellaify} style="btn-primary" loading={loading}>Hella-ify!</LoadingButton>
</form>