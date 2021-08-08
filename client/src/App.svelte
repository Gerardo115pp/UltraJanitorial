<script>
    import TxyContent from './TxySvelte/TxyContentGetter.svelte';
    import NavBar from './components/ultra-navbar.svelte';
    import Footer from './components/Footer.svelte';
    import Router from 'svelte-spa-router';
    import { routes } from './routes';
</script>

<style>
    @import url('https://fonts.googleapis.com/css2?family=Limelight&display=swap');

    @font-face {
        font-family: 'GatsbyFLF';
        src: url('fonts/58aaccbcd294a88d9b91af291895060d.woff2');
    }

    :global(:root) {
        --theme-color: #FF9900;
        --secondary-color: #009CBE;
        --dark-color: #D8790A;
        --theme-gradiant: linear-gradient(309deg, rgba(253,190,0,1) 0%, rgba(255,153,0,1) 100%);
        --main-font: GatsbyFLF;
        --title-font: Limelight, sans-serif;
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
    
    
    /*=============================================
    =            Responsive Rules            =
    =============================================*/
    
    @media screen and (max-width: 900px) {
        #main-component {
            overflow-x: hidden;
            max-width: 100vw;
        }
    }
    
    
    /*=====  End of Responsive Rules  ======*/
    
    

    :global(.title-text) {
        font-weight: lighter;
        font-family: var(--title-font), sans-serif;
    }
    
    #main-component {
        color: var(--theme-color);
    }

    #content-wrap {
        min-height: 63vh;
    }

</style>

<svelte:head>
    <link type="text/css" rel="stylesheet" href="normalize.css" />
</svelte:head>

<TxyContent font_family="var(--main-font)" theme_color="var(--theme-color)"/>
<div id="main-component">
    <NavBar/>
    <div id="content-wrap">
        <Router {routes} />
    </div>
    <Footer/>
</div>
