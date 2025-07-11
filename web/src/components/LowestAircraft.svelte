<script>
    import AircraftStats from './AircraftStats.svelte';
    import { IconArrowDownDashed } from '@tabler/icons-svelte';

    const columns = [
        { header: 'Reg', field: 'registration', class: 'font-mono' },
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
    endpoint="/api/stats/lowest?limit=5"
    title="Lowest Aircraft"
    {columns}
    icon={IconArrowDownDashed}
    iconColor="text-blue-500"
    iconBgColor="bg-blue-100"
/>