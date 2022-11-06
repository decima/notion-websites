<script>
    export let field;
    export let properties;
    export let value = []
    let selectedChecks = {}

    function onSelect(id) {
        return function (event) {
            selectedChecks[id] = event.target.checked;

            value = [];
            for (let [k, v] of Object.entries(selectedChecks)) {
                if (v) {
                    value.push(k)
                }
            }
        }

    }
</script>
<div class="form-control">
    <label class="label">
        <span class="label-text">{field}</span>
    </label>
    {#each properties.multi_select.options as option,key}
        <label class="label">
            <span class="label-text">{option.name}</span>
            <input type="checkbox" class="checkbox" on:change={onSelect(option.id)} bind:checked={selectedChecks[option.id]}>
        </label>
    {/each}
</div>
