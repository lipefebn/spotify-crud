#!/bin/bash 

ROOT_PATH_PUML="/home/felipenunes-e14/Documents/spotify/spotify-crud/docs/diagram"
test="primeiro.png"

function getTitle() {
    local file_name_puml=${1//.png/} # remove .png
    echo $file_name_puml
    file_name_puml+=".puml"
    local root_file_name_puml="$ROOT_PATH_PUML/$file_name_puml"
    local title=$(cat $root_file_name_puml | grep -Po "(?<=^title ).+$")
    echo $title
}

getTitle $test