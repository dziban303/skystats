<script>
    import MotionStats from './MotionStats.svelte';
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

<MotionStats 
    endpoint="/api/stats/motion/slowest"
    title="Slowest Aircraft"
    {columns}
    icon={IconWalk}
    iconColor="text-yellow-600"
    iconBgColor="bg-yellow-100"
/>