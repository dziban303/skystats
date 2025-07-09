<script>
    import AircraftStats from './AircraftStats.svelte';

    const columns = [
        { header: 'Registration', field: 'registration', class: 'font-mono' },
        { header: 'Type', field: 'type' },
        { header: 'Flight', field: 'flight' },
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
    endpoint="/api/stats/fastest"
    title="Fastest Aircraft"
    {columns}
/>