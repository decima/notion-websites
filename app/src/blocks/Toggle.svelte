<script>
    import Toggleable from "./Toggleable.svelte";
    import RichText from "./primaries/RichText.svelte";

    export let item;
    export let expanded = false;
</script>
<Toggleable toggleable={item.has_children} color={item.block.toggle.color}
            on:open
            on:close
            on:mountedToggle
            id={item.block.id}
>
    <h4 slot="title" class="text-accent color-{item.block.toggle.color??'none'}">
        <RichText richText={item.block.toggle.rich_text}/>
    </h4>
    <slot></slot>
</Toggleable>

