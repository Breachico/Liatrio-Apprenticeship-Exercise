# Repository for Liatrio Apprenceship Exercise

Bryce Emery

### Description
I will fill this out when I have a better idea of what I'm doing!


### Tasks Checklist:
- [ ] Create __Github__ Repo
- [ ] Create a single endpoint appliaction using __Golang__ and __Fiber__
  - [ ] returns JSON object with depicted content. 
    - Time stamp is dynamically generated.
  - [ ] JSON object should be "minified" (not exactly as depicted)
- [ ] Write __Dockerfile__ to build (containerize?) application.
   - Unsure Dockerfile is in repository.
- [ ] Create a __Github Actions Workflow__ which does the following:
  - [ ] Builds application's Docker Image.
  - [ ] Verifies application funcionality using Liatrio's GitHub_[apprenctice-action](https://github.com/liatrio/github-actions/tree/master/apprentice-action)
  - [ ] Pushes the image to an OCI complient repository (e.g. __Docker Hub__)
  - [ ] Uniquely version the image each "successful workflow run"
- [ ] Deploy application to a cloud platform using image from __Docker Hub__.

- Extra Credit:
- [ ] Add a __Github Workflow__ which automatically deploys versioned application to the cloud platform when changes are made to main branch of repo.
  - Ensure versioned images are deployed from OCI repository.
- [ ] Add a field in JSON object to verify change is deployed.

- Demo:
- [ ] On prior tasks completions update team on Slack, then refer back to given document to review information on Demonstrating.

