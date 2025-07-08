<script>
  import { onMount } from 'svelte';
  
  let slowestData = [];
  let loading = true;
  let error = null;
  
  onMount(async () => {
    try {
      console.log('Fetching data from API...');
      const response = await fetch('/api/stats/slowest');
      console.log('Response status:', response.status);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      console.log('Data received:', data);
      // Get the 10 most recent entries
      slowestData = data.slice(0, 10);
    } catch (err) {
      console.error('Fetch error:', err);
      error = err.message;
    } finally {
      loading = false;
    }
  });
</script>

<div class="card bg-base-100 shadow-xl">
  <div class="card-body">
    <h2 class="card-title text-2xl mb-4">Slowest Aircraft</h2>
    
    {#if loading}
      <div class="flex justify-center py-8">
        <span class="loading loading-spinner loading-lg"></span>
      </div>
    {:else if error}
      <div class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>Error: {error}</span>
      </div>
    {:else if slowestData.length === 0}
      <div class="alert alert-info">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>No data available</span>
      </div>
    {:else}
      <div class="overflow-x-auto">
        <table class="table table-zebra">
          <thead>
            <tr>
              <th>Registration</th>
              <th>Type</th>
              <th>Flight</th>
              <th>Ground Speed (kt)</th>
              <th>First Seen</th>
              <th>Last Seen</th>
            </tr>
          </thead>
          <tbody>
            {#each slowestData as aircraft}
              <tr>
                <td class="font-mono">{aircraft.registration}</td>
                <td>{aircraft.type}</td>
                <td>{aircraft.flight || '-'}</td>
                <td class="text-center">{aircraft.ground_speed}</td>
                <td>{new Date(aircraft.first_seen).toLocaleString()}</td>
                <td>{new Date(aircraft.last_seen).toLocaleString()}</td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  </div>
</div>