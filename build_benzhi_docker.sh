#!/usr/bin/env bash
set -euo pipefail
docker build -f benzhi.Dockerfile -t "${1:-app}:benzhi" .
