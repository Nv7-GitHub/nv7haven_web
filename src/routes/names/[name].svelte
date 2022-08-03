<script lang="ts">
  import { page } from "$app/stores";
import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";

  let name = $page.params["name"];

  type Data = {
    Name: string,
    IsMale: boolean,
    Population: number,
  }
  let data: Data;
  async function fetchData() {
    let res = await fetch(url + "/get_name/" + encodeURIComponent(name));
    data = await res.json();
  }
  fetchData();
</script>

<h1>Names</h1>

{#if data}
<p class="lead">
  The name {data.Name} is a {data.IsMale ? "male" : "female"} name. There are {data.Population.toLocaleString()} people who have this name!
</p>
{:else}
<Spinner></Spinner>
{/if}
