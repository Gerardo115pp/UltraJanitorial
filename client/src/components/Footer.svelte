<script>
    import Facebook_logo from "../images/Icons/facebook.svg";
    import Twitter_logo from "../images/Icons/twitter.svg";
    import Instagram_logo from "../images/Icons/instagram.svg";
    import Whatsapp_logo from "../images/Icons/whatsapp.svg";
</script>

<style>
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
        height: 30vh;
        justify-content: space-around;
        align-items: center;
    }

    .section-title {
        font-size: 2rem;
        color: papayawhip;
    }

    .contact-socials {
        display: flex;
        justify-content: space-around;
        align-items: center;
    }

    .social-icon {
        width: 2vw;
    }
</style>

<footer id="ultra-footer">
    <div id="footer-sections">
        <!-- Services -->
        <div class="footer-section">
            <div class="section-title">Services</div>
            <div class="section-content">
                <ul>
                    <li>Service 1</li>
                    <li>Service 2</li>
                    <li>Service 3</li>
                    <li>Service 4</li>
                    <li>Service 5</li>
                </ul>
            </div>
        </div>
        <!-- Contact section -->
        <div class="footer-section">
            <div class="section-title">Contact</div>
            <div class="section-content">
                <ul>
                    <li>+52 55 81234-5678</li>
                    <li>some.email@ultrajanitorial.com</li>
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
                        Some random address, some random city, some random
                        country, #1234.
                    </li>
                </ul>
            </div>
        </div>
    </div>
</footer>
