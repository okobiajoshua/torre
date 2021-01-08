function loadPeople() {
    let job_listing = document.querySelector("#job-listing");
    job_listing.innerHTML = ``;
    let iHtml = ``;
    postData("/torre/search/people", {}).then(
        res => {
            console.log(res);
            res.forEach(element => {
                iHtml += renderProfile(element);
            });
            console.log(iHtml);
            job_listing.innerHTML = iHtml;
        }
    )
}

function searchPeople(name, skill, organization) {
    data = {
        name,
        skill,
        organization
    };

    let job_listing = document.querySelector("#job-listing");
    job_listing.innerHTML = ``;
    let iHtml = ``;

    postData("/torre/search/people", data).then(
        res => {
            console.log(res);
            res.forEach(element => {
                iHtml += renderProfile(element);
            });
            console.log(iHtml);
            job_listing.innerHTML = iHtml;
        }
    )
}

function renderProfile(e) {
    return `
    <div class="col-lg-6">
    <div class="card card-chart">
      <div class="card-header">
        <div class="row">
          <div class="col-sm-9 col-lg-9 text-left">
            <h3 class="card-title">${e.name}</h3>
            <h5 class="card-category" style="white-space: nowrap; overflow: hidden; text-overflow: ellipsis;">${e.professionalHeadline}</h5>
          </div>
          <div class="col-sm-3 col-lg-3">
            <div class="photo float-right">
              <img
                src="${e.picture}"
                style="height: 60px; width: 60px; border-radius:30px;">
            </div>
          </div>
        </div>
      </div>
      <div class="card-body" style="padding-left: 15px;">
        <h5>${e.location.name} ${e.location.name}</h5>
        <label>
          <h5>Open To</h5>
          - <span class="skills">${openTo(e.openTo)}</span>
        </label>

      </div>
      <div class="card-footer">
        <button type="submit" class="btn btn-fill btn-primary float-right">View</button>
      </div>
    </div>
  </div>`
}

function openTo(skills) {
    let s = ``;
    skills.forEach(e => {
        s += `${e}, `
    })
    s.trim(',');
    return s;
}
