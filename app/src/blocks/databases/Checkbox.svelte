<script>
    export let field;
    export let properties;
    export let value
</script>
<div class="form-control">
    <label class="label">
        <span class="label-text">{field}</span>
        <input type="checkbox" bind:checked={value} class="checkbox"/>
    </label>

</div>