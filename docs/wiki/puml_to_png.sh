#!/bin/bash

PATH_PUML="docs/diagram/puml"
#docs/diagram/images


if [ -z "$PATH_PUML" ]; then
    echo "PATH_PUML is not specified"
    exit 1
fi

if [ -z "$PATH_OUTPUT" ]; then
    echo "OUTPUT is not specified"
    exit 1
fi

if [ -z "$TOKEN" ]; then
    echo "TOKEN is not specified"
    exit 1
fi

function sanitizePath() {
    local result=$1
    
    if [[ ! result =~ '$\.\/.*' ]]; then
        result="./$result"
    fi

    echo $result
}
PATH_PUML=$(sanitizePath $PATH_PUML)

function updatePath() {
    local count="${PATH_PUML//[^\/]}"
    local returns_paths
    for ((i=0; i<=${#count}-1; i++)); do
        returns_paths+="../"
    done
    echo $returns_paths
}

PATH_OUTPUT=$(updatePath)$PATH_OUTPUT

wget -q -O plantuml.jar https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
rm -r ./$PATH_OUTPUT
mkdir ./$PATH_OUTPUT
java -jar plantuml.jar -charset UTF-8 -output "../images" "./$PATH_PUML/**.puml"
ls ./$PATH_PUML
echo ----------------------------
ls ./$PATH_OUTPUT