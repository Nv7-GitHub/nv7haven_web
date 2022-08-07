<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  export let style: string;
  export let type = "button";
  export let loading: boolean;
  export let disabled = false;

  function click() {
    dispatch('click');
  }
</script>

<button disabled={loading || disabled} class={"btn " + style} type={type} on:click|preventDefault={click}>
  {#if loading}
    <span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
  {/if}

  <slot></slot>
</button>