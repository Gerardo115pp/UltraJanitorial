import { getPageContent, postFileContent, postTextContent } from './TxyConnections';
import { txy_state } from './TxyConf';

class TxyContent {
    constructor() {
        this.content = {};
        this.page_state = {
            name: "",
            ready: false
        };
        this.content.name = "Txy";
    }
    
    clearContent = () => {
        for (var key in this.content) {
            if (this.content.hasOwnProperty(key)) {
                delete this.content[key];
            }
        }
    }

    dispatchTxyEvent = (event_name, details) => {
        if (this.eventExists(event_name)) {
            if (this.isTxyActive()) {
                document.dispatchEvent(new CustomEvent(event_name, {
                    detail: details
                }));
            }
        } else {
            throw new Error("Txy event does not exist: " + event_name);
        }
    }

    eventExists = (event_name) => {
        for (let event of Object.values(txy_events)) {
            if (event === event_name) {
                return true;
            }
        }
        return false;
    }


    getPageName = () => {
        return this.page_state.name;
    }

    getContent = (content_name) => {
        return this.content[content_name];
    }

    hasContent = (content_name) => {
        return this.content.hasOwnProperty(content_name);
    }

    isPageReady = () => {
        return this.page_state.ready;
    }

    isTxyActive = () => {
        return txy_state === 1;
    }

    joinContent = (other) => {
        for (let key in other) {
            if (other.hasOwnProperty(key)) {
                this.content[key] = other[key];
            }
        }
    }

    suscribe = (event_name, callback, options=false) => {
        if (this.eventExists(event_name)) {
            document.addEventListener(event_name, callback, options);
        } else {
            throw new Error(`Event ${event_name} does not exist`);
        }
    }

    setPageName = (page_name) => {
        if (this.isTxyActive()) {
            this.page_state.name = page_name;
            this.page_state.ready = false;
            getPageContent(page_name).then(page_content => {
                this.clearContent();
                this.joinContent(page_content);
                this.page_state.ready = true;
                this.dispatchTxyEvent(txy_events.PAGE_READY, this.page_state);
            }); 
        }
    }

    setContent = (content_name, content) => {
        if (this.isTxyActive()) {
            this.content[content_name] = content;
            postTextContent(this.page_state.name, content, content_name).then(() => {
                this.dispatchTxyEvent(txy_events.CONTENT_CHANGE, {
                    content_name: content_name,
                    content: content
                });
            });
        }
    }

    setContentFromFile = (file_blob, content_name) => { 
        if (this.isTxyActive()) {
            postFileContent(this.page_state.name, file_blob, content_name).then(response => {
                this.dispatchTxyEvent(txy_events.CONTENT_CHANGE, {
                    content_name: content_name,
                    content: response
                });
            });
        }
    }

    triggerContentGetter = (content_name, content_type, current_content="a") => {
        this.dispatchTxyEvent(txy_events.GET_CONTENT, {
            content_name: content_name,
            content_type,
            current_content
        });
    }
}

export const txy_events = {
    CONTENT_CHANGE: "content-change",
    PAGE_READY: "page-name-change",
    GET_CONTENT: "get-content"
}

export const txy_types = {
    TEXT: "text",
    IMAGE: "image"
}

const txy_content = new TxyContent();
Object.freeze(txy_content);
export default txy_content;