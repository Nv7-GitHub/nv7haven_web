<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let text: string;
  export let style: string;
  export let type = "button";
  export let loading = false;

  function click() {
    dispatch('click');
  }
</script>

<button disabled={loading} class={"btn " + style} type={type} on:click|preventDefault={click}>
  {#if loading}
    <span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
  {/if}

  {text}
</button>