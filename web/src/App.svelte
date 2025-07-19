<script>
  import ThemeSelector from './components/ThemeSelector.svelte';
  import MotionFastestAircraft from './components/MotionHighestAircraft.svelte';
  import MotionSlowestAircraft from './components/MotionSlowestAircraft.svelte';
  import MotionLowestAircraft from './components/MotionLowestAircraft.svelte';
  import MotionHighestAircraft from './components/MotionHighestAircraft.svelte';
  import GeneralStats from './components/GeneralStats.svelte';
  import InterestingMilAircraft from './components/InterestingMilAircraft.svelte';
  import InterestingGovAircraft from './components/InterestingGovAircraft.svelte';
  import InterestingPolAircraft from './components/InterestingPolAircraft.svelte';
  import InterestingCivAircraft from './components/InterestingCivAircraft.svelte';
  import Status from './components/Status.svelte';
  import Above from './components/Above.svelte';
  import AboveTimeline from './components/AboveTimeline.svelte';
  import RouteTopAirlines from './components/RouteTopAirlines.svelte';
  import RouteTopCountriesOrigin from './components/RouteTopCountriesOrigin.svelte';
  import RouteTopCountriesDestination from './components/RouteTopCountriesDestination.svelte';
  import RouteTopAirportsInternational from './components/RouteTopAirportsInternational.svelte';
  import RouteTopAirportsDomestic from './components/RouteTopAirportsDomestic.svelte';
  import RouteTopRoutes from './components/RouteTopRoutes.svelte';
</script>


<div class="container max-w-8xl mx-auto p-8">
  <div class="flex justify-between items-center mt-8 mb-8">
    <div class="flex items-center gap-4">
      <h1 class="text-3xl font-bold">SkyStats</h1>
      <Status />
    </div>
    <ThemeSelector />
  </div>
  <div class="grid grid-cols-1">
    <GeneralStats />
  </div>
  <div class="grid grid-cols-1 mt-10 gap-6">
    <AboveTimeline />
  </div>
  <div class="grid grid-cols-3 mt-10 gap-6">
    <RouteTopAirlines />
    <RouteTopAirportsDomestic  />
    <RouteTopAirportsInternational />
    <RouteTopRoutes />
    <RouteTopCountriesOrigin />
    <RouteTopCountriesDestination />
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 mt-10 gap-6">
    <InterestingMilAircraft />
    <InterestingGovAircraft />
    <InterestingPolAircraft />
    <InterestingCivAircraft />
  </div>
  <div class="grid grid-cols-1 md:grid-cols-2 mt-10 gap-6">
    <MotionFastestAircraft />
    <MotionSlowestAircraft />
    <MotionHighestAircraft /> 
    <MotionLowestAircraft />
  </div>
</div>