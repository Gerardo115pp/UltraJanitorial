<script>
    import AboutCard from './page-components/about-card.svelte';
    import { onMount, onDestroy } from 'svelte';
    import { lorem } from '../../Utilitys'; 
    import Txy from '../../TxySvelte/Txy';
    

    const PAGE_NAME = "about";
    
    onMount(() => {
        Txy.setPageName(PAGE_NAME);
    });

    onDestroy(() => {
        Txy.unsetPage();
    });

    const isMobile = window.screen.width < 800;
</script>

<style>
    #about-page {
        padding-top: 6vh;
        padding-bottom: 5vh;
    }
</style>

<div id="about-page">
    {#if !isMobile}
        <AboutCard 
            name="Our team"
            description="Our team consists of the most proactive people who have the ability to work independently with minimal supervision and provide excellent customer service. We have the knowledge of health and safety standards to handle cleaning chemicals and equipment safely. We are dedicated to maintaining a schedule while responding to incidents.We believe that the most productive and happy employees are well taken care of, and we are dedicated to offering fair compensation for quality work."
            image="images/pexels-alexander-suhorucov-6457579.jpg"
        />
    {:else}
        <AboutCard
            name="Our team"
            hasDivisor={true}
            isTitleFlipped={true}
            description="Our team consists of the most proactive people who have the ability to work independently with minimal supervision and provide excellent customer service. We have the knowledge of health and safety standards to handle cleaning chemicals and equipment safely. We are dedicated to maintaining a schedule while responding to incidents.We believe that the most productive and happy employees are well taken care of, and we are dedicated to offering fair compensation for quality work."
            image="images/pexels-alexander-suhorucov-6457579.jpg"
        />        
    {/if}
    <AboutCard
        name="Our Mission"
        description={lorem(50)}
        image="images/pexels-alfauzikri-325135.jpg"
        hasDivisor={true}
        isContentFlipped={true}
    />
    <AboutCard
        name="Our Vision"
        description="Our vision is to create a safe and healthy workplace environment by providing outstanding quality cleaning services."
        image="images/pexels-alexander-suhorucov-6457513.jpg"
        hasDivisor={true}
        isTitleFlipped={true}
    />
</div>