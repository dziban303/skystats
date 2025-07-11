<script>
    import AircraftStats from './AircraftStats.svelte';
    import { IconWalk } from '@tabler/icons-svelte';

    const columns = [
        { header: 'Reg', field: 'registration', class: 'font-mono' },
        { header: 'Type', field: 'type' },
        // { header: 'Flight', field: 'flight' },
        { 
            header: 'Speed', 
            field: 'ground_speed',
            formatter: (value) => value ? `${value.toLocaleString()} kts` : '-'
        },
        { 
            header: 'First Seen', 
            field: 'first_seen',
            formatter: (value) => value ? new Date(value).toLocaleString() : '-'
        }
    ];
</script>

<AircraftStats 
    endpoint="/api/stats/slowest?limit=5"
    title="Slowest Aircraft"
    {columns}
    icon={IconWalk}
    iconColor="text-yellow-600"
    iconBgColor="bg-yellow-100"
/>