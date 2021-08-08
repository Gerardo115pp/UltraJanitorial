<script>
    import email_icon from '../../images/Icons/email.svg';
    import phone_icon from '../../images/Icons/Phone.svg';
    import address_icon from '../../images/Icons/AddressLocation.svg';
    import hours_icon from '../../images/Icons/WorkingHours.svg';
    import { onMount, onDestroy } from 'svelte';
    import Txy from '../../TxySvelte/Txy';
    import TxyText from '../../TxySvelte/txy_components/TxyText.svelte';
    
    const PAGE_NAME = 'contact';

    onMount(() => {
        Txy.setPageName(PAGE_NAME); 
    });

    onDestroy(() => {
        Txy.unsetPage();
    });

</script>

<style>
    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        #contact-page {
            display: flex;
            height: 63vh;
            justify-content: center;
            align-items: center;
        }

        #contact-card {
            width: 72%;
            height: 35vh;
            border-radius: 3px;
            border: 1px solid var(--theme-color);
            box-shadow: 0 10px  15px 0 rgba(0, 0, 0, 0.2);
        }

        #contact-title {
            display: flex;
            width: 100%;
            height: 10vh;
            font-size: 1.4rem;
            text-transform: uppercase;
            justify-content: center;
            align-items: center;
        }

        #contact-information {
            height: 25vh;
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            grid-column-gap: 10vh;
            padding: 0 5vw;
        }

        .contact-field {
            display: flex;
            font-size: 1.3em;
            align-items: center;
            justify-content: space-between;
        }

        .contact-icon{
            width: 8%;
            fill: var(--secondary-color);
        }
    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            MobileSupport            =
    =============================================*/
    
    @media only screen and (max-width: 800px) {
        #contact-page {
            height: 65vh;
        }

        #contact-card {
            width: 90%;
            height: 45vh;
        }

        #contact-information {
            grid-template-columns: repeat(1, 1fr);

        }

        .contact-field {
            justify-content: flex-start;
        }

        .contact-icon {
            margin-right: 3%;
        }
    }

    
    /*=====  End of MobileSupport  ======*/
    
    
    
</style>

<div id="contact-page">
    <div id="contact-card">
        <div id="contact-title">
            <TxyText
                content_key="contact-title"
                fallback="Contact"
                isHeader={true}
            />
        </div>
        <div id="contact-information">
            <div class="contact-field" id="contact-email">
                <i class="contact-icon">
                    {@html email_icon}
                </i>
                <span class="contact-field-data">
                    <TxyText
                        content_key="contact-email"
                        fallback="Email"
                    />
                </span>
            </div>
            <div class="contact-field" id="contact-phone">
                <i class="contact-icon">
                    {@html phone_icon}
                </i>
                <span class="contact-field-data">
                    <TxyText
                        content_key="contact-phone"
                        fallback="Phone"
                    />
                </span>
            </div>
            <div class="contact-field" id="contact-address">
                <i class="contact-icon">
                    {@html address_icon}
                </i>
                <span class="contact-field-data">
                    <TxyText
                        content_key="contact-address"
                        fallback="Address"
                    />
                </span>
            </div>
            <div class="contact-field" id="contact-hours">
                <i class="contact-icon">
                    {@html hours_icon}
                </i>
                <span class="contact-field-data">
                    <TxyText
                        content_key="contact-hours"
                        fallback="Hours"
                    />
                </span>
            </div>
        </div>
    </div>
</div>