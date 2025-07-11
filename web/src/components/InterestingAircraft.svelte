<script>
    import {onMount, onDestroy} from 'svelte'

    export let endpoint;
    export let title;
    export let icon;
    export let iconColor;
    export let iconBgColor;

    let refreshRate = 10000
    let data = [];
    let loading = true;
    let error = null;
    let interval = null;

    async function fetchData() {
        
        try {
            const response = await fetch(endpoint);
            if(!response.ok) {
                throw new Error(`{response.status}`);
            }
            const result = await response.json();
            data = result;
            error = null
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }

    onMount(() => {
        fetchData();
        interval = setInterval(fetchData, refreshRate)
    })

    onDestroy(() => {
        if (interval) {
            clearInterval(interval)
        }
    });

</script>

<div class="card bg-base-100 w96 shadow-sm rounded-xl border border-gray-200 hover:shadow-md transition-all duration-200">
    <div class="card-body">
        <h2 class="card-title -mt-2 -mb-2 flex items-center gap-2">
            {#if icon}
                <div class="w-8 h-8 {iconBgColor} rounded-lg flex items-center justify-center">
                    <svelte:component this={icon} class="w-5 h-5 {iconColor}" />
                </div> 
            {/if}
            {title}
        </h2>
        <div class="divider mt-1 mb-1 -mx-6"></div>
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
                    <thead class="uppercase tracking-wider">
                        <tr>
                            <th>Flight</th>
                            <th>Reg</th>
                            <th>Operator</th>
                            <th>Type</th>
                            <th>Category</th>
                            <th>Tags</th>
                            <th>Last Seen</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each data as aircraft}
                        <tr>
                            <td>{aircraft.flight || '-'}</td>
                            <td>{aircraft.registration}</td>
                            <td>{aircraft.operator}</td>
                            <td>{aircraft.type}</td>
                            <td>{aircraft.category}</td>
                            <!-- TODO -->
                            <td>{aircraft.tag1}</td>
                            <td>{aircraft.seen}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            {/if}
        </div>
    </div>
</div>

