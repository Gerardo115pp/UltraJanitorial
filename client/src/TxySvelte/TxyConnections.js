import { txy_server } from "./TxyConf";

export const composeImageUrl = ( partial_url ) => `http://${ txy_server }${ partial_url }`;

export const getPageContent = (page) => {
    return fetch(`http://${txy_server}/content?page=${page}`, {method: 'GET'})
            .then(response => response.json())
}

export const loadPages = () => {
    return fetch(`http://${txy_server}/pages`, {method: 'GET'})
            .then(response => response.json())
}

export const postFileContent = (page, file, name) => {
    const form = new FormData();
    form.append("file", file, file.name);
    form.append("page", page);
    form.append("name", name);

    const request = new Request(`http://${txy_server}/static/${page}`, { method: 'POST', body: form });
    return fetch(request).then(response => response.text());
}

export const postTextContent = (page, text, name) => {
    const form = new FormData();
    form.append('content', text);
    form.append('name', name);
    form.append('page', page);
    
    const request = new Request(`http://${txy_server}/content`, { method: 'POST', body: form });
    return fetch(request);
}