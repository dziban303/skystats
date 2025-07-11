<script>
  import ThemeSelector from './components/ThemeSelector.svelte';
  import HighestAircraft from './components/HighestAircraft.svelte';
  import SlowestAircraft from './components/SlowestAircraft.svelte';
  import LowestAircraft from './components/LowestAircraft.svelte';
  import FastestAircraft from './components/FastestAircraft.svelte';
  import GeneralStats from './components/GeneralStats.svelte';
</script>

<ThemeSelector />

<div class="container mx-auto p-8">
  <h1 class="text-4xl font-bold mb-8">Sky Stats</h1>
    <GeneralStats />
  <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
    <FastestAircraft />
    <SlowestAircraft />
    <HighestAircraft /> 
    <LowestAircraft />
  </div>
</div>