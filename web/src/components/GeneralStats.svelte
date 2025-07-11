<script>
    import { onMount, onDestroy } from 'svelte';
    import { IconPlaneTilt, IconCalendar, IconClock, IconCategoryPlus } from '@tabler/icons-svelte';


    let data = {};
    let endpoint = '/api/stats/general'
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
            console.log(data)
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


    <div class="stats stats-vertical lg:stats-horizontal shadow">
        <div class="stat">
            <div class="stat-figure">
                <div class="icon icon-tabler icons-tabler-outline">
                    <IconClock stroke={2}/>
                </div>
            </div>
            <div class="stat-title">Past Hour</div>
            <div class="stat-value">{data.hour_aircraft.toLocaleString()}</div>
            <div class="stat-desc">past hour</div>
        </div>
        <div class="stat">
            <div class="stat-figure">
                <div class="icon icon-tabler icons-tabler-outline">
                    <IconCalendar stroke={2}/>
                </div>
            </div>
            <div class="stat-title">Today</div>
            <div class="stat-value">{data.today_aircraft.toLocaleString()}</div>
            <div class="stat-desc">today</div>
        </div>
        <div class="stat">
            <div class="stat-figure">
                <div class="icon icon-tabler icons-tabler-outline">
                    <IconPlaneTilt stroke={2}/>
                </div>
            </div>
            <div class="stat-title">All Time</div>
            <div class="stat-value">{data.total_aircraft.toLocaleString()}</div>
            <div class="stat-desc">all time</div>
        </div>
        


        <!-- <div class="stat">
            <div class="stat-title">Unique Types</div>
            <div class="stat-value">{data.unique_aircraft_types.toLocaleString()}</div>
            <div class="stat-desc">↘︎ 90 (14%)</div>
        </div>
        <div class="stat">
            <div class="stat-title">Interesting Aircraft</div>
            <div class="stat-value">{data.interesting_aircraft_count.toLocaleString()}</div>
            <div class="stat-desc">↘︎ 90 (14%)</div>
        </div> -->
    </div>
{/if}