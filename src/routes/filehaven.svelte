<script lang="ts">
import { url } from "$lib/consts";


  let hasFile = false;
  let loading = false;
  let percent = 0;
  let link = "";
  let customFile: HTMLInputElement;

  function fileChange() {
    hasFile = true;
  }

  function upload() {
    loading = true;
    percent = 0;

    // Form data
    let dat = new FormData();
    dat.append("file", customFile.files![0]);
    let size = customFile.files![0].size;

    // Add handlers
    let req = new XMLHttpRequest();
    req.upload.addEventListener("progress", (e) => {
      percent = e.loaded/size * 100;
    })
    req.onreadystatechange = () => {
      if (req.readyState == XMLHttpRequest.DONE) {
        customFile.type = "text";
        customFile.type = "file";
        loading = false;
        hasFile = false;
        link = url + "/get_file/" + req.responseText;
      }
    }

    // Send request
    req.open("POST", url + "/upload");
    req.timeout = 99999999999999;
    req.send(dat);
  }
</script>

<h1>FileHaven</h1>
<p class="lead">Have you ever wanted to share a file but have had no place to share it? Well this is where you can share it!</p>

<div class="input-group mb-3 text-left">
  <input type="file" class="form-control" bind:this={customFile} name="filename" on:change={fileChange}>
  <button class="btn btn-success" disabled={!hasFile || loading} on:click={upload}>Upload Files</button>
</div>

{#if loading}
  <div class="progress" vg-if="c.Loading">
    <div class="progress-bar progress-bar-striped progress-bar-animated" style={"width: " + percent + "%"}></div>
  </div>
{/if}

{#if link != ""}
  <p class="lead"><a href={link}>Download Link</a></p>
{/if}