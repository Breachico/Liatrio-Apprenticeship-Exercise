# Repository for Liatrio Apprenceship Exercise

Bryce Emery

### Description
Creating and return a simple dynamically populated JSON object using Golang.
Showing that I can do the DevOps thing! Automating updates and workflows in online development environments.

### Tasks Checklist:
- [X] Create __Github__ Repo
- [X] Create a single endpoint appliaction using __Golang__ and __Fiber__
  - [X] returns JSON object with depicted content. 
    - Time stamp is dynamically generated.
  - [X] JSON object should be "minified" (not exactly as depicted)
- [X] Write __Dockerfile__ to build (containerize?) application.
   - Unsure Dockerfile is in repository.
- [X] Create a __Github Actions Workflow__ which does the following:
  - [X] Builds application's Docker Image.
  - [X] Verifies application funcionality using Liatrio's GitHub_[apprenctice-action](https://github.com/liatrio/github-actions/tree/master/apprentice-action)
  - [X] Pushes the image to an OCI complient repository (e.g. __Docker Hub__)
  - [X] Uniquely version the image each "successful workflow run"
- [X] Deploy application to a cloud platform using image from __Docker Hub__.

- Extra Credit:
- [ ] Add a __Github Workflow__ which automatically deploys versioned application to the cloud platform when changes are made to main branch of repo.
  - Ensure versioned images are deployed from OCI repository.
- [ ] Add a field in JSON object to verify change is deployed.

- Demo:
- [X] On prior tasks completions update team on Slack, then refer back to given document to review information on Demonstrating.

