<script>
    export let field;
    export let properties;
    export let value;
    export let show = false;
</script>

<span style="cursor: pointer" on:click={()=>show = !show}>&nbsp;</span>
{#if show}
    <div>{field}</div>
    {JSON.stringify(properties)}
{/if}