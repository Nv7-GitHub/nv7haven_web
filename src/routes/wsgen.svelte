<script lang="ts">
  const CHARS = "abcdefghijklmnopqrstuvwxyz";

  let val = "";
  $: {
    val = val.toLowerCase()
  }
  let items: string[] = [];

  let width: number;
  let height: number;
  let output = "";

  let copied = false;

  function duplBoard(board: string[][]): string[][] {
    return [...board].map(v => [...v]);
  }

  function generate() {
    let board = Array.from({length: height}, () => new Array(width).fill(""));

    // Put in words
    for (let word of items) {
      let finished = false;
      while (!finished) {
        // Calculate random vals
        let startx = Math.floor(Math.random() * width);
        let starty = Math.floor(Math.random() * height);
        let offx = Math.floor(Math.random() * 2) - 1;
        let offy = Math.floor(Math.random() * 2) - 1;
        if (offx == 0 && offy == 0) {continue;} // No offset
        let cop = duplBoard(board);

        let fail = false;
        for (let [i, lett] of Array.from(word).entries()) {
          // Calculate indexes
          let rv = starty + (offy * i);
          let cv = startx + (offx * i);
          if (rv < 0 || rv >= height || cv < 0 || cv >= width) { // Out of bounds
            fail = true;
            break;
          }
          if (board[rv][cv] != "") { // Intersects with other word
            fail = true;
            break;
          }
          cop[rv][cv] = lett;
        }

        if (!fail) {
          finished = true;
          board = cop;
        }
      }
    }

    // Put in random chars
    for (let [r, row] of board.entries()) {
      for (let [c, col] of row.entries()) {
        if (col == "") {
          board[r][c] = CHARS[Math.floor(Math.random() * CHARS.length)];
        }
      }
    }

    // Save
    output = board.map(v => v.join("")).join("\n");
    copied = false;
  }
</script>

<h1>Word Search Generator</h1>
<p class="lead">This website has a Word Search solver, so why not a word search generator? This generates word searches that can be solved using the Word Search solver!</p>

<form class="input-group mb-3">
  <input type="text" class="form-control" bind:value={val} placeholder="Enter a word for the word search..."/>
  <button class="btn btn-outline-secondary" type="submit" on:click|preventDefault={() => {items.push(val); items = items; val = "";}}><i class="bi bi-plus-lg"></i></button>
</form>

<ul class="list-group mb-3"> 
  {#each items as item, i (item)}
    <li class="list-group-item d-flex justify-content-between">
      <div class="p-0 m-0 flex-grow-1" style="transform: translateY(20%);">{item}</div> 
      <button class="btn btn-danger" on:click={() => {items.splice(i, 1); items = items}}>
        <i class="bi bi-trash"></i>
      </button>
    </li>
  {/each}
</ul>

<form class="input-group mb-5">
  <div class="form-floating">
    <input type="number" id="width" bind:value={width} class="form-control" placeholder="Width..."/>
    <label for="width">Width</label>
  </div>
  <div class="form-floating">
    <input type="number" bind:value={height} id="height" class="form-control" placeholder="Height..."/>
    <label for="height">Height</label>
  </div>
  <button type="submit" class="btn btn-primary" on:click|preventDefault={generate}>Generate!</button>
</form>

{#if output != ""}
  <textarea class="form-control" readonly rows={output.length - output.replaceAll("\n", "").length + 1}>{output}</textarea>

  <button class="btn btn-success mt-3 mb-5" on:click={async () => {await navigator.clipboard.writeText(output); copied = true;}}><i class="bi" class:bi-clipboard={!copied} class:bi-clipboard-check={copied}></i> {copied ? "Copied!" : "Copy"}</button>
{/if}
