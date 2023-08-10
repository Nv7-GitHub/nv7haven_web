<script lang="ts">
  import Papa from "papaparse";
  import Chart from "chart.js/auto";
  import { tick } from "svelte";

  let file: HTMLInputElement;
  let graph: HTMLCanvasElement;
  let hasData = false;
  let data: string[][] | null = null;
  let shown: boolean[] = [];
  let chart: Chart;

  function updateRowDisplay(row: number) {
    if (shown[row]) {
      chart.data.labels![row] = data![row + 1][0];
      for (let i = 0; i < chart.data.datasets!.length; i++) {
        chart.data.datasets![i].data![row] = parseFloat(data![row + 1][i + 1]);
      }
    } else {
      chart.data.labels![row] = null;
      for (let i = 0; i < chart.data.datasets!.length; i++) {
        chart.data.datasets![i].data![row] = null;
      }
    }
    chart.update();
  }

  async function makeGraph() {
    if (chart) {
      chart.destroy();
    }

    // Make data structures
    let labels = [];
    let datasets = [];
    for (let col of data![0].slice(1)) {
      datasets.push({
        label: col,
        data: [] as number[],
      });
    }

    // Fill data structures
    for (let [i, row] of data!.slice(1).entries()) {
      if (!shown[i]) continue;

      labels.push(row[0]);
      for (let i = 1; i < row.length; i++) {
        datasets[i - 1].data.push(parseFloat(row[i]));
      }
    }

    // Make graph
    chart = new Chart(graph, {
      type: "line",
      options: {
        responsive: true,
        elements: {
          point: {
            radius: 0, // default to disabled in all datasets
          },
        },
        animation: false,
      },
      data: {
        labels,
        datasets,
      },
    });
  }

  async function read(ev: Event) {
    hasData = false;
    data = null;
    shown = [];
    let files: File[] = (ev.target! as any).files;
    if (files.length < 1) return;

    // Parse CSV
    Papa.parse(files[0], {
      complete: async (res) => {
        hasData = true;
        data = res.data as string[][];
        shown = Array(data.length - 1).fill(true);
        await tick();
        makeGraph();
      },
    });
  }
</script>

<h1>CSV Viewer</h1>
<p class="lead">
  I created this to view the flight data from my <a
    href="https://github.com/Nv7-Github/rock"
    target="_blank">flight computer</a
  >, but this can be used to view any CSV with the first column as the x-axis
  and numerical data! Upload a file below and view the data!
</p>

<input
  type="file"
  class="form-control"
  bind:this={file}
  name="filename"
  on:change={read}
  accept="text/csv"
/>

{#if hasData}
  <div
    class="chart-container d-flex justify-content-center"
    style="width: 100%; height: 50vh;"
  >
    <canvas bind:this={graph} class="mt-3" />
  </div>

  <h1>Data</h1>
  {#if data}
    <div style="width: 100%; overflow: scroll">
      <table class="table table-striped">
        <thead>
          {#each data[0] as col}
            <th scope="col">{col}</th>
          {/each}
          <th scope="col">display</th>
        </thead>
        <tbody>
          {#each data.slice(1) as row, rowind}
            <tr>
              {#each row as val, i}
                {#if i == 0}
                  <th scope="row">{val}</th>
                {:else}
                  <td>{val}</td>
                {/if}
              {/each}
              <td
                ><input
                  type="checkbox"
                  bind:checked={shown[rowind]}
                  on:change={() => {
                    updateRowDisplay(rowind);
                  }}
                /></td
              >
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
{/if}
