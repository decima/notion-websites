<script>
    import {Route, Router} from "svelte-navigator";
    import Page from "./pages/Page.svelte";
</script>

<Router>


    <Route path="" component={Page} primary={false}/>
    <Route path=":subpage" component={Page} primary={false}/>
</Router>
