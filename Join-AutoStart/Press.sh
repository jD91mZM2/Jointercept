#!/bin/bash
if [[ "$1" == "press=:="* ]]; then
    key=$(echo "$1" | cut -c 9-)
    xdotool key $key
fi
