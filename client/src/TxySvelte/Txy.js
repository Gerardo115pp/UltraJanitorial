import { getPageContent } from './TxyConnections'

class TxyContent {
    constructor() {
        this.content = {};
        this.page_name = {page: ""};
        this.content.name = "Txy";
    }
    
    // Clears the content object without reassigning it.
    clearContent = () => {
        for (var key in this.content) {
            if (this.content.hasOwnProperty(key)) {
                delete this.content[key];
            }
        }
    }

    // dispatch a 'content change' event
    dispatchTxyEvent = (event_name, details) => {
        document.dispatchEvent(new CustomEvent(event_name, {
            detail: details
        }));
    }


    // get current page name
    getPageName = () => {
        return this.page_name.page;
    }

    // get content by name
    getContent = (content_name) => {
        return this.content[content_name];
    }

    // method that recives an object and adds all the properties to the content object
    joinContent = (other) => {
        for (let key in other) {
            if (other.hasOwnProperty(key)) {
                this.content[key] = other[key];
            }
        }
    }

    // set page name
    setPageName = (page_name) => {
        this.page_name.page = page_name;
        getPageContent(page_name).then(page_content => {
            this.clearContent();
            this.joinContent(page_content);
            this.dispatchTxyEvent(txy_events.PAGE_NAME_CHANGE, this.page_name);
        }); 
    }

    // add new content to the content object or reset the content name
    setContent = (content_name, content) => {
        this.content[content_name] = content;
        this.dispatchTxyEvent(txy_events.CONTENT_CHANGE, {
            content_name: content_name,
            content: content
        });
    }
}

export const txy_events = {
    CONTENT_CHANGE: "content-change",
    PAGE_NAME_CHANGE: "page-name-change"
}

const txy_content = new TxyContent();
Object.freeze(txy_content);
export default txy_content;