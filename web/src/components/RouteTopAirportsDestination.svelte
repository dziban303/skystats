<script>
    import { onMount, onDestroy } from 'svelte';

    let data = [];
    let endpoint = '/api/stats/routes'
    let loading = true;
    let error = null;
    let interval = null;

    async function fetchData() {

        try {
            const response = await fetch(endpoint);
            if (!response.ok) {
                throw new Error(`${response.status}`);
            }
            const result = await response.json();
            data = result["top_destination_airports"];
            error = null;
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }

    onMount(() => {
        fetchData();
        interval = setInterval(fetchData, 60000);
    })

    onDestroy(() => {
        if (interval) {
            clearInterval(interval);
        }
    });
</script>

 {#if loading}
    <div class="flex justify-center py-8">
        <span class="loading loading-spinner loading-lg"></span>
    </div>
{:else if error}
    <div class="flex alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>Something went wrong: {error}</span>
    </div>
{:else if data.length === 0}
    <div class="alert alert-info">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>No data available</span>
    </div>
{:else}
    <ul class="list bg-base-100 rounded-box shadow-md">
    <li class="p-4 pb-2 text-xs opacity-60 tracking-wide">Top Airports - Destination</li>


    {#each data as airport, index}
    <li class="list-row">
        <div class="text-4xl font-thin opacity-30 tabular-nums">{String(index + 1).padStart(2, '0')}</div>
        <div class="list-col-grow">
         <div class="font-medium">{airport.airport_code}</div>
        <div class="text-xs uppercase font-semibold opacity-60">{airport.country}</div>
        </div>
        <div class="text-right">
            <div class="font-semibold">{airport.count.toLocaleString()}</div>
            <div class="text-xs opacity-60">flights</div>
        </div>
    </li>
    {/each}
    </ul>
{/if}