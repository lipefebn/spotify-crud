#!/bin/bash 

PATH_DIAGRAMS=".ada/ada"
#if [ -z "$PATH_DIAGRAMS" || ["$PATH_DIAGRAMS" =~ ^[\.\/]] ]; then
if [ -z "$PATH_DIAGRAMS" ] || [[ "$PATH_DIAGRAMS" =~ ^[\.\/] ]]; then
    echo sla
fi