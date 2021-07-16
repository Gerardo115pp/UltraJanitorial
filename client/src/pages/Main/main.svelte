<script>
    import Hero from './page-components/home-hero.svelte';

</script>

<style>

    #main-page {
        position: relative;
    }

    .center-section {
        text-transform: capitalize;
        text-align: center;
        font-size: 2.6rem;
    }

    #why-ultra-janitorial {
        height: 40vh;
    }

    .leads-container {
        display: flex;
        flex-wrap: wrap;
        margin-top: 5vh;
        justify-content: space-around;
        padding: 1rem;
    }

    .lead {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        width: 22%;
        min-height: 20vh;
        font-size: 1.4rem;
    }

    .lead-name {
        height: 10vh;
        width: 100%;
        font-size: 1.8em;
        font-weight: bold;
        text-align: center;
    }

</style>

<div id="main-page" class="page-with-nav">
    <Hero/>
    <!--Why ultra janitorial section-->
    <section id="why-ultra-janitorial">
        <h2 class="title-text center-section">Why ultra janitorial?</h2>
        <div class="leads-container">
            <div class="lead">
                <h3 class="lead-name">
                    <span class="lead-name-text">
                        Integrity
                    </span>
                </h3>
                <p class="lead-text">
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Magni nulla quibusdam dignissimos, earum soluta est eum, obcaecati impedit voluptates expedita fuga illo, atque quisquam debitis error recusandae exercitationem praesentium? Tempora atque nulla ducimus sequi facere nemo odio, corrupti eos cupiditate, natus voluptatem rerum. Aperiam, obcaecati! Aut ab fuga sequi illo.
                </p>
            </div>
            <div class="lead">
                <h3 class="lead-name">
                    <span class="lead-name-text">
                        Familiarity with Material Safety Data Sheets (MSDS)
                    </span>
                </h3>
                <p class="lead-text">
                    We know that the most important thing in the world is to know how to use it. So we're super happy with the MSDS.
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Magni nulla quibusdam dignissimos, earum soluta est eum, obcaecati impedit voluptates expedita fuga illo, atque quisquam debitis error recusandae exercitationem praesentium? Tempora atque nulla ducimus sequi facere nemo odio, corrupti eos cupiditate, natus voluptatem rerum. Aperiam, obcaecati! Aut ab fuga sequi illo.
                </p>
            </div>
            <div class="lead">
                <h3 class="lead-name">
                    <span class="lead-name-text">
                        Knowledge of cleaning chemicals and supplies
                    </span>
                </h3>
                <p class="lead-text">
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Magni nulla quibusdam dignissimos, earum soluta est eum, obcaecati impedit voluptates expedita fuga illo, atque quisquam debitis error recusandae exercitationem praesentium? Tempora atque nulla ducimus sequi facere nemo odio, corrupti eos cupiditate, natus voluptatem rerum. Aperiam, obcaecati! Aut ab fuga sequi illo.
                    We know that cleaning is a very important thing in the world. So we're super happy with cleaning supplies.
                </p>
            </div>
        </div>
    </section>
</div>