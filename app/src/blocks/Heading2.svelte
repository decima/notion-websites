<script>
    import Toggleable from "./Toggleable.svelte";
    import RichText from "./primaries/RichText.svelte";

    export let item;
    export let expanded = false;
</script>
<Toggleable toggleable={item.has_children} color={item.block.heading_2.color??'none'}
            on:open
            on:close
            on:mountedToggle
            id={item.block.id}
>
    <h2 slot="title" class=" text:lg sm:text-xl p-1 sm:p-2 md:text-2xl xl:text-3xl 2xl:text-4xl 2xl:p-4 text-secondary color-{item.block.heading_2.color??'none'}  font-semibold">
        <RichText richText={item.block.heading_2.rich_text}/>
    </h2>
    <slot></slot>
</Toggleable>