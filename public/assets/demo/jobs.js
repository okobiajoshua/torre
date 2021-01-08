function loadJobs() {
    let job_listing = document.querySelector("#job-listing");
    job_listing.innerHTML = ``;
    let iHtml = ``;
    postData("/torre/search/job", {}).then(
        res => {
            res.forEach(element => {
                iHtml += renderJob(element);
            });
            console.log(iHtml);
            job_listing.innerHTML = iHtml;
        }
    )
}

function renderJob(e) {
    return `
    <div class="col-lg-6">
    <div class="card card-chart">
      <div class="card-header">
        <div class="row">
          <div class="col-sm-9 col-lg-9 text-left">
            <h5 class="card-category">${e.objective}</h5>
            <h3 class="card-title">${e.organizations[0].name}</h3>
          </div>
          <div class="col-sm-3 col-lg-3">
            <div class="photo float-right">
              <img
                src="${e.organizations[0].picture}"
                style="height: 60px; width: 60px; border-radius:30px;">
            </div>
          </div>
        </div>
      </div>
      <div class="card-body" style="padding-left: 15px;">
        <h5>${e.type}</h5>
        <h5>${e.compensation.data.currency} ${e.compensation.data.minAmount} - ${e.compensation.data.maxAmount} ${e.compensation.data.periodicity}</h5>
        <label>
          <h5>Skills Required</h5>
          - <span class="skills">${getSkills(e.skills)}</span>
        </label>

      </div>
      <div class="card-footer">
        <button type="submit" class="btn btn-fill btn-primary float-right">View</button>
      </div>
    </div>
  </div>`
}

function getSkills(skills) {
    let s = ``;
    skills.forEach(e => {
        s += `${e.name}, `
    })
    s.trim(' ,');
    return s;
}


async function postData(url = '', data = {}) {
    const response = await fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(data)
    });
    return response.json(); // parses JSON response into native JavaScript objects
}