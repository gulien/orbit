#!/usr/bin/env bash

go fmt $(go list ./... | grep -v /vendor/);
