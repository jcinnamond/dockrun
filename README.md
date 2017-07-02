# Dockrun #

Dockrun is a tool for running a command in a Docker instance. I wrote
it to make it easier to run rspec commands from Emacs on a host
Windows system, and have the command execute inside a docker image
without the overhead of starting a new instance and booting up the
rails stack.

## How to use it ##

In a docker instance run `dockrun server`. This will listen on
port 9178. In your `docker-compose.yml` forward port 9178 to this
instance. See the `docker-compose.yml` in this project for an example.

On the host system, run `dockrun client <command>`.

## Is this a security risk? ##

Yes.

## LICENCE ##

See file in this repo.
