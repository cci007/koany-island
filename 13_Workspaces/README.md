# Workspaces

**Description:**

In this challenge, 3 applications will be created in a workflow. One for Linux, Windows, and MacOS. Each build job will generate an app file that needs to be persisted to the final `deploy` job, which will then upload these apps as artifacts.

**Goals:**

- Persist each of the apps in the build jobs to the final deploy job.
- Upload all of the apps to your artifacts.
- Provide a link to your final workflow.

**Help:**
<details>
  <summary>Spoiler warning</summary>
  
  *         - store_artifacts: # See circleci.com/docs/2.0/artifacts/ for more details.
            path: app
            destination: apps
  * https://circleci.com/blog/build-cicd-piplines-using-docker/
  
</details>
