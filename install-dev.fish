#!/usr/bin/env fish

if not test -f go.mod
    echo "Run this from the Feedwatch project root." >&2
    exit 1
end

set -l bin_dir ~/.local/bin
mkdir -p "$bin_dir"; or exit 1
go build -o "$bin_dir/feedwatch" .; or exit 1

fish_add_path --universal --move "$bin_dir"
echo "Installed $bin_dir/feedwatch"
