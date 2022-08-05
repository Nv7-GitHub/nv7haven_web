<script lang="ts">
  import Spinner from "$lib/components/Spinner.svelte";
  import { url } from "$lib/consts";
  import Chart from "chart.js/auto";
  import { onMount, tick } from "svelte";

  let loading = true;

  type CmdCount = {
    time: number,
    timestring: string,
    counts: Record<string, number>,
  }

  type Data = {
    labels: string[],
    found: number[],
    elemcnt: number[],
    categorized: number[],
    combcnt: number[],
    usercnt: number[],
    servercnt: number[],
    commandcounts: CmdCount[],
  }

  let data: Data;
  let general: HTMLCanvasElement;
  let users: HTMLCanvasElement;
  let found: HTMLCanvasElement;

  onMount(() => {
    console.log("mount");
    fetch(url + "/eod_stats").then(async (val) => {
      data = await val.json();
      loading = false;
      await tick();
      draw();
    }) 
  })

  function draw() {
    let elemcnt = [
      {
        label: "Element Count",
        data: data.elemcnt,
        fill: true,
        borderColor: "rgba(54, 162, 235, 1)",
        backgroundColor: "rgba(54, 162, 235, 0.2)",
      },
      {
        label: "Combination Count",
        data: data.combcnt,
        fill: true,
        borderColor: "rgba(255, 206, 86, 1)",
        backgroundColor: "rgba(255, 206, 86, 0.2)",
      },
      {
        label: "Categorized Count",
        data: data.categorized,
        fill: true,
        borderColor: "rgba(75, 192, 192, 1)",
        backgroundColor: "rgba(75, 192, 192, 0.2)",
      }
    ];

    new Chart(general, {
      type: "line",
      data: {
        labels: data.labels,
        datasets: elemcnt,
      }
    });

    let usercnt = [
      {
        label: "User Count",
        data: data.usercnt,
        fill: true,
        borderColor: "rgba(153, 102, 255, 1)",
        backgroundColor: "rgba(153, 102, 255, 0.2)",
      },
      {
        label: "Server Count",
        data: data.servercnt,
        fill: true,
        borderColor: "rgba(255, 159, 64, 1)",
        backgroundColor: "rgba(255, 159, 64, 0.2)",
      }
    ];

    new Chart(users, {
      type: "line",
      data: {
        labels: data.labels,
        datasets: usercnt,
      }
    });

    let foundcnt = [
      {
        label: "Found",
        data: data.found,
        fill: true,
        borderColor: "rgba(255, 99, 132, 1)",
        backgroundColor: "rgba(255, 99, 132, 0.2)",
      }
    ];

    new Chart(found, {
      type: "line",
      data: {
        labels: data.labels,
        datasets: foundcnt,
      }
    });
  }
</script>

<h1>EoD Statistics</h1>
<p class="lead">Check out EoD statistics! Learn more about EoD <a href="https://discord.gg/KPmbJmNtxQ">here</a>!</p>

{#if loading} 
<Spinner></Spinner>
{:else}
<h2>General Statistics</h2>
<p class="lead">Statistics on element count, combination count, and categorized element count.</p>
<canvas bind:this={general} style="width: 75vw; height: 30vw;"></canvas>

<h2>User Statistics</h2>
<p class="lead">Statistics on user count and server count.</p>
<canvas bind:this={users} id="usercnt" style="width: 75vw; height: 30vw;"></canvas>

<h2>Elements Found</h2>
<p class="lead">The number of elements found, across all servers and users.</p>
<canvas bind:this={found} id="found" style="width: 75vw; height: 30vw;"></canvas>
{/if}