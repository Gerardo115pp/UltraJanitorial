<script>
    import { composeImageUrl } from '../TxyConnections';
    import Txy, { txy_events, txy_types } from '../Txy';

    export let fallback;
    export let content_key;

    let image_url = fallback;
    const loadImageFromTxy = () => {
        const new_image_url = Txy.getContent(content_key);
        if (new_image_url) {
            image_url = new_image_url;
        }
    }

    if (Txy.isPageReady()) {
        image_url = Txy.getContent(content_key);
    } else {
        Txy.suscribe(txy_events.PAGE_READY, loadImageFromTxy, { once: true });
    }

    Txy.suscribe(txy_events.CONTENT_CHANGE, loadImageFromTxy);
</script>

<style>
    .txy_image_wrapper {
        cursor: pointer;
        transition: all .2s linear;
    }

    .txy_image_wrapper:hover {
        filter: blur(1px);
    }
</style>

<div on:click={() => Txy.triggerContentGetter(content_key, txy_types.IMAGE)} class="txy_image_wrapper">
    <img class="txy_image" src="{composeImageUrl(image_url)}" alt="txy_image" />
</div>