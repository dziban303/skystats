<script>
    import { onDestroy, onMount } from 'svelte';

    export let endpoint;
    export let title;
    export let columns = [];

    let data = [];
    let loading = true;
    let error = null;
    let interval = null;

    async function fetchData() {
        console.log('fetchData')
        try {
            const response = await fetch(endpoint);
            if (!response.ok) {
                throw new Error(`${response.status}`);
            }
            const result = await response.json();
            data = result;
            error = null;
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }

    onMount(() => {
        fetchData();
        interval = setInterval(fetchData, 10000);
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
    <div class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>Something went wrong: {error}</span>
    </div>
{:else if data.length === 0}
    <div class="alert alert-info">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>No data available</span>
    </div>
{:else}

<div class="card bg-base-100 w96 shadow-sm">
    <div class="card-body">
        <h2 class="card-title">{title}</h2>
        <div class="overflow-x-auto">
            <table class="table text-center">
                <thead>
                    <tr class="uppercase tracking-wider">
                        {#each columns as column}
                            <th>{column.header}</th>
                        {/each}
                    </tr>
                </thead>
                <tbody>
                    {#each data as aircraft}
                    <tr class="hover:bg-base-300">
                        {#each columns as column}
                            <td class={column.class || ''}>
                                {#if column.formatter}
                                    {@html column.formatter(aircraft[column.field])}
                                {:else}
                                    {aircraft[column.field] || '-'}
                                {/if}
                            </td>
                        {/each}
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>
    </div>
</div>
{/if}