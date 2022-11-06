<script>
    import RichText from "../blocks/primaries/RichText.svelte";

    export let breadcrumb = []
</script>
<div class="text-sm breadcrumbs">
    <ul>
        <li><a href="/"><i class="fas fa-house"></i></a></li>
        {#each breadcrumb as item }
            <li>
                <a href="/{item.id}">
                    <RichText richText={item.properties.title.title}/>
                </a>
            </li>
        {/each}
    </ul>
</div>