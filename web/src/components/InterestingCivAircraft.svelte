<script>
    import InterestingAircraft from "./InterestingAircraft.svelte";
    import { IconCurrencyDollar } from '@tabler/icons-svelte';
</script>

<InterestingAircraft
    endpoint='api/stats/interesting/civilian?limit=5'
    title="Civilian Aircraft"
    icon={IconCurrencyDollar}
    iconColor="text-pink-500"
    iconBgColor="bg-pink-100"
/>