<script>
    import Toggleable from "./Toggleable.svelte";
    import RichText from "./primaries/RichText.svelte";

    export let item;
    export let expanded = false;
</script>
<Toggleable toggleable={item.has_children} color={item.block.heading_1.color??'none'}
            on:open
            on:close
            on:mountedToggle
            id={item.block.id}

>
    <h1 slot="title"
        className=" text-xl sm:text-2xl p-1 sm:p-2 md:text-3xl xl:text-4xl 2xl:text-5xl 2xl:p-4 color-{item.block.heading_1.color??'none'} font-semibold">
        <RichText richText={item.block.heading_1.rich_text}/>
    </h1>
    <slot></slot>
</Toggleable>