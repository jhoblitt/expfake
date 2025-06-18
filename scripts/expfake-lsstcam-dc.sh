#!/usr/bin/env bash

set -euo pipefail

EXPFAKE_VER=v1.1.2
EXPFAKE_DIR="${PWD}/expfake"
DATASET_DIR="${EXPFAKE_DIR}/test2"
FILESIZES="${EXPFAKE_DIR}/filesizes-2025-05-05-to-2025-06-01.yaml"
HOSTMAP="${EXPFAKE_DIR}/hostmap.yaml"
REMOTE_CMD="${EXPFAKE_DIR}/expfake \
  --dir $DATASET_DIR \
  --filesizes $FILESIZES \
  --hostmap $HOSTMAP \
  --hostname \$(hostname -s)"  # evaluated on the remote host

mkdir -p "$EXPFAKE_DIR"

(
  cd "$EXPFAKE_DIR"
  curl -L "https://github.com/jhoblitt/expfake/releases/download/${EXPFAKE_VER}/expfake-linux-amd64.tar.gz" | tar -xz
  curl -O "https://raw.githubusercontent.com/jhoblitt/expfake/refs/tags/${EXPFAKE_VER}/hostmap.yaml"
  curl -O https://raw.githubusercontent.com/jhoblitt/expsize/refs/tags/v1.0.0/filesizes-2025-05-05-to-2025-06-01.yaml
)

clush -g dc --copy expfake

# === fan out to lsstcam-dc01 - lsstcam-dc10 ===
for i in {1..10}; do
  host=$(printf 'lsstcam-dc%02d' "$i")
  echo "[$host] starting …"
  # shellcheck disable=SC2029
  ssh "$host" "$REMOTE_CMD" &
done

wait
echo 'All hosts finished.'
