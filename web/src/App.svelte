<script>
  import ThemeSelector from './components/ThemeSelector.svelte';
  import HighestAircraft from './components/HighestAircraft.svelte';
  import SlowestAircraft from './components/SlowestAircraft.svelte';
  import LowestAircraft from './components/LowestAircraft.svelte';
  import FastestAircraft from './components/FastestAircraft.svelte';
  import GeneralStats from './components/GeneralStats.svelte';
  import InterestingMilAircraft from './components/InterestingMilAircraft.svelte';
  import InterestingGovAircraft from './components/InterestingGovAircraft.svelte';
  import InterestingPolAircraft from './components/InterestingPolAircraft.svelte';
  import InterestingCivAircraft from './components/InterestingCivAircraft.svelte';
  import Status from './components/Status.svelte';
</script>

<div class="container max-w-7xl mx-auto p-8">
  <div class="flex justify-between items-center mb-8">
    <div class="flex items-center gap-4">
      <h1 class="text-4xl font-bold">SkyStats</h1>
      <Status />
    </div>
    <ThemeSelector />
  </div>
  <div class="grid grid-cols-1">
    <GeneralStats />
  </div>
  <div class="grid grid-cols-1 mt-10 gap-6">
    <InterestingMilAircraft />
    <InterestingGovAircraft />
    <InterestingPolAircraft />
    <InterestingCivAircraft />
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 mt-10 gap-6">
    <FastestAircraft />
    <SlowestAircraft />
    <HighestAircraft /> 
    <LowestAircraft />
  </div>
</div>