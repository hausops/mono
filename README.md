# mono

Monorepo for HausOps

## Local Dev Environment

We use [Dapr](https://dapr.io/) to manage service-to-service communication. To run multiple HausOps services locally and allow them to communicate with each other, you have to run them through Dapr.

### Gettings started

1. Install [dapr-cli](https://docs.dapr.io/getting-started/install-dapr-cli/).
1. Install [Podman](https://podman.io/getting-started/installation). For example, on MacOS, you can install it with `brew install podman`.
1. Start Podman:

   ```sh
   podman machine init
   podman machine start
   ```

1. Run all services locally as configured in `hausops/mono/dapr.yaml` using [Dapr Multi-App Run](https://docs.dapr.io/developing-applications/local-development/multi-app-dapr-run/multi-app-overview/):

   ```sh
   # cd hausops/mono
   dapr run -f .
   ```

1. To work on a service (temporary measure; let's figure out something better):

   - Comment it out from `dapr.yaml`.
   - Running the service using `dapr run`:

   ```sh
   # Example

   # cd apps/dashboard-api
   dapr run --app-id dashboard-api -- make run
   ```

   This way, you don't have to start/stop all services running via Multi-App Run when you need to start/stop the service under development.

### Clean up

To clean up your local environment:

```sh
# This command is the reverse of `dapr init --container-runtime podman`.
# It will remove Dapr containers from Podman as well.
dapr uninstall --container-runtime podman --all

podman machine stop
podman machine rm
```
