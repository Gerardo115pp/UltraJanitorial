<script>
    import ultra_janitorial_logo from '../images/logo.svg';

    let self = null;

    const addSticky = () => {
        if (self) {
            const isAtBottom = window.pageYOffset > self.offsetTop;
            if (isAtBottom) {
                if (!self.classList.contains('sticky')) {
                    self.classList.add('sticky');
                }
            } else {
                if (self.classList.contains('sticky')) {
                    self.classList.remove('sticky');
                }
            }
        }
    };

    window.onscroll = addSticky
</script>

<style>
    
    #uj-navbar {
        --navbar-height: 7vh;
    }

    /*=============================================
    =            BaseStyle            =
    =============================================*/
    
        /* main nav bar container style */
        #uj-navbar {
            display: flex;
            width: 100vw;
            height: var(--navbar-height);
            background: #fff;
            border-bottom: 2px solid var(--theme-color);
            justify-content: space-between;
            align-items: center;
            padding: .7vh 1.3vw;
            box-shadow: 0 0 10px 5px rgba(0,0,0,.05);
            z-index: 3;
            transition: all .2s ease-in-out;
        }

        :global(#ultra-janitorial-nav-logo svg) {
            width: 5vw;
        }

        :global(#uj-navbar.sticky){
            position: fixed;   
            top: 0;
        }

        /* nav bar style */
        #right-content {
            display: flex;
            width: 30%;
            font-size: 1.5rem;
            justify-content: space-evenly;
            align-items: center;
        }
        
        #right-content span {
            cursor: pointer;
            margin-right: 4em;
        }

    /*=====  End of BaseStyle  ======*/
    
    
    /*=============================================
    =            Desktop Only            =
    =============================================*/
    
    @media (any-hover: hover) {

        /* @keyframes expandBackground {
            0% {

        } */

        #right-content span {
            width: 30%;
            display: flex;
            height: var(--navbar-height);
            justify-content: center;
            align-items: center;
            margin-right: 0;
        }

        #right-content span:hover {
            
        }
    }
    
    /*=====  End of Desktop Only  ======*/
    
    
    
    /*=============================================
    =            MobileSupport            =
    =============================================*/
        
        @media screen and (max-width: 992px) {
            #uj-navbar {
                height: 6vh;
                padding: 0 1.3vw;
            }

            :global(#ultra-janitorial-nav-logo svg) {
                width: 18vw;
            }

            #right-content {
                width: 70%;
            }

            #right-content span {
                cursor: pointer;
                margin-right: 2em;
            }
        }
    
    /*=====  End of MobileSupport  ======*/
    
    



</style>

<div bind:this={self} on:scroll={addSticky} id="uj-navbar" class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div id="ultra-janitorial-nav-logo">
        <a href="#/">
            {@html ultra_janitorial_logo}
        </a>
    </div>
    <div id="right-content">
        <span class="nav-content">
            <a href="#/services">
                Services
            </a>
        </span>
        <span class="nav-content">
            <a href="#/contact">
                Contact
            </a>
        </span>
        <span class="nav-content">
            <a href="#/about">
                About
            </a>
        </span>
    </div>
</div>