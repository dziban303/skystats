<script>
    import { onMount, onDestroy } from 'svelte'

    export let endpoint;
    export let title;
    export let icon;
    export let iconColor;
    export let iconBgColor;
    export let aircraftType;

    let refreshRate = 10000
    let data = [];
    let loading = true;
    let error = null;
    let interval = null;
    let selectedAircraft = null;

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

    function showAircraftModal(aircraft) {
        selectedAircraft = aircraft;
        // @ts-ignore
        document.getElementById(aircraftType).showModal();
    }

    function closeModal() {
        selectedAircraft = null;
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

<!-- <div class="card bg-base-100 w96 shadow-sm rounded-xl border border-secondary hover:shadow-md transition-all duration-200"> -->
<h1 class="card-title -mb-2 flex items-center gap-2">
    {#if icon}
        <div class="w-8 h-8 {iconBgColor} rounded-lg flex items-center justify-center">
            <svelte:component this={icon} class="w-5 h-5 {iconColor}" />
        </div> 
    {/if}
    {title}
</h1>
<div class="card bg-base-100 mb-4 w96 shadow-sm rounded-xl hover:shadow-md transition-all duration-200">
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
                    <thead class="uppercase tracking-wider">
                        <tr>
                            <th>Reg</th>
                            <th>Flight</th>
                            <th>Operator</th>
                            <th>Type</th>
                            <th>Category</th>
                            <!-- <th>Tags</th> -->
                            <th>Last Seen</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each data as aircraft}
                        <tr class="hover:bg-base-300 cursor-pointer" on:click={() => showAircraftModal(aircraft)}>
                            <td>{aircraft.registration}</td>
                            <td>{aircraft.flight || '-'}</td>
                            <td>{aircraft.operator}</td>
                            <td>{aircraft.type}</td>
                            <td>{aircraft.category}</td>
                            <td>{aircraft.seen ? new Date(aircraft.seen).toLocaleString() : '-'}</td>
                        </tr>
                        {/each}
                    </tbody>
                </table>
            {/if}
        </div>
    </div>
</div>

<dialog id={aircraftType} class="modal" on:close={closeModal}>
    <div class="modal-box max-w-4xl">
        {#if selectedAircraft}
            <div class="flex items-center justify-between mb-2">
                <h3 class="text-lg font-bold">{selectedAircraft.registration} - {selectedAircraft.type}</h3>
                <div class="flex gap-2">
                    {#if selectedAircraft.tag1}
                        <div class="badge badge-accent text-white">{selectedAircraft.tag1}</div>
                    {/if}
                     {#if selectedAircraft.tag2}
                        <div class="badge badge-accent text-white">{selectedAircraft.tag2}</div>
                    {/if}
                    {#if selectedAircraft.tag3}
                        <div class="badge badge-accent text-white">{selectedAircraft.tag3}</div>
                    {/if}
                </div>
            </div>
            <p class="text-sm text-gray-600 mb-4">{selectedAircraft.operator}</p>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {#if selectedAircraft.image_link_1}
                    <img src={selectedAircraft.image_link_1} alt="{selectedAircraft.registration} photo 1" class="w-full h-auto rounded-lg" />
                {/if}
                {#if selectedAircraft.image_link_2}
                    <img src={selectedAircraft.image_link_2} alt="{selectedAircraft.registration} photo 2" class="w-full h-auto rounded-lg" />
                {/if}
                {#if selectedAircraft.image_link_3}
                    <img src={selectedAircraft.image_link_3} alt="{selectedAircraft.registration} photo 3" class="w-full h-auto rounded-lg" />
                {/if}
            </div>
            {#if !selectedAircraft.image_link_1 && !selectedAircraft.image_link_2 && !selectedAircraft.image_link_3}
                <p class="text-center text-gray-500 py-8">No photos available for this aircraft</p>
            {/if}
        {/if}
        <div class="modal-action">
            <form method="dialog">
                <button class="btn">Close</button>
            </form>
        </div>
    </div>
    <form method="dialog" class="modal-backdrop">
        <button>close</button>
    </form>
</dialog>
