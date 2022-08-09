<script lang="ts">
  let val = "";
  let items: string[] = [];

  let width: number;
  let height: number;

  function generate() {
    console.log("TODO")
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

<form class="input-group">
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
