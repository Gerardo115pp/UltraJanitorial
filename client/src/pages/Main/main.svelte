<script>
    import Hero from './page-components/home-hero.svelte';
    import Leads from './page-components/home-leads.svelte';
    import Parallax from './page-components/home-parallax.svelte';
    import ServiceDescriptor from './page-components/home-service-descriptor.svelte';
    import Button from '../../components/Button.svelte';
    import { push } from 'svelte-spa-router';

    // a function that generates a lorem ipsum string of a given length
    const lorem = (length) => {
        let result = "";
        const words = ["lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore", "magna", "aliqua."];
        for (let i = 0; i < length; i++) {
            result += result.length > 1 ? words[Math.floor(Math.random() * words.length)] : words[0];
            if (i !== length - 1) {
                result += " ";
            }
        }
        return result;
    };
</script>

<style>
 
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
        margin-left: 10vw;
    }
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
        <Button onClick={() => push("/contact")} button_color="var(--theme-gradiant)" width="20vw" label="FIND OUT MORE" font_size="2em" padding="1vh 1vw" extraClasses="find-out-btn"/>
    </div>
</div>