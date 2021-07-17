<script>
    export let image_url = 'https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png';
    export let service_name = 'ultra janitorial';
    export let service_description = 'A tool to help you keep your files clean.';

    // style variables
    export let isFliped = false;

    const classNameComposer = () => {
        let class_name = 'service-descriptor';
        if (isFliped) {
            class_name += ' flipped';
        }
        return class_name;
    };


</script>

<style>
    .service-descriptor {
        width: 100%;
        height: 45vh;
        display: flex;
        padding: 1vh 1vw;
        align-items: center; 
     }

    .service-descriptor.flipped {
        flex-direction: row-reverse;
    }
    
    .service-descriptor div:first-child {
        margin-right: 10%;
    }
    .service-descriptor-image {
        width: 80vw;
    }

    .service-descriptor-image img {
        width: 100%;
        clip-path: circle(40%);
    }

    .service-descriptor-content {
        margin-right: 10vw;
    }

    .flipped .service-descriptor-text {
        margin-right: 0;
        margin-left: 10vw;
    }

    .service-descriptor-text {
        font-size: 1.5em;
        color: var(--sencondary-color);
        text-indent: 2em;
        text-align: center;
    }
</style>

<div class="{classNameComposer()}">
    <div class="service-descriptor-image">
        <img src="{image_url}" alt="{service_name}">
    </div>
    <div class="service-descriptor-content">
        <p class="service-descriptor-text">{service_description}</p>
    </div>
</div>