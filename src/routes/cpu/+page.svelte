<script lang="ts">
    import { url } from "$lib/consts";
    import { onMount } from "svelte";

    onMount(getUsage);

    let usage = {
        cpu_usage: 0,
        memory_usage: 0,
        cpu_temperature: 0,
        disk_usage: 0,
    };
    let loading = false;
    async function getUsage() {
        loading = true;
        const res = await fetch(url + "/cpu");
        usage = await res.json();
        loading = false;
    }
</script>

<h1>System Usage</h1>
<p class="lead">See the CPU usage of the Nv7haven server!</p>
<button
    class="btn btn-lg w-100 btn-primary mb-5"
    disabled={loading}
    on:click={getUsage}
>
    <i class="bi bi-cloud-download-fill"></i> Refresh</button
>
<h2>CPU Usage</h2>
<div class="progress mb-5" role="progressbar">
    <div class="progress-bar" style="width: {usage.cpu_usage}%">
        {usage.cpu_usage.toFixed(1)}%
    </div>
</div>

<h2>Memory Usage</h2>
<div class="progress mb-5" role="progressbar">
    <div class="progress-bar bg-info" style="width: {usage.memory_usage}%">
        {usage.memory_usage.toFixed(1)}%
    </div>
</div>

<h2>CPU Temperature</h2>
<div class="progress mb-5" role="progressbar">
    <div
        class="progress-bar bg-warning"
        style="width: {(usage.cpu_temperature / 85) * 100}%"
    >
        {usage.cpu_temperature.toFixed(1)} C
    </div>
</div>

<h2>Disk Usage</h2>
<div class="progress mb-5" role="progressbar">
    <div class="progress-bar bg-success" style="width: {usage.disk_usage}%">
        {usage.disk_usage.toFixed(1)}%
    </div>
</div>

<style>
    .progress {
        height: 2rem;
        border-radius: 0.75rem;
    }

    .progress div {
        font-size: 1rem;
    }
</style>
