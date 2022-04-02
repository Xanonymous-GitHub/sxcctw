#!/usr/bin/env bash

start_go_watch() {
  local entry_name="$1"
  readonly current_script_pos="$(dirname "$0")"
  local entry_top_path="$current_script_pos/cmd"
  local full_entry_path="$entry_top_path/$entry_name"

  if [ ! -d "$full_entry_path" ]; then
    echo "$full_entry_path does not exists."
    exit 1
  elif [ ! -f "$full_entry_path/gowatch.yml" ]; then
    echo "no gowatch.yml exists in $full_entry_path."
    exit 1
  fi

  go mod tidy

  if ! pgrep -x "$entry_name" >/dev/null; then
    pkill "$entry_name" -9
  fi

  cd "$full_entry_path" && gowatch
  return 0
}

if [ -z "$1" ]; then
  echo "No go watch entry point name supplied."
  exit 1
fi

start_go_watch "$1"
exit 0
