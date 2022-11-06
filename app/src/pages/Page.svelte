<script>
    import {onMount} from "svelte";
    import Renderer from "../blocks/Renderer.svelte";
    import '../assets/app.scss'
    import RichText from "../blocks/primaries/RichText.svelte";
    import {loadAPI, urlAPI} from "../store/api.js";
    import Breadcrumb from "./Breadcrumb.svelte";

    let counter = 0;
    let data = null;
    export let subpage = null
    onMount(async () => {
        if (subpage) {
            subpage = "/subpage/" + subpage
        }
        data = await loadAPI(subpage)
    })
</script>
{#if data}
    {#if data.breadCrumb }
        <div class="container mx-auto">
            <Breadcrumb breadcrumb={data.breadCrumb}></Breadcrumb>
        </div>
    {/if}

    <div class="container mx-auto">

        <div class="hero min-h-screen" style="background-image: url({ urlAPI('/cover') });">
            <div class="hero-overlay bg-opacity-60"></div>
            <div class="hero-content text-center text-neutral-content">
                <div class="max-w-md">
                    <h1 class="mb-5 text-5xl font-bold">
                        {#each Object.entries(data.page.page.properties) as [key, properties]}
                            {#if properties.type == "title"}
                                <RichText richText={properties.title}/>
                            {/if}
                        {/each}
                    </h1>
                </div>
            </div>
        </div>

        <article class="prose lg:prose-xl mx-auto">
            <Renderer items={data.page.blocks}/>
        </article>
    </div>
{/if}

<footer class="footer mt-10 p-10 bg-neutral text-neutral-content">
    <!--
      <div>
          <span class="footer-title">Services</span>
          <a class="link link-hover">Branding</a>
          <a class="link link-hover">Design</a>
          <a class="link link-hover">Marketing</a>
          <a class="link link-hover">Advertisement</a>
      </div>
  -->
</footer>