#!/bin/bash

if [ -z "$PATH_UML" ]; then
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


wget -O plantuml.jar https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
java -jar plantuml.jar -charset UTF-8 -output $OUTPUT "./$PATH_UML/**.puml"