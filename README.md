# Drone - Datadog Integration
Generates a small dockerized tool that notifies Datadog whenever a build is successfully deployed.

The format for the event content is:

release-[environment]: [version]

where [environment] is generally something like "prod" or "staging" and [version] is whatever you are using for version identification (usually a semver, a hash, or a timestamp)

When running the program or the docker instance, the following environment variables need to be set:

* `dd_api_key` - a write-allowed API key created in https://app.datadoghq.com/account/settings#api
* `dd_release_environment` - the name of the environment being deployed to
* `dd_release_version` - the version of the code being deployed

Optional parameter:
* `dd_event_tags` - a list of tags to add to the event

Example usage:
```yaml
pipeline:
  datadog-event:
    image: urbint/drone-datadog
    dd_api_key: 1234567890abcdefg1234567890abcde
    dd_release_version: ${DRONE_COMMIT}
    dd_release_environment: ${DRONE_DEPLOY_TO=dev}
    dd_event_tags: [drone, deployment, release-${DRONE_DEPLOY_TO=dev}]
```

In addition, a `.tags` file placed at the root of the workspace will be parsed and all 'words' found will be appended to any tags provided in the pipeline definition.
