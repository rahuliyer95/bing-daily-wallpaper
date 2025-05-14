#!/usr/bin/env bash

set -eu -o pipefail

exec > >(sed -u "s/^/[$(date '+%Y-%m-%d %H:%M:%S')] /")

exec 2>&1

check_internet() {
  ping -c 1 8.8.8.8 >/dev/null 2>&1
}

while ! check_internet; do
  echo "Waiting for internet..."
  sleep 5
done

OS="$(uname)"
case "$OS" in
*Darwin*)
  rm -f "$HOME/Pictures/Shortcuts Desktop Pictures/"*.jpeg
  shortcuts run "Bing Wallpaper for iOS" -i 0
  killall Dock WallpaperAgent
  ;;
*)
  # REGION="en-IN"
  # REGION="en-US"
  # BING_API="https://bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=$REGION"
  # WALLPAPER_URL="https://bing.com$(curl -sSL "$BING_API" | jq -r '.images[0].urlbase')_UHD.jpg"
  ;;
esac
