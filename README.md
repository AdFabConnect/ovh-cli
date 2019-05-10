# OVH-CLI

## Configuration

You can store configuration into `.ovh-cli.toml` in paths:
* `/etc/ovh-cli/`
* `/usr/local/etc/ovh-cli/`
* `~/` (user home)
* `./`

Configuration sample:

```toml
# File ~/.ovh-me-cli.toml

log-level = "info"
# profile = account2  

[default]
endpoint = "ovh-eu"
application-key = "XXXXXXXXXxx"
application-secret = "XXXXXXXXXXXXXXXXXXXXXXXX"
consumer-key = "XXXXXXXXXXXXXXXXXXXXXXXX"

[account2]
endpoint = "ovh-eu"
application-key = "XXXXXXXXXxx"
application-secret = "XXXXXXXXXXXXXXXXXXXXXXXX"
consumer-key = "XXXXXXXXXXXXXXXXXXXXXXXX"

...
```

Following fields MUST be in each sections:
* `endpoint`
* `application-key`
* `application-secret`
* `consumer-key`

It is possible to save which account to select in configuration file with `profile`.