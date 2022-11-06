<script>
    import Default from "./Default.svelte";

    export let item;
</script>

<a class="btn btn-primary" href="{item.block.id}">{item.block.child_page.title}</a>
