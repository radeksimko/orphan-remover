# orphan-remover

## Motivation

Removing orphaned resources in cloud environments like AWS, GCloud, Azure etc. may be time consuming and boring in the web UI.
The goal of this project is to make the removal as easy as hit of a key.

## This can (and will) nuke real resources

Do make sure you've specified the right credentials for the right environment which you intend to clean up.
There is no undo after you delete the resources.

## Usage

For now it doesn't have any fancy CLI nor it supports anything except AWS, so the use is super simple:

```
AWS_PROFILE=myprofile go run main.go
```
