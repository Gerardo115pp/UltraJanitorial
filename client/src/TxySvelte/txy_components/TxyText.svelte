<script>
    import { txy_state } from '../TxyConf';
    import Txy, { txy_events, txy_types } from '../Txy';
    
    export let allow_empty = false;
    
    // style parameters
    export let font_size = '';
    export let font_color = '';
    export let text_align = '';
    export let extra_classes = '';

    // content parameters
    export let isHeader = false;
    export let content_key = "";
    export let fallback = "Empty content";

    // state parameters
    let txy_content = fallback;

    const getContentFromTxyStore = (e) => {
        console.log("getContentFromTxyStore:" + e.detail.content);
        if (Txy.isPageReady() && Txy.hasContent(content_key)) {
            let new_content = Txy.getContent(content_key);
            if (new_content !== "" || allow_empty) {
                txy_content = new_content;
            }            
        }
    }


    const composeStyleFromProps = () => {
        let style = "";
        if (extra_classes === "") {
            if (font_size !== '') {
                style += `font-size: ${font_size};`;
            } 
            if (font_color !== '') {
                style += `color: ${font_color};`;
            }
            if (text_align !== '') {
                style += `text-align: ${text_align};`;
            }
        }
        return style;
    }

    Txy.suscribe(txy_events.CONTENT_CHANGE, getContentFromTxyStore);
    
    if (Txy.isPageReady()) {
        txy_content = Txy.getContent(content_key);
    } else {
        Txy.suscribe(txy_events.PAGE_READY, getContentFromTxyStore, {once: true});
    }
</script>

<style>
    .txy-text {
        margin: 0;
    }

    .txy-active:hover {
        background-color: #ff634754;
    }
</style>

<div style="cursor: { txy_state ? "pointer" : "default" };" on:click={() => Txy.triggerContentGetter(content_key, txy_types.TEXT, txy_content)} class="txy-text-wrapper {txy_state ? "txy-active" : ""}">

    {#if isHeader}
        <h2 style={composeStyleFromProps()} class="txy-text {extra_classes}">
            {txy_content}
        </h2>
    {:else}
        <p style={composeStyleFromProps()} class="txy-text {extra_classes}">
            {txy_content}
        </p>
    {/if}
</div>
