<script>
    import { txy_state, txy_server } from './TxyConf';
    import Txy, { txy_events } from './Txy';
    
    export let allow_empty = false;
    
    // style parameters
    export let font_size = '1em';
    export let font_color = '#000';
    export let text_align = 'left';
    export let extra_classes = '';

    // content parameters
    export let isHeader = false;
    export let content_key = "";
    export let fallback = "Empty content";

    // state parameters
    let show_textarea = false;

    let txy_content = Txy.getContent(content_key);
    if (txy_content === undefined) {
        txy_content = fallback;
    }

    const changeContent = (e) => {
        if (e.key === 'Enter') {
            if (show_textarea) {
                const new_content = e.target.value;
                if (new_content !== "" || allow_empty) {
                    Txy.setContent(content_key, new_content);
                }
                show_textarea = false;
            }
        }
    }

    const composeStyleFromProps = () => {
        return `
            font-size: ${font_size};
            color: ${font_color};
            text-align: ${text_align};
            cursor: ${ txy_state == 1 ? 'pointer' : 'default' };
            `;
    }

    const showTxyTextArea = () => {
        if (txy_state == 1) {
            show_textarea = true;
        }
    }

    document.addEventListener(txy_events.CONTENT_CHANGE, (e) => {
        const new_content = Txy.getContent(content_key);
        if (new_content !== undefined) {
            txy_content = new_content;
        }
    }, false);
    
</script>

<div on:click={showTxyTextArea} class="txy-text-wrapper">
    {#if show_textarea}
        <textarea on:keyup={changeContent} name="{content_key}"  cols="30" rows="10">
            {txy_content}
        </textarea>
    {:else if isHeader}
        <h2 style={composeStyleFromProps()} class="txy-text {extra_classes}">
            {txy_content}
        </h2>
    {:else}
        <p style={composeStyleFromProps()} class="txy-text {extra_classes}">
            {txy_content}
        </p>
    {/if}
</div>
