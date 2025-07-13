<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconBuildings } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/government?limit=5'
    title="Government Aircraft"
    icon={IconBuildings}
    iconColor="text-secondary-content"
    iconBgColor="bg-secondary"
    aircraftType="gov"
/>
    <!-- iconColor="text-green-500"
    iconBgColor="bg-green-100" -->
    