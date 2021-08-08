<script>
    import TxyText from '../../../TxySvelte/txy_components/TxyText.svelte';

    // content props
    export let step_name = "Generic name";
    export let step_description = "Generic description";
    export let step_icon = "generic";

    // style props
    export let width = "";
    export let extraClasses = "";

    const composeStyle = () => {
        let style = "";
        if (width) {
            style += `width: ${width};`;
        }
        return style;
    };
</script>

<style>
    .cleaing-step {
        display: flex;
        height: 36vh;
        background-color: #F8BC64;
        color: papayawhip;
        border-radius: 5px;
        padding: 2vh 1vw;
        flex-direction: column;
        align-items: center;
    }

    .step-icon {
        width: 20%;
        fill: papayawhip;
        margin-bottom: 1em;
    }

    .step-name {
        font-size: 2.1em;
        text-align: center;

    }

    .step-description {
        font-size: 1.2em;
        text-align: center;
        margin-top: 1em;
    }
</style>

<div style="{composeStyle()}" class="cleaing-step {extraClasses}">
    <div class="step-icon">
        {@html step_icon}
    </div>
    <div class="step-text">
        <h2 class="step-name">
            {step_name}
        </h2>
        <p class="step-description">
            <TxyText
                content_key={step_name.replace(/\s/g, '_')}
                fallback="{step_description}"
            />
        </p>
    </div>
</div>