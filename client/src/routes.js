import Services from './pages/Services/Services.svelte';
import Contact from './pages/Contact/contact.svelte';
import About from './pages/About/About.svelte'
import Main from './pages/Main/main.svelte';

const routes = {
    '/': Main,
    '/contact': Contact,
    '/services': Services,
    '/about': About
};

export { routes };