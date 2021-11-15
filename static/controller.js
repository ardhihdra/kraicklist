const Controller = {
    search: (ev) => {
        ev.preventDefault();
        const data = Object.fromEntries(new FormData(form));
        const response = fetch(`/search?q=${data.query}`).then((response) => {
            response.json().then((results) => {
                if (!results) {
                    alert(`No result for ${data.query}`);
                    return
                }
                Controller.updateList(results);
            });
        });
    },

    openSearch: (ev) => {
        ev.preventDefault();
        const data = Object.fromEntries(new FormData(form));
        if(data.query.length) window.open(`/search?q=${data.query}&p=${0}`,"_self")
    },

    openPage: (ev) => {
        ev.preventDefault();
        const urlParams = new URLSearchParams(window.location.search);
        const query = urlParams.get('q');
        const page = ev.srcElement.innerHTML || 1;
        window.open(`/search?q=${query}&p=${page - 1}`,"_self")
    },

    updateList: (results) => {
        const rows = [];
        Controller.changeStyling()
        for (let result of results) {
            rows.push(
                `
                    <li>
                        <div>
                            <div class="result-title">
                                ${result.title}
                            </div>
                            <div class="result-highlight>
                                ${result.content}
                            </div>
                        </div>
                    </li>
                `
            );
        }
        resultList.innerHTML = rows.join(" ");
    },

    changeStyling: () => {
        searchResult.style.display = null
        divider.style.display = null
        // home.classList.remove("flex-center")
        // home.classList.remove("search-container")
        // home.classList.add("search-container-result")
        // home.classList.add("ds-row")

        // subtitle.style.display = "none"
        // title.classList.remove("search-title")
        // title.classList.add("search-title-result")
    }
};