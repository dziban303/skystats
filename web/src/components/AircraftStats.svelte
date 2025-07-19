<script>
// @ts-nocheck
    import { onMount, onDestroy } from 'svelte';

    export let endpoint;
    export let title;
    export let columns = [];
    export let icon = null;
    export let iconColor = '';
    export let iconBgColor = '';

    let refreshRate = 60000
    let data = [];
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
        interval = setInterval(fetchData, refreshRate);
    })

    onDestroy(() => {
        if (interval) {
            clearInterval(interval);
        }
    });
</script>

<div class="grid grid-cols-1 gap-2">

<div>
<h1 class="card-title mb-3 flex items-center gap-2">
    {#if icon}
        <div class="w-8 h-8 {iconBgColor} rounded-lg flex items-center justify-center">
            <svelte:component this={icon} class="w-5 h-5 {iconColor}" />
        </div> 
    {/if}
    {title}
</h1>
<!-- <div class="card bg-base-100 w96 shadow-sm rounded-xl border border-secondary hover:shadow-md transition-all duration-200"> -->
 <div class="card bg-base-100 w96 mb-4 shadow-sm rounded-xl hover:shadow-md transition-all duration-200">
    <div class="card-body">
        <!-- <div class="divider mt-1 mb-1 -mx-6"></div> -->
        <div class="overflow-x-auto">
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
                <table class="table">
                    <thead>
                        <tr class="uppercase tracking-wider">
                            {#each columns as column}
                                <th>{column.header}</th>
                            {/each}
                        </tr>
                    </thead>
                    <tbody>
                        {#each data as aircraft}
                        <!-- <tr class="hover:bg-base-300"> -->
                         <tr>
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

        {/if}
        </div>
    </div>
</div>
</div>
</div>