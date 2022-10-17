#!/bin/bash 

if [ -z "$TOKEN" ]; then
    echo "TOKEN is not specified"
    exit 1
fi
if [ -z "$WIKI_PAGE_NAME" ]; then 
    echo "WIKI_PAGE_NAME is no specified"
    exit 1
fi
if [ -z "$PATH_PUML" ] || [[ "$PATH_PUML" =~ ^[\.\/] ]]; then
    echo "PATH_PUML must follow the established pattern"
    exit 1 
fi
if [ -z "$OUTPUT" ] || [[ "$OUTPUT" =~ ^[\.\/] ]]; then
    echo "OUTPUT must follow the established pattern"
    exit 1
fi
# set the default message if no message is specified.
if [ -z "$COMMIT_MESSAGE" ]; then
    COMMIT_MESSAGE="Update wiki"
fi

# Stop the execution
function hasError() {
    if [ $? -gt 0 ]; then
        echo "ERROR:  $1"
        exit 1
    fi
}

cd .. # return to root path

# Put '/' in paths
OUTPUT="/$OUTPUT"
PATH_PUML="/$PATH_PUML"

# The default name for the wiki repository.
TEMP_REPO_NAME="wiki-repo" 

# get absolute path
ROOT_OUTPUT="$(pwd)/${TEMP_REPO_NAME}${OUTPUT}"
ROOT_PATH_PUML="${GITHUB_WORKSPACE}${PATH_PUML}"


# Clone the wiki repository and change working directory to wiki repository
function getWikiRepository() {
    git clone "https://$GITHUB_ACTOR:$TOKEN@github.com/$GITHUB_REPOSITORY.wiki.git" "$TEMP_REPO_NAME"
    hasError "Could not clone repo"
    
    # move to wiki repository
    cd "$TEMP_REPO_NAME"
    hasError "Could not move to wiki repository"
}

# Function to generate the png files
function pumlToPng() {
    # default name to plantuml
    local file_jar="plantuml.jar" 

    wget -q -O $file_jar https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
    hasError "Could not get plantuml.jar"

    java -jar $file_jar -charset UTF-8 -output $ROOT_OUTPUT "$ROOT_PATH_PUML/**.puml"
    hasError "Could not generate png files"
    
    rm $file_jar
}

# for each in png files and put in markdown
function putEachPngFile() {
    # remove old markdown file
    rm "$WIKI_PAGE_NAME"

    # get all png files 
    local files_png
    files_png=$(ls "$ROOT_OUTPUT" -t -U | grep -oP "^.+\.png$")
    hasError "Could not get png files"
    
    for file in $files_png; do
        doMarkdown $file
    done
}
# build markdown for a file
function doMarkdown() {
    local file="[[$OUTPUT/$1|alt=$1]]"
    local title=$(getTitle $1)

    echo "## $title" >> "$WIKI_PAGE_NAME"
    echo "$file" >> "$WIKI_PAGE_NAME" # image
}
# get the tittle in puml file
function getTitle() {
    local file_name_puml=${1//.png/} # remove .png
    file_name_puml+=".puml"

    local root_file_name_puml="$ROOT_PATH_PUML/$file_name_puml"
    local title=$(cat $root_file_name_puml | grep -Po "(?<=^title ).+$")
    echo $title
}

function SetConfigsGit() {
    # get configs git
    local email=`git log -1 --format="%ae"`
    
    # set configs git
    git config --global user.email "$email"
    hasError "Could not config git"

    git config --global user.name "$GITHUB_ACTOR"
    hasError "Could not config git"
}

function doPush() {
    if [ -z "$(git status --porcelain)" ]; then 
        echo "there are no modified files."
        exit 0
    elif [[ ! $(pwd) =~ \/$TEMP_REPO_NAME$ ]]; then
        echo "error: incorrect folder"
        exit 1
    fi

    git add .
    git commit -m "$COMMIT_MESSAGE" && git push "https://$GITHUB_ACTOR:$TOKEN@github.com/$GITHUB_REPOSITORY.wiki.git"
}


echo "cloning the wiki repository..."
getWikiRepository

echo "generating png files..."
pumlToPng

echo "generating the markdown file..."
putEachPngFile

echo "configuring git..."
SetConfigsGit

echo "starting the function doPush..."
doPush