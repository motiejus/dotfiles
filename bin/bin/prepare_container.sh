#!/bin/bash
set -euo pipefail

dir=$1
container=$2

TMPDIR=/bigdisk/tmp
GETIMAGE=/bigdisk/code/moby/contrib/download-frozen-image-v2.sh

prepare_container() {
    if [[ -d "$dir" ]]; then
        if [[ -f "$dir/.extract_done" ]]; then
            echo "$dir already has the filesystem extracted"
            return
        fi

        if [[ ! -f "$dir/.extract_started" ]]; then
            2>&1 echo "$dir exists and does not have magic file, bailing"
            exit 1
        else
            rm -fr "$dir"
        fi
    fi
    mkdir -p "$dir"
    touch "$dir/.extract_started"

    tmpdir=$(mktemp -d -p "$TMPDIR" "$container"XXXXX)
    trap "rm -fr $tmpdir" EXIT
    2>&1 echo "Downloading $container to $tmpdir..."
    "$GETIMAGE" "$tmpdir" "$container"
    2>&1 echo -n "Extracting $tmpdir to $dir... "
    jq -r '.[] | .Layers | .[]' "$tmpdir/manifest.json" | xargs -I{} tar -C "$dir" -xf $tmpdir/{}
    2>&1 echo "done"
    touch "$dir/.extract_done"
    2>&1 echo -n "Removing $tmpdir..."
    rm -fr "$tmpdir"
    2>&1 echo "done"
    trap - EXIT
}

prepare_container
