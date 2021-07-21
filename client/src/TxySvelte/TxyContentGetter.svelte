<script>
    import { txy_state } from './TxyConf';
    import Txy, { txy_events, txy_types } from './Txy';
import { bind } from 'svelte/internal';


    export let font_family = 'Monospace';
    export let theme_color = 'yellowgreen';

    let is_showing = false;
    let content_type = "";
    let content_name = "";
    let content_text = "";
    let file_input = null;


       
    const closeModal = e => {
        if(e.target === e.currentTarget) {
            reset();
        }
    }

    const handleGetContentEvent = (e) => {
        is_showing = true;
        content_type = e.detail.content_type;
        content_name = e.detail.content_name;
        content_text = e.detail.current_content;
    }

    const handleTextContent = (e) => {
        if(e.key == "Enter"){
            Txy.setContent(content_name, content_text);
            reset();
        }
    }

    const handleImageContent = (e) => {
        // get the image from the file input and encoded as blob
        const file = file_input.files[0];
        const reader = new FileReader();
        reader.onload = (e) => {
            const data = e.target.result;
            alert("loaded");
            Txy.setContentFromFile(file, content_name);
            reset();
        }
        reader.readAsDataURL(file);
    }

    const reset = () => {
        is_showing = false;
        content_type = "";
        content_name = "";
        content_text = "";
    }

    document.addEventListener( txy_events.GET_CONTENT, handleGetContentEvent, false );

</script>

<style>
    /* modal background */
    #txy-content-getter-wrapper {
        cursor: pointer;
        position: fixed;
        display: flex;
        top: 0;
        left: 0;
        width: 100%;
        height: 100vh;
        justify-content: center;
        align-items: center;
        background-color: rgba(255, 214, 175, 0.397);
        z-index: 3;
    }

    #txy-content-getter {
        cursor: default;
        width: 50%;
        height: 40vh;
        background-color: white;
        border-radius: 15px;
        box-shadow: 0 15px 15px 5px rgba(0, 0, 0, 0.05);
    }

    #txy-content-getter-title {
        font-size: 1.5em;
        height: 3em;
        text-align: center;
        padding: .75em;
        font-weight: bold;
    }

    #txy-content {
        display: flex;
        height: calc(40vh - 3em);
        justify-content: center;
        align-items: center;
    }

    #txy-textarea {
        display: block;
        width: 90%;
        height: 20vh;
        padding: 1vh .5vw;
        margin: 5vh auto;
        color: gray;
        font-size: 1.3em;
        border-radius: 15px;
        border: 1px solid yellowgreen;
        resize: none;
        outline: none;
        scrollbar-width: none;
        -ms-overflow-style: none;
    }

    #txy-textarea::-webkit-scrollbar {
        display: none;
    }
    
    #txy-file-input {
        display: none;
    }

    #txy-file-btn {
        cursor: pointer;
        width: 20%;
        text-align: center;
        font-size: 1.5em;
        padding: 1vh 1vw;
        border-radius: 15px;
        color: white;
    }
</style>

{#if is_showing}
    <div on:click={closeModal} style="font-family: {font_family};" id="txy-content-getter-wrapper">

            <div id="txy-content-getter">
                <div style="color: {theme_color};" id="txy-content-getter-title">
                    New content
                </div>
                <div id="txy-content">
                    {#if content_type == "text"}
                        <textarea bind:value={content_text} on:keyup={handleTextContent} style="border-color: {theme_color};" name="txy-text" id="txy-textarea" placeholder="new content"></textarea>
                    {:else if content_type == "image"}
                        <input bind:this={file_input} on:change={handleImageContent} type="file" name="txy-file" id="txy-file-input" accept="image/*" />
                        <div on:click={() => file_input.click()} style="background-color: {theme_color};" id="txy-file-btn">
                            Upload Image
                        </div>
                    {/if}
                </div>
            </div>

    </div>
{/if}