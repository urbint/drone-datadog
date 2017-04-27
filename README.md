# Drone - Datadog Integration
Generates a small dockerized tool that notifies Datadog whenever a build is successfully deployed.

The format for the event content is:

release-[environment]: [version]

where [environment] is generally something like "prod" or "staging" and [version] is whatever you are using for version identification (usually a semver, a hash, or a timestamp)
