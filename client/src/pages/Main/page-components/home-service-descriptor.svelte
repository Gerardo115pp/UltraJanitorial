<script>
    import TxyText from '../../../TxySvelte/txy_components/TxyText.svelte';
    import TxyImage from '../../../TxySvelte/txy_components/TxyImage.svelte';

    export let image_url = 'https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png';
    export let service_name = 'ultra janitorial';
    export let service_description = 'A tool to help you keep your files clean.';

    // style variables
    export let isFliped = false;

    const classNameComposer = () => {
        let class_name = 'service-descriptor';
        if (isFliped) {
            class_name += ' flipped';
        }
        return class_name;
    };


</script>

<style>

    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        .service-descriptor {
            display: flex;
            width: 100%;
            height: 45vh;
            padding: 1vh 1vw;
            align-items: center; 
        }

        .service-descriptor.flipped {
            flex-direction: row-reverse;
        }
        
        .service-descriptor div:first-child {
            margin-right: 10%;
        }
        .service-descriptor-image {
            width: 80vw;
        }

        :global(.service-descriptor-image img) {
            width: 100%;
            clip-path: circle(40%);
        }

        .service-descriptor-content {
            margin-right: 10vw;
        }

        :global(.flipped .service-descriptor-text) {
            margin-right: 0;
            margin-left: 10vw;
        }

        :global(.service-descriptor-text) {
            font-size: 1.5em;
            color: var(--secondary-color);
            text-indent: 2em;
            text-align: center;
        }    
    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            MobileSupport            =
    =============================================*/
    
    @media only screen and (max-width: 800px) {
        .service-descriptor, .service-descriptor.flipped {
            display: flex;
            width: 100%;
            flex-direction: column;
            height: 70vh;
            padding: 1vh 1vw;
            justify-content: center;
            align-items: center; 
        }

        .service-descriptor div:first-child {
            margin-right: 0;
        }
        .service-descriptor-image {
            width: 80vw;
            margin-bottom: 5vh;
        }

        :global(.service-descriptor-image img) {
            clip-path: circle(60%);
        }

        .service-descriptor-content {
            margin-right: 0;
        }

        :global(.flipped .service-descriptor-text) {
            margin-right: 0;
            margin-left: 0;
        }

        :global(.service-descriptor-text) {
            font-size: 1.3em;
            color: var(--secondary-color);
            text-align: center;
        }    
    }
    
    /*=====  End of MobileSupport  ======*/
    
    
</style>

<div class="{classNameComposer()}">
    <div class="service-descriptor-image">
        <TxyImage fallback={image_url} content_key="{service_name}-image" />
    </div>
    <div class="service-descriptor-content">
        <TxyText
            fallback={service_description}
            content_key={service_name}
            extra_classes="service-descriptor-text"
        />
    </div>
</div>