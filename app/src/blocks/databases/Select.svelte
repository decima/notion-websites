<script>
    export let field;
    export let properties;
    export let value;
</script>
<div class="form-control">
    <label class="label">
        <span class="label-text">{field}</span>
    </label>
    <select bind:value class="select select-bordered">
        <option disabled selected>{field}</option>
        {#each properties.select.options as option}
            <option value={option.id}>{option.name}</option>
        {/each}
    </select>
</div>