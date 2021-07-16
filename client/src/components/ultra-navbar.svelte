<script>
    import ultra_janitorial_logo from '../images/logo.svg';

    let self = null;

    const addSticky = () => {
        if (self) {
            const isAtBottom = window.pageYOffset > self.offsetTop
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
    
    /* main nav bar container style */
    #uj-navbar {
        display: flex;
        width: 100vw;
        max-width: 100%;
        background: #fff;
        border-bottom: 1px solid var(--theme-color);
        justify-content: space-between;
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
        align-items: center;
    }
    
    #right-content span {
        cursor: pointer;
        margin-right: 4em;
    }
</style>

<div bind:this={self} on:scroll={addSticky} id="uj-navbar" class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div id="ultra-janitorial-nav-logo">
        <a href="#uj-navbar">
            {@html ultra_janitorial_logo}
        </a>
    </div>
    <div id="right-content">
        <span class="nav-content">
            Services
        </span>
        <span class="nav-content">
            Contact
        </span>
        <span class="nav-content">
            About
        </span>
    </div>
</div>