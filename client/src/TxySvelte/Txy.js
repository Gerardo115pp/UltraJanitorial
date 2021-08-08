import { getPageContent, postFileContent, postTextContent, loadPages } from './TxyConnections';
import { txy_state, mobild_edit } from './TxyConf';

class TxyContent {
    constructor() {
        this.content = {};
        this.page_state = {
            name: "",
            ready: false
        };
        this.content.name = "Txy";
        this.loadGeneralContent();
    }

    createPage = (name) => {
        
        /**
         *
         * Creates a new page locally and also sends a request to the server to create the page but it does not
         * wait for the server to respond because a new page has no content, so it is not necessary to wait for the
         * server to respond. 
         *
         */
        
        this.content[name] = {};
        getPageContent(name); // the server will create the page even if it does not exist
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

    getPageName = () => this.page_state.name;

    getContent = (content_name, page_name="") => {
        page_name = page_name || this.getPageName();
        console.log("getContent: " + content_name + " " + page_name);
        if (this.content.hasOwnProperty(page_name)) {
            if (this.content[page_name].hasOwnProperty(content_name)) {
             
                if (this.page_state.name !== page_name) {
                    console.warn(`Content from page '${page_name}' was requested but current page is '${page_name}'`);
                }
             
                return this.content[page_name][content_name];
            }
        } else {
            throw new Error(`Page doesnt exists: ${page_name}`);
        }
    };
        
    hasContent = (content_name, page_name="") => {
        page_name = page_name || this.getPageName();
        return this.content[page_name].hasOwnProperty(content_name);
    };
    
    isPageReady = () => this.page_state.ready;

    isGeneralContentReady = () => Object.keys(this.content).length > 0;

    isTxyActive = () => txy_state === 1;

    isContentEditable = () => {
        if (!this.isTxyActive()) {
            return false;
        }
        if (is_mobile && !mobild_edit) {
            return false;
        }
        return true;
    }
    
    joinContent = (other, page_name) => {
        if(!this.content.hasOwnProperty(page_name)) { 
            this.content[page_name] = {};
        }
        for (let key in other) {
            this.content[page_name][key] = other[key];
        }
    }

    loadGeneralContent = () => {
        loadPages().then(pages => {
            Object.keys(pages).forEach(page_name => {
                this.content[page_name] = pages[page_name];
            });
            let current_page = this.getPageName();
            console.log(`current page: ${current_page}`)
            this.dispatchTxyEvent(txy_events.GENERAL_CONTENT_LOADED, {
                pages: Object.keys(this.content),
                current_page: current_page
            });

            // in case setPageName was called before the general content was loaded
            if (current_page !== "") { 
                if (!pages.hasOwnProperty(current_page)) {
                    this.createPage(current_page);
                }
                this.page_state.ready = true;
                this.dispatchTxyEvent(txy_events.PAGE_READY, this.page_state);
            } 
        });
    }

    suscribe = (event_name, callback, options=false) => {
        if (this.eventExists(event_name)) {
            document.addEventListener(event_name, callback, options);
        } else {
            throw new Error(`Event ${event_name} does not exist`);
        }
    }

    setPageName = (page_name) => {
        if (!this.page_state.ready) {
            this.page_state.name = page_name;
            if (!this.content.hasOwnProperty(page_name)) {
                this.page_state.ready = false;
                
                // if the general content is not ready yet await for it, in case the page is not loaded in the general content, the loader will
                // realize the current page is not loaded then it doesnt exists so it will tell the server to create it.
                if (this.isGeneralContentReady()) {               
                    getPageContent(page_name).then(page_content => {
                        this.joinContent(page_content, page_name);
                        this.page_state.ready = true;
                        this.dispatchTxyEvent(txy_events.PAGE_READY, this.page_state);
                    }); 
                }
            } else {
                this.page_state.ready = true;
                this.dispatchTxyEvent(txy_events.PAGE_READY, this.page_state);
            }
        }
        console.log(`${page_name} status: ${this.page_state.ready}`);
    }
    
    setContent = (content_name, content, page_name="") => {
        page_name = page_name || this.getPageName();
        if (this.isContentEditable()) {
            this.content[page_name][content_name] = content;
            postTextContent(page_name, content, content_name).then(() => {
                this.dispatchTxyEvent(txy_events.CONTENT_CHANGE, {
                    content_name: content_name,
                    content: content
                });
            });
        }
    }

    setContentFromFile = (file_blob, content_name, page_name="") => {
        page_name = page_name || this.getPageName(); 
        if (this.isContentEditable()) {
            postFileContent(page_name, file_blob, content_name).then(response => {
                this.content[page_name][content_name] = response;
                this.dispatchTxyEvent(txy_events.CONTENT_CHANGE, {
                    content_name: content_name,
                    content: response
                });
            });
        }
    }

    triggerContentGetter = (content_name, content_type, current_content="a", page_name="") => {
        page_name = page_name || this.getPageName();
        if (this.isContentEditable()) {
            this.dispatchTxyEvent(txy_events.GET_CONTENT, {
                content_name: content_name,
                content_type,
                current_content,
                page_name
            });
        }
    }

    unsetPage = () => {
        this.page_state.name = "";
        this.page_state.ready = false;
    }
}


export const txy_events = {
    CONTENT_CHANGE: "content-change",
    PAGE_READY: "page-name-change",
    GET_CONTENT: "get-content",
    GENERAL_CONTENT_LOADED: "general-content-loaded"
}

export const txy_types = {
    TEXT: "text",
    IMAGE: "image"
}

export const is_mobile = window.screen.width <= 768;

const txy_content = new TxyContent();
Object.freeze(txy_content);
export default txy_content;