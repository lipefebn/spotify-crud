#!/bin/bash

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

# This function put ./
function sanitizePath() {
    local result=$1
    
    if [[ ! result =~ '$\.\/.*' ]]; then
        result="./$result"
    fi

    echo $result
}
PATH_PUML=$(sanitizePath $PATH_PUML)

# back to the root path of the repository
function updatePath() {
    local count="${PATH_PUML//[^\/]}"
    local returns_paths
    for ((i=0; i<=${#count}-1; i++)); do
        returns_paths+="../"
    done
    echo $returns_paths
}


rm -r ./$PATH_OUTPUT
mkdir ./$PATH_OUTPUT

FILE_JAR="plantuml.jar"
wget -q -O $FILE_JAR https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
java -jar $FILE_JAR -charset UTF-8 -output $(updatePath)$PATH_OUTPUT "./$PATH_PUML/**.puml"

rm $FILE_JAR
