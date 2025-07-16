<script>
    import { onMount, onDestroy } from 'svelte'
    import NumberFlow from '@number-flow/svelte'
    import { fade } from 'svelte/transition'
    import { flip } from 'svelte/animate'
    import { IconBinoculars } from '@tabler/icons-svelte';

    let endpoint = 'api/stats/above'

    let refreshRate = 2000
    let data = []
    let loading = true;
    let error = null;
    let interval = null;

    async function fetchData() {
        console.log('refresh above')
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

<style>
    .table-container {
        /* Calculate height for 5 rows + header */
        /* This uses CSS calc with custom properties that work with DaisyUI */
        min-height: calc(5 * 3rem + 3rem); /* 5 rows at ~3rem each + header ~3rem */
        height: calc(5 * 3rem + 3rem);
        display: flex;
        flex-direction: column;
    }
</style>

<div>
<h1 class="card-title mb-3 flex items-center gap-2">
    <div class="w-8 h-8 bg-secondary rounded-lg flex items-center justify-center">
        <svelte:component this={IconBinoculars} class="w-5 h-5 text-secondary-content" />
    </div> 
    Above Me
</h1>
<div class="card bg-base-100 mb-4 w96 shadow-sm rounded-xl hover:shadow-md transition-all duration-200">
<div class="card-body">
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
            <!-- {:else if data.length === 0}
                <div class="alert alert-info">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                    <span>Nothing above you right now!</span>
                </div> -->
            {:else}
                <div class="table-container">
                <table class="table">
                    <thead class="uppercase tracking-wider">
                        <tr>
                            <th>Reg</th>
                            <th>Flight</th>
                            <th>Type</th>
                            <th>Distance (km)</th>
                            <th>Last Seen</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each data as aircraft (aircraft.registration)}
                        <tr transition:fade={{ duration: 2000 }}
                            animate:flip={{ duration: 2000 }}
                            class="hover:bg-base-300">
                            <td>{aircraft.registration}</td>
                            <td>{aircraft.flight || '-'}</td>
                            <td>{aircraft.type || '-'}</td>
                            <td><NumberFlow value={aircraft.last_seen_distance} willChange={true} respectMotionPreference={false} /></td>
                            <td>{aircraft.last_seen ? new Date(aircraft.last_seen).toLocaleString() : '-'}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
                </div>
            {/if}
        </div>
    </div>
</div>
</div>
