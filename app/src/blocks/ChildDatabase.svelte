<script>
    import RichText from "./primaries/RichText.svelte";
    import {resolveField} from "./databases/fields.js";
    import {storeInDB} from "../store/api.js";

    export let item = {}
    let values = {}
    let save = async function () {
        await storeInDB(item.block.id, values)
        values={}
    }
</script>

<fieldset>
    <legend class="text-2xl border-b w-full">
        <RichText richText={item.database.title}/>
    </legend>
    <form on:submit={save} action="#">
        <div class="grid grid-col-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
            {#each Object.entries(item.database.properties) as [field, properties] }
                {#if !field.startsWith("_")}
                    <div>

                        <svelte:component this="{resolveField(properties.type)}" bind:value={values[field]}
                                          field={field}
                                          properties={properties}/>
                    </div>
                {/if}

            {/each}
        </div>
        <button type="submit" class="btn btn-primary btn-wide">Send</button>
    </form>
</fieldset>