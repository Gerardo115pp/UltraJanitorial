<script>
    import Txy from '../../../TxySvelte/Txy';
    import TxyText from '../../../TxySvelte/txy_components/TxyText.svelte';
    import TxyImage from '../../../TxySvelte/txy_components/TxyImage.svelte';


    // content propts
    export let name = "generic";
    export let description = "This card should have a description";
    export let image = "https://www.example.com/image.png";

    // style props
    export let isTitleFlipped = false;
    export let isContentFlipped = false;
    export let hasDivisor = false;


    // compose className base on props
    const composeClassName = () => {
        let className = "about-card ";
        if (isTitleFlipped) className += "title-flipped ";
        if (isContentFlipped) className += "content-flipped ";
        if (hasDivisor) className += "has-divisor ";
        return className;
    }
</script>

<style>

    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        .about-card {
            margin: 0 0 3vh;
        }

        .about-card-header {
            display: flex;
            font-size: 1.5em;
            padding: 1vh 2vw;
            text-transform: uppercase;
            align-items: center;
            margin-bottom: 4vh;
        }
        
        .title-flipped .about-card-header {
            flex-direction: row-reverse;
        }

        .about-card-divisor {
            width: 30vw;
            border-bottom: 3px solid var(--theme-color);
            margin: 0 2vw;
        }

        .about-card-content {
            display: flex;
            justify-content: space-around;
            align-items: center;
        }

        .content-flipped .about-card-content {
            flex-direction: row-reverse;
        }

        .about-card-description {
            max-width: 60%;
        }
        
        :global(.about-card-description p.about-card-description-txt) {
            width: 80%;
            display: flex;
            font-size: 1.8em;
            justify-content: flex-end;
            align-items: center;
        }

        .about-card-image {
            width:  30vw;
        }

        :global(.about-card-image img) {
            width: 100%;
            border-radius: 5px;
        }
    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            MobileSupport            =
    =============================================*/
    
    @media only screen and (max-width: 800px) {
        
        .about-card-header {
            font-size: 1.1em;
            padding: 1vh 4vw;
        }

        .about-card-divisor {
            width: 50vw;
        }

        .about-card-content {
            flex-direction: column-reverse;
            justify-content: center;
            align-items: center;
        }

        .content-flipped .about-card-content {
            flex-direction: column-reverse;
        }


        .about-card-description {
            max-width: 100%;
        }
        
        :global(.about-card-description p.about-card-description-txt) {
            width: 90%;
            text-align: center;
            justify-content: center;
            margin: 2vh auto;
        }

        .about-card-image {
            width: 80%;
        }
    }
    
    /*=====  End of MobileSupport  ======*/
    
        
    
</style>

<div id="about-{name}" class="{composeClassName()}">
    <div class="about-card-header">
        <h1 class="about-card-title">{name}</h1>
        {#if hasDivisor}
            <div class="about-card-divisor"></div>
        {/if}
    </div>
    <div class="about-card-content">
        <div class="about-card-description">
            <TxyText
                content_key={name.replace(/\s/g, '_')}
                fallback="{description}"
                extra_classes="about-card-description-txt"
            />
        </div>
        <div class="about-card-image">
            <TxyImage 
                content_key={name.replace(/\s/g, '_')+'_image'}
                fallback="{image}"
            />
        </div>
    </div>
</div>