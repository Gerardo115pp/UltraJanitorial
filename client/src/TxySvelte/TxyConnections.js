import { txy_server } from "./TxyConf";

export const getPageContent = (page) => {
    return fetch(`http://${txy_server}/content?page=${page}`, {method: 'GET'})
            .then(response => response.json())
}