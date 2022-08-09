<script lang="ts">
  const orig = `xymbitackkj
kxfyofasokd
sxtmttcbmwh
rranatbspkh
htotcdwlzor
oeptvvzopif
rkemcsgjndn
ssecqacizlw
eahcfxrbkyk
cbslmqxtrla`;
  let inp = orig;

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
    let q = Array.from(query.toLowerCase());
    for (let [r, row] of board.entries()) { // Loop over rows & cols
      for (let [c, col] of row.entries()) {
        if (board[r][c] == q[0]) { // Character matches, see if areas nearby match
          for (let roff = -1; roff <= 1; roff++) { // Loop through offsets
            for (let coff = -1; coff <= 1; coff++) {
              if (roff == 0 && coff == 0) { // Not an offset
                continue;
              }

              // Loop through letters & match
              let valid = true;
              for (let [i, lett] of q.entries()) {
                if (i == 0) {
                  continue; // We have already matched first letter
                }
                
                // Calc indexes
                let rv = r + (roff * i);
                let cv = c + (coff * i);

                // Checks
                if (rv < 0 || rv >= board.length || cv < 0 || cv >= board[rv].length) {
                  valid = false;
                  break; // Continue to next offset if out of bounds
                }
                if (board[rv][cv] != lett) {
                  valid = false;
                  break; // Continue if doesn't match
                }
              }

              if (valid) { // Found word, add to matched
                for (let i = 0; i < q.length; i++) {
                  matched[r + (roff * i)][c + (coff * i)] = true;
                }
              }
            }
          }
        }
      }
    }
  }
</script>

<h1>WordSearch</h1>
<p class="lead">Tired of solving word searches by hand? Have this solve it for you!</p>

{#if board.length == 0}
<textarea class="form-control" bind:value={inp} rows={inp.length - (inp.replaceAll("\n", "").length) + 1}></textarea>
<button class="btn btn-primary mt-2" on:click={load}>Load Board</button>
{:else}

{#if inp == orig}
<div class="text-start">
  <p>This is the default puzzle! Some words to look for are:</p>
  <ul>
    <li>cow</li>
    <li>tomato</li>
    <li>tractor</li>
    <li>horse</li>
    <li>sheep</li>
    <li>basket</li>
    <li>cat</li>
  </ul>
</div>
{/if}

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