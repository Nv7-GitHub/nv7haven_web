<script lang="ts">
  import chesshtml from "$lib/assets/chess.html?raw";
    import { onMount } from "svelte";
  let frame: HTMLIFrameElement;
  onMount(() => {
    frame.contentWindow?.document.open();
    frame.contentWindow?.document.write(chesshtml);
    frame.contentWindow?.document.close();
  });
</script>

<h1>Chess</h1>
<iframe bind:this={frame} width="480" height="540" title="Chess"></iframe>