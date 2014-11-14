This will allow you to run `docker` commands under the CloudFoundry `cf` command.  Why?  Why not?  :-)
It provides the appearance of giving a single UX for clients but still giving them access to `cf` and `docker` commands.
When combined with [cf-docker](https://github.com/duglin/cf-docker) some interesting usage patterns might arise.

Installing
----------
Download the appropiate binary from the [`binaries`](binaries) dir.

Run:
`cf install-plugin /full/path/to/your/cf-docker-plugin`

Running
-------
Just use `cf docker` wherever you would use `docker`.
Make sure you set your `DOCKER_HOST` environment variable first.
