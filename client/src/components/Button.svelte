<script>
    export let label = "button";
    export let isEnabled = true;
    export let isDimmed = false;
    export let isRounded = false;
    export let isClear = false;
    export let isDanger = false;
    export let isMaterial = false;
    export let extraClasses = "";
    
    export let font_size = "1rem";
    export let padding = "0 5%";
    export let width = "100%";
    export let button_color = "var(--theme-color)";
    
    export let onClick = () => {};
    
    const composeClassName = () => {
        let class_name = "control-btn"
        // Red class
        if (isDanger) {
            class_name += " red";
        }

        //Enable State
        class_name += isEnabled ? " enable" : " disable";

        //Enable Rounded
        class_name += isRounded ? " rounded " : "";

        // Enable clear style
        class_name += isClear ? " clear-btn " : "";

        // add a discret box shadow
        class_name += isMaterial ? " material " : "";

        //Dimmed class
        if (isDimmed) {
            class_name += " dimmed";
        }
        
        return class_name + " " + extraClasses;
    }

</script>

<style>

    @keyframes Enable {
        0% {filter: grayscale(90%);}
        25% {filter: grayscale(60%);}
        50% {filter: grayscale(40%);}
        75% {filter: grayscale(20%);}
        100% {filter: grayscale(0%);}
    }

    .control-btn {
        font-family: var(--main-font);
        cursor: pointer;
        color: var(--dark-color);
        text-align: center;
        border: none;
        border-radius: 15px;
        user-select: none;
        transition: all .4s ease-in-out;
    }

    .control-btn:hover {
        filter: brightness(1.2);
    }

    .control-btn.red {
        background-color: var(--danger) !important;
    }

    .control-btn.enable {
        animation-name: Enable;
        animation-duration: 1s;
        animation-iteration-count: 1;
    }

    .control-btn.disable {
        filter: grayscale(90%);
    }

    .control-btn.dimmed {   
        background: none !important;
        color: var(--theme-color);
        border: 1px solid var(--theme-color);
    }

    .control-btn.dimmed:hover {
        color: white;
        background-color: var(--theme-color) !important;
    }

    .control-btn.rounded {
        border-radius: 50%;
    }

    .control-btn.material {
        box-shadow: 0 0 2px 5px rgba(0, 0, 0, 0.2);
    }

    .control-btn.clear-btn {
        background-color: white !important;
        color: var(--theme-color);
        border-radius: 0;
        border-bottom: 2px solid var(--theme-color);
    }


</style>

<div style="font-size: {font_size};width: {width}; background: {button_color}; padding: {padding};" on:click={isEnabled ?  onClick : ()=>{}}
    class={composeClassName()}>
        {label}
</div>