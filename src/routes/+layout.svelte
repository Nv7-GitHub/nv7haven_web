<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";
  import logo from "$lib/assets/navbar.svg";
  import { onMount } from "svelte";

  type Page = {
    name: string;
    path: string;
  };

  const structure: Record<string, Page[]> = {
    Fun: [
      { name: "Hella", path: "/hella" },
      { name: "Names", path: "/names" },
      { name: "NameFind", path: "/namefind" },
      { name: "Chess", path: "/chess" },
    ],
    Utilities: [
      { name: "FileHaven", path: "/filehaven" },
      { name: "QwikNotes", path: "/notes" },
      { name: "BreakDown", path: "/breakdown" },
      { name: "CSV", path: "/csv" },
      { name: "Studier", path: "https://studier.nv7haven.com" },
      { name: "When3meet", path: "https://w3m.evang.dev" },
    ],
    Social: [
      { name: "BestEver", path: "/bestever" },
      { name: "Elem7", path: "/e7" },
      { name: "EoD", path: "/eod" },
    ],
    Puzzles: [
      { name: "LetterBoxed", path: "/letterboxed" },
      { name: "WordSearch", path: "/wordsearch" },
      { name: "WSGen", path: "/wsgen" },
    ],
    Dev: [{ name: "B#", path: "/bsp" }],
  };

  // Theme
  /*enum Theme {
    LIGHT = "light",
    DARK = "dark",
    AUTO = "auto",
  }*/

  const themes = ["light", "dark", "auto"];
  const themeIcons: Record<string, string> = {
    light: "brightness-high-fill",
    dark: "moon-stars-fill",
    auto: "circle-half",
  };

  let theme = "auto";

  function switchTheme(newTheme: string) {
    theme = newTheme;
    localStorage.setItem("theme", newTheme);
    updateColorMode();
  }

  function updateColorMode() {
    switch (theme) {
      case "light":
        document.documentElement.setAttribute("data-bs-theme", "light");
        break;

      case "dark":
        document.documentElement.setAttribute("data-bs-theme", "dark");
        break;

      case "auto":
        document.documentElement.setAttribute(
          "data-bs-theme",
          window.matchMedia("(prefers-color-scheme: dark)").matches
            ? "dark"
            : "light",
        );
        break;
    }
  }

  // Auto theme
  onMount(() => {
    theme = localStorage.getItem("theme") ?? "auto";
    updateColorMode();
    window
      .matchMedia("(prefers-color-scheme: dark)")
      .addEventListener("change", () => {
        updateColorMode();
      });
  });

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
  });
</script>

<svelte:head>
  <title>Nv7's Projects</title>
</svelte:head>

<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <div class="container-fluid">
    <a class="navbar-brand" href="/">
      <img
        src={logo}
        width="30"
        height="30"
        class="d-inline-block align-top"
        alt="logo"
      />
      Nv7's Projects
    </a>

    <button
      class="navbar-toggler"
      type="button"
      data-bs-toggle="collapse"
      data-bs-target=".navbar-collapse"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon" />
    </button>

    <nav class="collapse navbar-collapse justify-content-end">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a
            class="nav-link"
            href="https://discord.com/oauth2/authorize?client_id=788185365533556736&scope=bot%20applications.commands&permissions=2617388096&redirect_uri=https%3A%2F%2Fnv7haven.com"
            >Nv7 Bot</a
          >
        </li>

        <li class="nav-item">
          <a
            class="nav-link"
            href="https://urbandictionary.store/products/sweatshirt?defid=7227413"
            >Bobby</a
          >
        </li>

        {#each Object.entries(structure) as [key, values]}
          <li class="nav-item dropdown">
            <span
              class="nav-link dropdown-toggle cursor"
              class:active={active == key}
              data-bs-toggle="dropdown">{key}</span
            >
            <ul class="dropdown-menu dropdown-menu-end">
              {#each values as value}
                <a class="dropdown-item cursor" href={value.path}>
                  {value.name}
                </a>
              {/each}
            </ul>
          </li>
        {/each}

        <li class="nav-item dropdown">
          <span
            class="nav-link dropdown-toggle cursor"
            data-bs-toggle="dropdown"
            ><i class={"bi bi-" + themeIcons[theme]} /></span
          >
          <ul class="dropdown-menu dropdown-menu-end">
            {#each themes as t}
              <li
                class="dropdown-item cursor"
                on:click={() => {
                  switchTheme(t);
                }}
              >
                <i class={"bi bi-" + themeIcons[t]} />
                <span class="capitalize">{t}</span>
              </li>
            {/each}
          </ul>
        </li>
      </ul>
    </nav>
  </div>
</nav>

<div class="container text-center mt-3">
  <slot />
</div>

<style>
  .cursor {
    cursor: pointer;
  }

  .capitalize {
    text-transform: capitalize;
  }
</style>
