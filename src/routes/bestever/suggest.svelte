<script lang="ts">
  import LoadingButton from "$lib/components/LoadingButton.svelte";
  import { url } from "$lib/consts";

  let sugg = "";
  let loading = false;

  async function suggest() {
    loading = true;
    await fetch(url + "/bestever_new_suggest/" + encodeURIComponent(sugg))
  }
</script>

<p class="lead">Enter your suggestion below! If people like it, you might see it in the leaderboard!</p>

<form class="input-group">
  <input class="form-control" type="text" bind:value={sugg} placeholder="Suggestion..."/>
  <LoadingButton style="btn-primary" loading={loading} type="submit" on:click={suggest}>Suggest</LoadingButton>
</form>