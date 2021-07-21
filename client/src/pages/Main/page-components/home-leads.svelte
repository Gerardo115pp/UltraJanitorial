<script>
    import TxyText from '../../../TxySvelte/txy_components/TxyText.svelte';
import { lorem } from '../../../Utilitys';

    const txy_keys = {
        why_ultra: 'leads-why-ultra',
        lead_I_title: 'lead-I-title',
        lead_I_text: 'lead-I-text',
        lead_II_title: 'lead-II-title',
        lead_II_text: 'lead-II-text',
        lead_III_title: 'lead-III-title',
        lead_III_text: 'lead-III-text'
    };

</script>

<style>
    #why-ultra-janitorial {
        max-height: 60vh;
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

    .lead-text {
        color: var(--secondary-color);
    }
</style>

<section id="why-ultra-janitorial">
    <h2 class="title-text center-section">
        <TxyText fallback={"Why Ultra Janitorial"} content_key={txy_keys.why_ultra}/>
    </h2>
    <div class="leads-container">
        <div class="lead">
            <h3 class="lead-name">
                <span class="lead-name-text">
                    <TxyText fallback={"Integrity"} content_key={txy_keys.lead_I_title}/>
                </span>
            </h3>
            <p class="lead-text">
                <TxyText fallback={lorem(50)} content_key={txy_keys.lead_I_text}/>
            </p>
        </div>
        <div class="lead">
            <h3 class="lead-name">
                <span class="lead-name-text">
                    <TxyText 
                        fallback={"Familiarity with Material Safety Data Sheets (MSDS)"} 
                        content_key={txy_keys.lead_II_title}
                    />
                    
                </span>
            </h3>
            <p class="lead-text">
                <TxyText 
                    fallback={"We know that the most important thing in the world is to know how to use it. So we're super happy with the MSDS. Lorem, ipsum dolor sit amet consectetur adipisicing elit. Magni nulla quibusdam dignissimos, earum soluta est eum, obcaecati impedit voluptates expedita fuga illo, atque quisquam debitis error recusandae exercitationem praesentium? Tempora atque nulla ducimus sequi facere nemo odio, corrupti eos cupiditate, natus voluptatem rerum. Aperiam, obcaecati! Aut ab fuga sequi illo."} 
                    content_key={txy_keys.lead_II_text}
                />
            </p>
        </div>
        <div class="lead">
            <h3 class="lead-name">
                <span class="lead-name-text">
                    <TxyText
                        fallback={"Knowledge of cleaning chemicals and supplies"}
                        content_key={txy_keys.lead_III_title}
                    />
                </span>
            </h3> 
            <p class="lead-text">
                <TxyText
                    fallback={lorem(50)}
                    content_key={txy_keys.lead_III_text}
                />                
            </p>
        </div>
    </div>
</section>