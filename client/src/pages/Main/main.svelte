<script>
    import Hero from './page-components/home-hero.svelte';
    import Leads from './page-components/home-leads.svelte';
    import Parallax from './page-components/home-parallax.svelte';
    import ServiceDescriptor from './page-components/home-service-descriptor.svelte';
    import Button from '../../components/Button.svelte';
    import { push } from 'svelte-spa-router';
    import { lorem } from '../../Utilitys';
    import Txy from '../../TxySvelte/Txy';
    import { onMount, onDestroy } from 'svelte';
    

    const PAGE_NAME = "main";
    
    onMount(() => {
        Txy.setPageName(PAGE_NAME);
    });

    onDestroy(() => {
        Txy.unsetPage();
    });

    const scrollToTop = () => {
        window.scrollTo(0, 0);
        push('/about');
    }
</script>

<style>
 


    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        #main-page {
            position: relative;
            min-height: 100vh;
        }

        :global(.center-section) {
            text-transform: capitalize;
            text-align: center;
            font-size: 2.6rem;
        }

        #home-services {
            margin: 0 0 8rem;
        }

        :global(.find-out-btn) {
            width: 20vw;
            margin-left: 10vw;
        }
    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            MobileStyle            =
    =============================================*/
    
    @media only screen and (max-width: 800px) {
    
        #home-services {
            margin: 0 0 4rem;
        }

        :global(.find-out-btn) {
            width: 60vw;
        }
    }
    
    /*=====  End of MobileStyle  ======*/
    
    


</style>

<div id="main-page" class="page-with-nav">
    <Hero/>
    <!--Why ultra janitorial section-->
    <Leads/>
    <Parallax/>
    <div id="home-services">
        <ServiceDescriptor 
            image_url="images/pexels-karolina-grabowska-4239146.jpg" 
            service_name="Generic service I"
            service_description="{lorem(100)}"
        />
        <ServiceDescriptor 
            image_url="images/pexels-karolina-grabowska-4239146.jpg" 
            service_name="Generic service II"
            service_description="{lorem(100)}"
            isFliped={true}
        />
        <Button onClick={scrollToTop} button_color="var(--theme-gradiant)" label="FIND OUT MORE" font_size="2em" padding="1vh 1vw" extraClasses="find-out-btn"/>
    </div>
</div>