<script>
    export let field;
    export let properties;
    export let value
</script>
<div class="form-control">
    <label class="label">
        <span class="label-text">{field}</span>
    </label>
    <input type="number" bind:value  placeholder="{field}" class="input input-bordered"/>
</div>