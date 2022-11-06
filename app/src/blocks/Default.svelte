<script>
    export let item

    function toJSON(obj) {
        return JSON.stringify(obj, null, 2)
    }

</script>
<div>
    Component <span class="font-mono bg-red-50 cursor-pointer text-xs">{item?.type ?? "ITEM"}</span> not implemented yet
    <pre class="h-52 overflow-scroll"><code>{toJSON(item)}</code></pre>
    <slot></slot>
</div>