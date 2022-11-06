

<script>
    import RichText from "./primaries/RichText.svelte";
    import Icon from "./primaries/Icon.svelte";

    export let item;
</script>
<div class="alert color-{item.block.callout.color} my-2 w-full">
    <div class="w-full">

        <Icon icon={item.block.callout.icon}></Icon>
        <div class="w-full">
            <RichText richText={item.block.callout.rich_text}/>
            <slot/>
        </div>
    </div>
</div>