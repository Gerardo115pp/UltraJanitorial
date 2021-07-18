<script>

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
    }

    .content-flipped .about-card-content {
        flex-direction: row-reverse;
    }

    .about-card-description {
        display: flex;
        max-width: 60%;
        font-size: 1.8em;
        justify-content: flex-end;
        align-items: center;
    }
    
    .about-card-description p {
        width: 80%;
    }

    .about-card-image {
        width:  30vw;
    }

    .about-card-image img {
        width: 100%;
        border-radius: 5px;
    }
    
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
            <p>{description}</p>
        </div>
        <div class="about-card-image">
            <img src="{image}" alt="{name}">
        </div>
    </div>
</div>