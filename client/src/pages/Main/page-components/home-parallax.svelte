<script>
    import Button from '../../../components/Button.svelte';
</script>

<style>

    
    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        .home-parallax {
            --parallax-content-width: 10vw;
            --blured-background-color: #3d3d3d77;
        }

        @supports (backdrop-filter: blur(1px)) or (-webkit-backdrop-filter: blur(1px)) {
            .home-parallax {
                --blured-background-color: rgba(46, 46, 46, 0.2);
            }
        }

        @keyframes shake {
            0% {
                transform: rotateZ(0deg) scale(1);
            }
            20% {
                transform: rotateZ(-10deg);
            }
            40% {
                transform: rotateZ(10deg);
                    }
            60% {
                transform: rotateZ(-10deg);
            }
            80% {
                transform: rotateZ(10deg);
            }
            100% {
                transform: rotateZ(0deg) scale(1.1);
            }
        }
        
        @keyframes expandBackground {
            0% {
                backdrop-filter: blur(0px);
                -webkit-backdrop-filter: blur(0px);
            }
            5% {
                background-color: aqua;
                backdrop-filter: blur(1px);
                -webkit-backdrop-filter: blur(1px);
            }
            60% {
                transform: scaleX(2) scaleY(2);
                background-color: var(--blured-background-color);
            }
            80% {
                transform: scaleX(5) scaleY(8);
            }
            100% {
                transform: scaleX(5) scaleY(25);
                background-color: var(--blured-background-color);
                backdrop-filter: grayscale(1) blur(.02px);
                -webkit-backdrop-filter: grayscale(1) blur(.02px);
            }
        }

        .home-parallax {
            overflow: hidden;
            position: relative;
            display: flex;
            height: 40vh;
            background-image: url('images/ashwini-chaudhary-Iu6parQAO-U-unsplash.jpg');
            background-position: center;
            background-attachment: fixed;
            background-repeat: no-repeat;
            background-size: cover;
            margin: 5vh 0;
            justify-content: center;
            align-items: center;
            perspective: 1px;
            
        }
        
        .blur-background {
            color: rgba(0,0,0,0);
            width: 20vw;
        }

        .home-parallax .parallax-content {
            position: absolute;
            display: flex;
            width: var(--parallax-content-width);
            justify-content: center;
            align-items: center;
            z-index: 2;
        }

    
    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            DesktopOnly            =
    =============================================*/
    
    @media only screen and (min-width: 1000px) {

        .home-parallax:hover .blur-background{
            animation: expandBackground 0.2s;
            animation-fill-mode: both;
            animation-delay: 0.1s;
            transition: all 0.1s ease-in-out;
            z-index: 1;
        }

        .home-parallax:hover .parallax-content {
            animation: shake 0.4s;
            animation-fill-mode: both;
        }
    }
    
    /*=====  End of DesktopOnly  ======*/
    
    
    /*=============================================
    =            MobileOnly            =
    =============================================*/
    
    @media only screen and (max-width: 700px) {
        .home-parallax {
            height: 20vh;
            background: var(--secondary-color);
            background-position: center;
            background-attachment: unset;
            background-repeat: no-repeat;
            background-size: cover;
            perspective: none;
            --parallax-content-width: 50%;
        }
    }
    
    /*=====  End of MobileOnly  ======*/
    
    
    
    


</style>

<div class="home-parallax">
    <div class="parallax-content">
        <Button width="var(--parallax-content-width)" font_size="1.4em" label="CONTACT" padding="2vh 2vw" />
    </div>
    <div class="blur-background">a</div>
</div>