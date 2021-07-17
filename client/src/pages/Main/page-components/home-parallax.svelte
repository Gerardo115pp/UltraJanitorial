<script>
    import Button from '../../../components/Button.svelte';
</script>

<style>
    .home-parallax {
        display: flex;
        height: 40vh;
        background-image: url('images/ashwini-chaudhary-Iu6parQAO-U-unsplash.jpg');
        background-position: center;
        background-attachment: fixed;
        background-repeat: no-repeat;
        background-size: cover;
        margin: 5vh 0;
        justify-content: center;
        align-items: center;
        perspective: 1px;

    }


</style>

<div class="home-parallax">
    <div class="parallax-content">
        <Button font_size="1.4em" label="CONTACT" padding="2vh 2vw" />
    </div>
</div>