<script lang="ts">
  import { url } from "$lib/consts";
  import { onMount } from "svelte";
  let invalid = false;
  let valid = false;
  let feedback = "";

  type Row = {
    name: string,
    count: number,
  }
  let rows: Row[] = [];
  let total = 0;
  let found = 0;

  let name = "";

  let loading = false;

  function sortRows() {
    rows.sort((a, b) => {
      if (a.count < b.count) return 1;
      if (a.count > b.count) return -1;
      return 0;
    });
    rows = rows;
  }

  // Init DB
  let db: IDBDatabase;
  function initDb() {
    loading = true;

    let req = indexedDB.open("namefind", 1);
    req.onupgradeneeded = (ev: any) => {
      db = ev.target.result;
      let store = db.createObjectStore("namefind", {keyPath: "name"});
      store.createIndex("name", "name", {unique: true})
      store.createIndex("count", "count", {unique: false})
    }
    req.onsuccess = (ev: any) => {
      db = ev.target.result;
      
      // Load DB
      let trans = db.transaction(["namefind"], "readonly");
      let store = trans.objectStore("namefind");
      let req = store.getAll();
      req.onsuccess = (ev: any) => {
        rows = ev.target.result;
        for (let row of rows) {
          found += row.count;
        }
        sortRows();
      }

      // Get total
      let totalVal = localStorage.getItem("total");
      if (totalVal) {
        total = parseInt(totalVal);
        loading = false;
      } else {
        fetch(url + "/name_count/all").then(async (val) => {
          total = parseInt(await val.text());
          localStorage.setItem("total", total.toString());
          loading = false;
        })
      }
    }
  }

  onMount(initDb);

  type nameCountResult = {
    Name: string,
    Count: number,
  }

  async function addName() {
    loading = true;

    let res = await fetch(url + "/name_count/" + encodeURIComponent(name));
    if (res.status != 200) {
      invalid = true;
      valid = false;
      feedback = "That name is not in our database!"
      loading = false;
      return;
    }

    let val: nameCountResult = await res.json();

    // Check if already exists, if not, add to DB
    let trans = db.transaction(["namefind"], "readonly");
    let store = trans.objectStore("namefind");
    let req = store.get(val.Name);
    req.onsuccess = (ev: any) => {
      if (req.result) {
        // Name is in DB
        feedback = "You already did that name!"
        invalid = true;
        valid = false;
        loading = false;
        return;
      }

      // Add to DB
      let trans = db.transaction(["namefind"], "readwrite");
      let store = trans.objectStore("namefind");
      let newReq = store.add({
        name: val.Name,
        count: val.Count,
      });
      newReq.onsuccess = () => {
        feedback = "New name found!";
        valid = true;
        invalid = false;
        rows.push({
          name: val.Name,
          count: val.Count,
        });
        found += val.Count;
        name = "";
        loading = false;
        sortRows();
        return;
      }
    }
  }
</script>

<h1>NameFind</h1>
<p class="lead">How many names can you find?</p>
<span class="text-secondary">If you want to cheat, you can find the most common names <a href="/names">here</a>.</span>

<form class="input-group mt-3 has-validation">
  <input type="text" class="form-control" placeholder="Name..." class:is-invalid={invalid} class:is-valid={valid} bind:value={name}/>
  <button class="btn btn-primary" type="submit" disabled={loading} on:click|preventDefault={addName}>Go!</button>
  {#if invalid || valid}
    <div class:valid-feedback={valid} class:invalid-feedback={invalid}>{feedback}</div>
  {/if}
</form>

<h2 class="mt-5">Progress</h2>
<p class="lead">
  You have found the names of {found.toLocaleString()} people, out of a total of {total.toLocaleString()} people.
</p>
<div class="progress">
  <div class="progress-bar" role="progressbar" style={"width: " + found/total*100 + "%"}></div>
</div>

<h2 class="mt-5">Names Found</h2>
<table class="table">
  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">Name</th>
      <th scope="col">Population</th>
    </tr>
  </thead>
  <tbody>
    {#each rows as row, i}
    <tr>
      <th scope="row" vg-content="i+1">{i+1}</th>
      <td>{row.name}</td>
      <td>{row.count.toLocaleString()}</td>
    </tr>
    {/each}
  </tbody>
</table>