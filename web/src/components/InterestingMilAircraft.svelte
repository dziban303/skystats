<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconTank } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/military'
    title="Military Aircraft"
    icon={IconTank}
    iconColor="text-secondary-content"
    iconBgColor="bg-secondary"
    aircraftType="mil"
/>