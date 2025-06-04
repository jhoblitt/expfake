#!/usr/bin/env bash
set -euo pipefail

# === tweak these once, reuse forever ===
DIR=/home/jhoblitt/test2
FILESIZES=/home/jhoblitt/expfake/filesizes-2025-05-05-to-2025-06-01.yaml
HOSTMAP=/home/jhoblitt/expfake/hostmap.yaml
REMOTE_CMD="./expfake/expfake \
  --dir $DIR \
  --filesizes $FILESIZES \
  --hostmap $HOSTMAP \
  --hostname \$(hostname -s)"   # evaluated on the remote host

# === fan out to lsstcam-dc01 … lsstcam-dc10 ===
for i in {1..10}; do
  host=$(printf 'lsstcam-dc%02d' "$i")
  echo "[$host] starting …"
  ssh "$host" "$REMOTE_CMD" &
done

wait        # comment this out if you prefer sequential execution
echo "All hosts finished."
