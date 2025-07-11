<script>
    import AircraftStats from './AircraftStats.svelte';
    import { IconArrowUpDashed } from '@tabler/icons-svelte';

    const columns = [
        { header: 'Registration', field: 'registration', class: 'font-mono' },
        { header: 'Model', field: 'type' },
        // { header: 'Flight', field: 'flight' },
        { 
            header: 'Altitude', 
            field: 'barometric_altitude',
            formatter: (value) => value ? `${value.toLocaleString()} ft` : '-'
        },
        { 
            header: 'First Seen', 
            field: 'first_seen',
            formatter: (value) => value ? new Date(value).toLocaleString() : '-'
        }
    ];
</script>

<AircraftStats 
    endpoint="/api/stats/highest?limit=5"
    title="Highest Aircraft"
    {columns}
    icon={IconArrowUpDashed}
    iconColor="text-green-500"
    iconBgColor="bg-green-100"
/>