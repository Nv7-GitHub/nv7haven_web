<script lang="ts">
  import { page } from "$app/stores";
  let url = $page.url.pathname;
  page.subscribe((v) => {
    url = v.url.pathname;
  })
</script>

<h1>BestEver</h1>
<ul class="nav nav-pills nav-fill mb-3">
  <li class="nav-item">
    <a class="nav-link" class:active={url.endsWith("bestever")} href="/bestever">About</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" class:active={url.endsWith("suggest")} href="/bestever/suggest">Suggest</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" class:active={url.endsWith("vote")} href="/bestever/vote">Vote</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" class:active={url.endsWith("leaderboard")} href="/bestever/leaderboard">Leaderboard</a>
  </li>
</ul>
<slot></slot>