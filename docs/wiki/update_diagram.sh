#!/bin/bash 

# if [ -z "$TOKEN" ]; then
#     echo "TOKEN is not specified"
#     exit 1
# fi
if [ -z "$PATH_DIAGRAMS" ] || [[ "$PATH_DIAGRAMS" =~ ^[\.\/] ]]; then
    echo "PATH_DIAGRAMS is not specified" #TODO
    exit 1
fi
if [ -z "$OUTPUT_DIAGRAMS" ] || [[ "$OUTPUT_DIAGRAMS" =~ ^[\.\/] ]]; then
    echo "OUTPUT_DIAGRAMS is no specified"
    exit 1
fi

OUTPUT_DIAGRAMS="/$OUTPUT_DIAGRAMS"
PATH_DIAGRAMS="/$PATH_DIAGRAMS"

# [[/diagrams/diagrama_de_test.png|alt=diagrama_de_test]]
PATH_DIAGRAMS_GIT="https://github.com/${GITHUB_REPOSITORY}/wiki/diagrams"
MESSAGE_COMMIT=`git log -1 --format="%s"`

# Stop the execution
function hasError() {
    if [ $? -gt 0 ]; then
        echo "ERROR:  $1"
        exit 1
    fi
}

function SetConfigsGit() {
    # get configs git
    local author=`git log -1 --format="%an"`
    local email=`git log -1 --format="%ae"`

    # set configs git
    git config --global user.email "$email"
    git config --global user.name "$author"
}

# The default name for the wiki repository.
TEMP_REPO_NAME="wiki-repo" 
# Clone the wiki repository and change working directory to wiki repository
function getWikiRepository() {
    cd .. # return to root path
    hasError "Could not return to root path"

    git clone "https://$GITHUB_ACTOR:$TOKEN@github.com/$GITHUB_REPOSITORY.wiki.git" "$TEMP_REPO_NAME"
    hasError "Could not clone repo"
    
    # move to wiki repository
    cd "$TEMP_REPO_NAME"
    hasError "Could not move to wiki repository"
    
    # remove old file
    rm Diagrams.md
}

function pumlToPng() {
    PATH_OUTPUT=$(pwd$PATH_OUTPUT)
    FILE_JAR="plantuml.jar"
    wget -q -O $FILE_JAR https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
    hasError "Could not get plantuml.jar"
    java -jar $FILE_JAR -charset UTF-8 -output $PATH_OUTPUT "${GITHUB_WORKSPACE}${PATH_PUML}/**.puml"
    hasError "Could not generate png files"
    ls $PATH_OUTPUT
    rm $FILE_JAR
}

# for each in png files and put in markdown
function putEachPngFile() {
    local files_png
    files_png=$(ls "$PATH_OUTPUT" -t -U | grep -oP "^[a-z]+(_[a-z]+)*\.png$")
    hasError "Could not get png files"
    
    for i in $files_png; do
        doMarkdown $i
    done
}

# build markdown
function doMarkdown() {
    local file_path="$PATH_DIAGRAMS_GIT/$1"
    local title=$(getNameToTitle $1)

    echo "## $title" >> Diagrams.md
    echo "![$1]($file_path)" >> Diagrams.md # image
}

# build the header for each diagram
function getNameToTitle() {
    local title=`expr match "$1" '\([a-z_]*\)'` # remove .png
    title=${title//_/ } # replace "_" to blank space
    title=${title^} # first letter to uppercase
    echo $title
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
    git commit -m "$MESSAGE_COMMIT" && git push "https://$GITHUB_ACTOR:$TOKEN@github.com/$GITHUB_REPOSITORY.wiki.git"
}


echo "cloning the wiki repository..."
getWikiRepository
echo "generating png files..."
pumlToPng
# echo "generating the markdown file..."
# putEachPngFile
 echo "configuring git..."
 SetConfigsGit
 echo "starting the function doPush..."
 doPush