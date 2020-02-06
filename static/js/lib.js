function setup() {
    if (document.querySelector('.list-of-contents') !== null) {
        createContentsList();
    }
}

/*
 * Creates contents list for a post if contents section is available.
 */
function createContentsList() {
    const sections = enumerateSections();
    const contents = document.querySelector('.list-of-contents ul');

    for (const sectionDef of sections) {
        const anchorLink = document.createElement('a');
        anchorLink.href = `#${sectionDef['id']}`;
        anchorLink.appendChild(document.createTextNode(sectionDef['header']));
        const listItem = document.createElement('li');
        listItem.classList += " post-section-link list-unstyled";
        listItem.appendChild(anchorLink);
        contents.appendChild(listItem);
    }
}

/*
 * Appends numerical indices to each section's header. Returns total number of
 * discovered sections.
 */
function enumerateSections(selector = "h4.header") {
    let content = [];
    const postSectionHeaders = Array.from(document.querySelectorAll(selector));

    for (const [index, header] of postSectionHeaders.entries()) {
        const headerText = `[${index}] ${header.textContent}`;
        const contentEntry = {
            id: header.id,
            header: headerText,
        };
        content.push(contentEntry);
        header.innerHTML = '';
        header.appendChild(document.createTextNode(headerText));
    }

    return content;
}
