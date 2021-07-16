<script>
    import Button from '../../../components/Button.svelte';
</script>

<style>
    @keyframes popUp {  
        0% {
            transform: translateY(200px);
        }
        100% {
            opacity: 100;
            transform: translateY(0);
        }
    }

    #hero-container {
        height: 45vh;
        background: url("images/cdc-usplash.cleaing.jpg");
        background-size: cover;
        background-position: 100% 40%;
        color: white;
        transition: all 0.5s ease-in-out;
        clip-path: polygon(100% 0%, 0% 0% , 0% 95.83%, 1% 95.83%, 2% 95.75%, 3% 95.75%, 4% 95.67%, 5% 95.50%, 6% 95.42%, 7% 95.25%, 8% 95.08%, 9% 94.92%, 10% 94.67%, 11% 94.42%, 12% 94.17%, 13% 93.92%, 14% 93.58%, 15% 93.25%, 16% 92.92%, 17% 92.58%, 18% 92.17%, 19% 91.83%, 20% 91.42%, 21% 91.00%, 22% 90.50%, 23% 90.08%, 24% 89.58%, 25% 89.17%, 26% 88.67%, 27% 88.17%, 28% 87.67%, 29% 87.17%, 30% 86.58%, 31% 86.08%, 32% 85.58%, 33% 85.00%, 34% 84.50%, 35% 83.92%, 36% 83.42%, 37% 82.83%, 38% 82.33%, 39% 81.75%, 40% 81.25%, 41% 80.67%, 42% 80.17%, 43% 79.67%, 44% 79.17%, 45% 78.67%, 46% 78.17%, 47% 77.67%, 48% 77.17%, 49% 76.67%, 50% 76.25%, 51% 75.83%, 52% 75.42%, 53% 75.00%, 54% 74.58%, 55% 74.17%, 56% 73.83%, 57% 73.50%, 58% 73.17%, 59% 72.83%, 60% 72.58%, 61% 72.33%, 62% 72.08%, 63% 71.83%, 64% 71.67%, 65% 71.42%, 66% 71.33%, 67% 71.17%, 68% 71.08%, 69% 71.00%, 70% 70.92%, 71% 70.83%, 72% 70.83%, 73% 70.83%, 74% 70.83%, 75% 70.92%, 76% 71.00%, 77% 71.08%, 78% 71.25%, 79% 71.33%, 80% 71.50%, 81% 71.75%, 82% 71.92%, 83% 72.17%, 84% 72.42%, 85% 72.67%, 86% 73.00%, 87% 73.33%, 88% 73.67%, 89% 74.00%, 90% 74.33%, 91% 74.75%, 92% 75.17%, 93% 75.58%, 94% 76.00%, 95% 76.50%, 96% 76.92%, 97% 77.42%, 98% 77.92%, 99% 78.33%, 100% 78.92%);
        margin-bottom: 20vh;

    }
    
    #hero-container:hover {
        clip-path: polygon(100% 0%, 0% 0% , 0% 72.25%, 1% 72.00%, 2% 71.83%, 3% 71.58%, 4% 71.42%, 5% 71.25%, 6% 71.17%, 7% 71.00%, 8% 70.92%, 9% 70.92%, 10% 70.83%, 11% 70.83%, 12% 70.83%, 13% 70.83%, 14% 70.92%, 15% 71.00%, 16% 71.08%, 17% 71.25%, 18% 71.42%, 19% 71.58%, 20% 71.75%, 21% 71.92%, 22% 72.17%, 23% 72.42%, 24% 72.75%, 25% 73.00%, 26% 73.33%, 27% 73.67%, 28% 74.00%, 29% 74.42%, 30% 74.75%, 31% 75.17%, 32% 75.58%, 33% 76.08%, 34% 76.50%, 35% 76.92%, 36% 77.42%, 37% 77.92%, 38% 78.42%, 39% 78.92%, 40% 79.42%, 41% 79.92%, 42% 80.50%, 43% 81.00%, 44% 81.50%, 45% 82.08%, 46% 82.58%, 47% 83.17%, 48% 83.67%, 49% 84.25%, 50% 84.75%, 51% 85.33%, 52% 85.83%, 53% 86.42%, 54% 86.92%, 55% 87.42%, 56% 87.92%, 57% 88.42%, 58% 88.92%, 59% 89.42%, 60% 89.92%, 61% 90.33%, 62% 90.75%, 63% 91.25%, 64% 91.58%, 65% 92.00%, 66% 92.42%, 67% 92.75%, 68% 93.08%, 69% 93.42%, 70% 93.75%, 71% 94.08%, 72% 94.33%, 73% 94.58%, 74% 94.83%, 75% 95.00%, 76% 95.17%, 77% 95.33%, 78% 95.50%, 79% 95.58%, 80% 95.67%, 81% 95.75%, 82% 95.83%, 83% 95.83%, 84% 95.83%, 85% 95.83%, 86% 95.75%, 87% 95.67%, 88% 95.58%, 89% 95.50%, 90% 95.33%, 91% 95.17%, 92% 95.00%, 93% 94.75%, 94% 94.58%, 95% 94.25%, 96% 94.00%, 97% 93.75%, 98% 93.42%, 99% 93.08%, 100% 92.75%);
    }

    #hero-blue-filter {
        background: rgba(0, 147, 206, 0.6);
        width: 100%;
        height: inherit;
    }
    

    #hero-content {
        padding: 5vh 5vw;
    }

    #hero-text {
        user-select: none;
    }

    #hero-text h1{ 
        color: var(--theme-color);
        font-size: 2.5vw;
    }
    
    #hero-text p{
        width: 90%;
        font-size: 1.5rem;
        margin: 2vh auto 0 auto;
    }

    #contact-btn-container {
        position: relative;
        display: flex;
        padding: 3vh 0;
        justify-content: flex-end;
    }

    :global(#contact-btn-container .popper) {
        opacity: 0;
        pointer-events: none;
    }

    :global(#hero-container:hover .popper) {
        pointer-events: all;
        animation: popUp 1s ease-out;
        animation-iteration-count: 1;
        animation-fill-mode: forwards;
        animation-delay: .02s;
    }
</style>

<div id="hero-container">
    <div id="hero-blue-filter">
        <div id="hero-content">
            <div id="hero-text">
                <h1>Welcome to Ultra Janitorial</h1>
                <p>
                    We are professionals, trained in proper way to clean, dust, and sanitize any office space. We know that having a clean office is more than just being organized, we know the way to give our customers offices a clean environment it needs. We value or customers the same way we value or team. we service each of our customers as unique 
                </p>
            </div>
            <div id="contact-btn-container">
                <Button width="10vw" padding="1vh 1vw" font_size="2rem" label={"CONTACT"} extraClasses="popper"/>
            </div>
        </div>
    </div>
</div>