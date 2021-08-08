<script>
    import Facebook_logo from "../images/Icons/facebook.svg";
    import Twitter_logo from "../images/Icons/twitter.svg";
    import Instagram_logo from "../images/Icons/instagram.svg";
    import Whatsapp_logo from "../images/Icons/whatsapp.svg";
    import TxyText from "../TxySvelte/txy_components/TxyText.svelte";
</script>

<style>
    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        #ultra-footer {
            height: 30vh;
            width: 100%;
            background-color: var(--theme-color);
            color: var(--dark-color);
            font-size: 1.5em;
            bottom: 0;
        }

        #footer-sections {
            display: flex;
            width: 100%;
            height: 30vh;
            justify-content: space-around;
            align-items: center;
        }
        
        .footer-section {
            width: calc((100% / 3) - 6%);
            height: 27vh;
        }

        .section-title {
            font-size: 2rem;
            color: papayawhip;
        }

        .contact-socials {
            display: flex;
            width: 100%;
            margin: 5vh 0;
            justify-content: space-around;
            align-items: center;
        }

        .social-icon {
            width: 2vw;
        }
    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            MobileSupport            =
    =============================================*/
    
    @media only screen and (max-width: 800px) {
        #ultra-footer {
            height: 30vh;
            font-size: 1em;
            bottom: 0;
        }

        #footer-sections {
            display: flex;
            height: inherit;
            justify-content: space-around;
            align-items: center;
        }

        .footer-section {
            width: 25%;
            height: calc(inherit - 6vh);
            text-align: center;
        }

        .footer-section ul {
            list-style: none;
            padding: 0;
            margin: 0;
        }

        .section-title {
            text-align: center;
            font-size: 2rem;
            color: papayawhip;
        }

        .contact-socials {
            display: flex;
            height: 15vh;
            flex-direction: column;
            margin: 3vh 0 0 7vw;
            justify-content: space-between;
            align-items: flex-start;
        }

        .social-icon {
            width: 5vw;
        }
    }
    
    /*=====  End of MobileSupport  ======*/
    
        


</style>

<footer id="ultra-footer">
    <div id="footer-sections">
        <!-- Services -->
        <div class="footer-section">
            <div class="section-title">Services</div>
            <div class="section-content">
                <ul>
                    <li>
                        <TxyText
                            content_key="footer-service-i"
                            fallback="not a service"
                            isGeneralContent={true}

                        />
                    </li>
                    <li>
                        <TxyText
                            content_key="footer-service-ii"
                            fallback="not a service"
                            isGeneralContent={true}
                        />
                    </li>
                    <li>
                        <TxyText
                            content_key="footer-service-iii"
                            fallback="not a service"
                            isGeneralContent={true}
                        />
                    </li>
                    <li>
                        <TxyText
                            content_key="footer-service-iv"
                            fallback="not a service"
                            isGeneralContent={true}
                        />
                    </li>
                    <li>
                        <TxyText
                            content_key="footer-service-v"
                            fallback="not a service"
                            isGeneralContent={true}
                        />
                    </li>
                </ul>
            </div>
        </div>
        <!-- Contact section -->
        <div class="footer-section">
            <div class="section-title">Contact</div>
            <div class="section-content">
                <ul>
                    <li>
                        <TxyText
                            content_key="footer-contact-i"
                            fallback="not a contact"
                            isGeneralContent={true}
                        />
                    </li>
                    <li>
                        <TxyText
                            content_key="footer-contact-ii"
                            fallback="not a contact"
                            isGeneralContent={true}
                        />
                    </li>
                </ul>
                <div class="contact-socials">
                        <a href="https://www.facebook.com/ultrajanitorial" target="_blank" class="social-icon social-facebook">
                            <i class="fa fa-facebook">
                                {@html Facebook_logo}
                            </i>
                        </a>
                        <a href="https://twitter.com/ultrajanitorial" target="_blank" class="social-icon social-twitter">
                            <i class="fa fa-twitter">
                                {@html Twitter_logo}
                            </i>
                        </a>
                        <a href="https://www.instagram.com/ultrajanitorial" target="_blank" class="social-icon social-instagram">
                            <i class="fa fa-instagram">
                                {@html Instagram_logo}
                            </i>
                        </a>
                        <a href="https://www.whatsapp.com/ultrajanitorial" target="_blank" class="social-icon social-whatsapp">
                            <i class="fa fa-whatsapp">
                                {@html Whatsapp_logo}
                            </i>
                        </a>
                    <div />
                </div>
            </div>
        </div>
        <!-- Address section -->
        <div class="footer-section">
            <div class="section-title">Address</div>
            <div class="section-content">
                <ul>
                    <li>
                        <TxyText
                            content_key="footer-address-i"
                            fallback="not a address"
                            isGeneralContent={true}
                        />
                    </li>
                </ul>
            </div>
        </div>
    </div>
</footer>
