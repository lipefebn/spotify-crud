#!/bin/bash

if [ -z "$PATH_PUML" ]; then
    echo "PATH_UML is not specified"
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


wget -q -O plantuml.jar https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
java -jar plantuml.jar -charset UTF-8 -output "./$PATH_OUTPUT" "./$PATH_PUML/**.puml"
ls ./$PATH_PUML
echo ----------------------------
ls ./$PATH_OUTPUT