<script>
    export let field;
    export let properties;
    export let value
</script>

<div class="form-control  w-full max-w-xs">
    <label class="label">
        <span class="label-text">{field}</span>
    </label>
    <input type="text" bind:value placeholder={field} class="input input-bordered"/>
</div>