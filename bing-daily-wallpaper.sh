#!/bin/sh

OS="$(uname)"
REGION_IN="en-IN"
REGION_US="en-US"
BING_API="https://bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=$REGION_US"

check_internet() {
  ping -c 1 8.8.8.8 > /dev/null 2>&1
}

check_internet
while [ $? -ne 0 ]
do
  echo "Waiting for internet..."
  sleep 5
  check_internet
done

WALLPAPER_URL="https://bing.com$(curl -sSL "$BING_API" | jq -r '.images[0].url')"

case "$OS" in
  *Darwin*)
    echo "Downloading wallpaper"
    curl -sSLo "$HOME/Pictures/wallpaper.jpg" "$WALLPAPER_URL"
    echo "Setting wallpaper"
    osascript -e "tell application \"Finder\" to set desktop picture to \"$HOME/Pictures/wallpaper.jpg\" as POSIX file"
    killall Dock
  ;;
esac
