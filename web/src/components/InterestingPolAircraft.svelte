<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconShield } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/police?limit=5'
    title="Police Aircraft"
    icon={IconShield}
    iconColor="text-secondary-content"
    iconBgColor="bg-secondary"
    aircraftType="pol"
/>
    <!-- iconColor="text-blue-500"
    iconBgColor="bg-blue-100" -->