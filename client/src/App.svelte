<script>
    import Router from 'svelte-spa-router';
    import { routes } from './routes';
</script>

<style>
    @import url(//db.onlinewebfonts.com/c/58aaccbcd294a88d9b91af291895060d?family=GatsbyFLF);

    :global(:root) {
        --theme-color: #FF9900;
        --theme-font: GatsbyFLF
    }    

    /*=============================================
    =            Normalization            =
    =============================================*/
    
    :global() {
        margin: 0;
    }

    :global(body) {
        font-family: var(--theme-font), sans-serif;
        margin: 0;
    }
    
    /*=====  End of Normalization  ======*/
    
    
    #main-component {
        font-family: var(--theme-font);
        color: var(--theme-color);
    }
</style>

<svelte:head>
    <link type="text/css" rel="stylesheet" href="normalize.css" />
</svelte:head>

<div id="main-component">
    <Router {routes} />
</div>
