<script>
    import NavBar from './components/ultra-navbar.svelte';
    import Router from 'svelte-spa-router';
    import { routes } from './routes';
</script>

<style>
    @import url(//db.onlinewebfonts.com/c/58aaccbcd294a88d9b91af291895060d?family=GatsbyFLF);
    @import url('https://fonts.googleapis.com/css2?family=Abril+Fatface&display=swap');

    :global(:root) {
        --theme-color: #FF9900;
        --main-font: GatsbyFLF;
        --title-font: Abril Fatface, sans-serif;
    }    

    /*=============================================
    =            Normalization            =
    =============================================*/
    
    :global(html) {
        font-size: 16px;
    }

    :global(body) {
        font-family: var(--main-font), sans-serif;
        margin: 0;
    }
    
    /*=====  End of Normalization  ======*/
    
    :global(.title-text) {
        font-weight: lighter;
        font-family: var(--title-font), sans-serif;
    }
    
    #main-component {
        color: var(--theme-color);
    }
</style>

<svelte:head>
    <link type="text/css" rel="stylesheet" href="normalize.css" />
</svelte:head>

<div id="main-component">
    <NavBar/>
    <Router {routes} />
</div>
