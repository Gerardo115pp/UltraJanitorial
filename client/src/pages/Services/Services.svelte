<script>
    import CleaningStep from './page-components/cleaing-step.svelte';
    import Button from '../../components/Button.svelte';
    import { push } from 'svelte-spa-router';
    import { lorem } from '../../Utilitys';
    import step_logo_I from '../../images/Icons/CleaningI.svg';
    import step_logo_II from '../../images/Icons/CleaningII.svg';
    import step_logo_III from '../../images/Icons/CleaningIII.svg';

    let contact_button = null;

    const handleButtonClick = () => {
        if (contact_button) {
            contact_button.classList.add('accelerable-btn');
            // wait 500ms before pushing contact route
            setTimeout(() => {
                contact_button.classList.remove('accelerable-btn');
                push('/contact');
            }, 400);
        }
    }

</script>

<style>
    @keyframes accelerate {
        0% {
            transform: skewX(30deg) translateX(-2em);
        }
        10% {
            transform: skewX(0) translateX(5em);
        }
        30% {
            opacity: .8;
        }
        100% {
            transform: translateX(80vw);
            opacity: 0;
        }
    }

    #services-header {
        height: 30vh;
        padding: 5vh 10vh;
    }

    #sh-title {
        background-image: linear-gradient(90deg, #D8790A 0%, #FF9900 30%);
        background-clip: text;
        -webkit-background-clip: text;
        font-size: 3.5em;
        text-transform: uppercase;
        -webkit-text-fill-color: transparent;
        -moz-text-fill-color: transparent;
    }

    #sh-sub-title {
        font-size: 1.9em;
        color: var(--secondary-color);
        text-indent: .1em;
    }
    
    #header-contact-btn-wrapper {
        margin-top: 5vh;
    }

    :global(.accelerable-btn) {
        animation: accelerate .5s !important;
        animation-iteration-count: 1 !important;
        animation-fill-mode: forwards !important;
    }

    #service-information {
        height: 110vh;
        background-color: papayawhip;
    }

    #service-callout {
        height: 30vh;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0 30vh;
    }

    #service-callout-title {
        font-family: var(--title-font);
        font-size: 1.5em;
    }

    #service-callout-description {
        width: 60%;
        font-size: 1.4em;
        color: var(--secondary-color);
        text-align: center;
    }

    #how-we-work {
        margin: 6vh 0 4vh;
    }

    #how-we-work-title {
        text-align: center;
        font-size: 1.5em;
    }

    #hww-divisor {
        display: flex;
        justify-content: center;
    }

    #hww-divisor-squares {
        display: flex;
        width: 8vw;
        margin: 2vh 0;
        justify-content: space-around;
    }

    .hww-divisor-square {
        background-color: var(--theme-color);
        width: 1em;
        height: 1em;
        transform: rotate(45deg);
    }

    #hww-steps {
        height: 50vh;
        display: flex;
        padding: 5vh 5vw;
        align-items: center;
        justify-content: space-around;
    }
</style>

<div id="services-page">
    <div id="services-header">
        <div id="sh-content-wraper">
            <div id="header-title">
                <div id="sh-title">
                    PROFESIONAL CLEAING
                </div>
                <div id="sh-sub-title">
                    and a promise of satisfaction
                </div>
            </div>
            <div id="header-contact-btn-wrapper">
                <div id="sh-contact-btn">
                    <!-- Contact Button -->
                    <Button 
                        bind:self_ref={contact_button}
                        onClick={handleButtonClick}
                        label="CONTACT"
                        font_size="2em" 
                        button_color="var(--theme-gradiant)" 
                        padding="1vh 1vw"
                        width="10vw"
                    />
                </div>
            </div>
        </div>
    </div>
    <div id="service-information">
        <div id="service-callout">
            <div id="service-callout-title">
                <h2>Efective,</h2>
                <h2>Eco-Friendly</h2>
                <h2>Cleaning</h2>
            </div>
            <div id="service-callout-description">
                <p>
                    We provide a variety of services to help you clean your buissnes, from small home cleaning to full-service floor cleaning.
                    {lorem(100)}
                </p>
            </div>
        </div>
        <div id="how-we-work">
            <div id="how-we-work-title">
                <h2>HOW WE WORK</h2>
            </div>
            <div id="hww-divisor">
                <div id="hww-divisor-squares">
                    <div class="hww-divisor-square"></div>
                    <div class="hww-divisor-square"></div>
                    <div class="hww-divisor-square"></div>
                </div>
           </div>
           <div id="hww-steps">
               <CleaningStep 
                    step_icon={step_logo_I}
                    step_name="Cleaing step I"
                    step_description="{lorem(40)}"
                    width="20vw"
                />
               <CleaningStep 
                    step_icon={step_logo_II}
                    step_name="Cleaing step II"
                    step_description="{lorem(40)}"
                    width="20vw"
                />
               <CleaningStep 
                    step_icon={step_logo_III}
                    step_name="Cleaing step I"
                    step_description="{lorem(40)}"
                    width="20vw"
                />
           </div>
        </div>
    </div>
</div>