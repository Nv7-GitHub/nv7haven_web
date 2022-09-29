<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import logo from "$lib/assets/navbar.svg";

  type Page = {
    name: string,
    path: string,
  }

  const structure: Record<string, Page[]> = {
    "Fun": [{name: "Hella", path: "/hella"}, {name: "Names", path: "/names"}, {name: "NameFind", path: "/namefind"}, {name: "Chess", path: "/chess"}],
    "Utilities": [{name: "FileHaven", path: "/filehaven"}, {name: "QwikNotes", path: "/notes"}, {name: "BreakDown", path: "/breakdown"}],
    "Social": [{name: "BestEver", path: "/bestever"}, {name: "Elem7", path: "/e7"}, {name: "EoD", path: "/eod"}],
    "Puzzles": [{name: "LetterBoxed", path: "/letterboxed"}, {name: "WordSearch", path: "/wordsearch"}, {name: "WSGen", path: "/wsgen"}],
    "Dev": [{name: "B#", path: "/bsp"}],
  };

  let active = "";
  page.subscribe(() => {
    for (let [key, value] of Object.entries(structure)) {
      for (let val of value) {
        if ($page.url.pathname.startsWith(val.path)) {
          active = key;
          return;
        }
      }
    }
    active = "";
  })
</script>

<svelte:head>
  <title>Nv7's Projects</title>
</svelte:head>

<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <div class="container-fluid">
    <a class="navbar-brand" href="/">
    <img src={logo} width="30" height="30" class="d-inline-block align-top" alt="logo">
      Nv7's Projects
    </a>

    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target=".navbar-collapse" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <nav class="collapse navbar-collapse justify-content-end">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a class="nav-link" href="https://discord.com/oauth2/authorize?client_id=788185365533556736&scope=bot%20applications.commands&permissions=2617388096&redirect_uri=https%3A%2F%2Fnv7haven.com">Nv7 Bot</a>
        </li>

        <li class="nav-item">
          <a class="nav-link" href="https://urbandictionary.store/products/sweatshirt?defid=7227413">Bobby</a>
        </li>

        {#each Object.entries(structure) as [key, values]}
          <li class="nav-item dropdown">
            <span class="nav-link dropdown-toggle cursor" class:active={active == key} data-bs-toggle="dropdown">{key}</span>
            <ul class="dropdown-menu">
              {#each values as value}
                <li class="dropdown-item cursor" on:click={() => {goto(value.path)}}>
                  {value.name}
                </li>
              {/each}
            </ul>
          </li>
        {/each}
      </ul>
    </nav>
  </div>
</nav>

<div class="container text-center mt-3">
  <slot></slot>
</div>

<style>
  .cursor {
    cursor: pointer;
  }
</style>