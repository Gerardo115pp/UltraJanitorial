<script>
    import email_icon from '../../images/Icons/email.svg';
    import phone_icon from '../../images/Icons/Phone.svg';
    import address_icon from '../../images/Icons/AddressLocation.svg';
    import hours_icon from '../../images/Icons/WorkingHours.svg';
</script>

<style>
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
        width: 2vw;
        fill: var(--secondary-color);
    }
</style>

<div id="contact-page">
    <div id="contact-card">
        <div id="contact-title">
            <h2>Contact Information</h2>
        </div>
        <div id="contact-information">
            <div class="contact-field" id="contact-email">
                <i class="contact-icon">
                    {@html email_icon}
                </i>
                <span class="contact-field-data">
                    contact@ultrajanitorial.com
                </span>
            </div>
            <div class="contact-field" id="contact-phone">
                <i class="contact-icon">
                    {@html phone_icon}
                </i>
                <span class="contact-field-data">
                    (555) 555-5555
                </span>
            </div>
            <div class="contact-field" id="contact-address">
                <i class="contact-icon">
                    {@html address_icon}
                </i>
                <span class="contact-field-data">
                    123 Fake Street, Fake City, FAKE ST, FAKE CITY, FAKE STATE, FAKE COUNTRY
                </span>
            </div>
            <div class="contact-field" id="contact-hours">
                <i class="contact-icon">
                    {@html hours_icon}
                </i>
                <span class="contact-field-data">
                    Monday - Friday: 9:00 AM - 5:00 PM
                </span>
            </div>
        </div>
    </div>
</div>