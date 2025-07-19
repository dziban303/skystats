<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconBuildings } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/government'
    title="Government Aircraft"
    icon={IconBuildings}
    iconColor="text-secondary-content"
    iconBgColor="bg-secondary"
    aircraftType="gov"
/>
    