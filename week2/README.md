# HTTP Server

## Overview

A simple HTTP server. It has two paths:

**/**<br />
Return "Version: _VERSION_". _VERSION_ is set as an environment variable at the HTTP server.

**/health**<br />
Return "ok". An simple health-check path.

<br>Level-based logs are enabled at level 4 at the server end:

- Level 2: Request IP, response status
- Level 4: env VERSION

## Release

```bash
make release
```

Build the httpserver into a docker image. Multi-stage build is used to ensure a minimal-sized image including only the `httpserver` binary file.

## Publish

```bash
make publish
```

This command push the docker image to dockerhub repo [dono1992](https://hub.docker.com/r/dono1992/cnhttpserver/tags).

<br>

> **_Settings_**<br>
> VERSION = tag = v1.0<br>
> image = dono1992/cnhttpserver
