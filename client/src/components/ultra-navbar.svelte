<script>
    import ultra_janitorial_logo from '../images/logo.svg';
</script>

<style>
    
    /* main nav bar container style */
    #uj-navbar {
        display: flex;
        background: #fff;
        border-bottom: 1px solid var(--theme-color);
        justify-content: space-between;
        padding: .7vh 1.3vw;
        box-shadow: 0 0 10px 5px rgba(0,0,0,.05);
    }

    :global(#ultra-janitorial-nav-logo svg) {
        width: 5vw;
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

<div id="uj-navbar" class="navbar navbar-default navbar-fixed-top" role="navigation">
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
            Contact Us
        </span>
        <span class="nav-content">
            About Us
        </span>
    </div>
</div>