function getBioData(username = 'okobiajoshua') {
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
    document.querySelector("#summaryOfBio").innerHTML = el.person.summaryOfBio;

    if (el.person.links.length > 0) {
        let links = ``;
        el.person.links.forEach(lnk => {
            links += `<button href="${lnk.address}" class="btn btn-icon btn-round btn-${lnk.name}">
            <i class="fab fa-${lnk.name}"></i>
          </button>`
        });
        document.querySelector("#social-media-links").innerHTML = links;
    }

    renderProfile(el);
}

function renderProfile(el) {
    let iHtml = ``;
    iHtml += renderExperiences(el.jobs, "Jobs");
    iHtml += renderExperiences(el.awards, "Awards");
    iHtml += renderExperiences(el.education, "Education");
    document.querySelector("#experiences").innerHTML = iHtml;
}

function renderExperiences(exps, title) {
    if (exps.length < 1) return ``;
    let workExperiences = ``;

    exps.forEach(exp => {
        workExperiences += renderSingleExperience(exp);
    });
    let r = `
    <div class="card">
        <div class="card-header">
        <h5 class="title">${title}</h5>
        </div>
        <div class="card-body">
        <div class="row">
            ${workExperiences}
        </div>
        </div>
    </div>`;

    return r;
}

function renderSingleExperience(exp) {
    let responsibilities = renderResponsibilities(exp.responsibilities);
    let org = exp.organizations.length > 0 ? `<h5 class="card-category"
    style="white-space: nowrap; overflow: hidden; text-overflow: ellipsis;">
    ${exp.organizations[0].name}
    </h5>` : '';
    let resp = responsibilities.trim().length > 0 ? `<div class="card-body" style="padding-left: 15px;">
    <h5>Responsibilities</h5>
      ${responsibilities}
    </div>` : '';

    return `<div class="card card-chart">
        <div class="card-header">
          <div class="row">
            <div class="col-sm-9 col-lg-9 text-left">
              <h3 class="card-title">${exp.name}</h3>
                ${org}
            </div>
            <div class="col-sm-3 col-lg-3">
            </div>
          </div>
          ${resp}
        </div>
        
    </div>`;
}

function renderResponsibilities(responsibilities) {
    if (responsibilities.length < 1) return ``;
    let res = `
    <ul>`
    responsibilities.forEach(r => {
        res += `<li>${r}</li>`;
    });
    res += `</ul>`;
    console.log(res);
    return res;
}
