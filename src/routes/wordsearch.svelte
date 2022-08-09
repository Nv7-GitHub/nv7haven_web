<script lang="ts">
  let inp = `xymbitackkj
kxfyofasokd
sxtmttcbmwh
rranatbspkh
htotcdwlzor
oeptvvzopif
rkemcsgjndn
ssecqacizlw
eahcfxrbkyk
cbslmqxtrla`;

  let board: string[][] = [];

  function load() {
    for (let line of inp.toLowerCase().trim().split("\n")) {
      let row = [];
      for (let char of line) {
        row.push(char);
      }
      board.push(row);
    }
    reset_matched();
    board = board;
  }

  let query = "";
  let matched: boolean[][] = [];

  function reset_matched() {
    matched = Array.from({length: board.length}, (_, i) => new Array(board[i].length).fill(false));
  }

  $: {
    reset_matched();
    let q = query.toLowerCase();
    
  }
</script>

<h1>WordSearch</h1>
<p class="lead">Tired of solving word searches by hand? Have this solve it for you!</p>

{#if board.length == 0}
<textarea class="form-control" bind:value={inp} rows={inp.length - (inp.replaceAll("\n", "").length) + 1}></textarea>
<button class="btn btn-primary mt-2" on:click={load}>Load Board</button>
{:else}
<input type="text" class="form-control" bind:value={query} placeholder="Word to find..."/>

<div class="table-responsive">
  <table class="table" vg-if="c.BoardCalculated">
    <tbody>
      {#each board as row, r}
        <tr>
          {#each row as char, c}
            <td class:table-success={matched[r][c]}>{char}</td>
          {/each}
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<button class="btn btn-danger" on:click={() => {board = [];}}>Reset</button>
{/if}