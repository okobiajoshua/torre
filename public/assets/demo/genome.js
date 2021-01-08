function getBioData(username) {
    getData(`/torre/bio/${username}`).then(
        res => {
            console.log(res);
            loadNameAndDesignation(res);
        }
    )
}

function loadNameAndDesignation(el) {
    document.querySelector("#picture").setAttribute("src", el.person.picture);
    document.querySelector("#name").innerHTML = el.person.name;
    document.querySelector("#professionalHeadline").innerHTML = el.person.professionalHeadline;
}