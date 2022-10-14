#!/bin/bash 

if [ -z "$TOKEN" ]; then
    echo "TOKEN is not specified"
    exit 1
fi
if [ -z "$PAGE_NAME" ]; then 
    echo "PAGE_NAME is no specified"
    exit 1
fi
if [ -z "$PATH_DIAGRAMS" ] || [[ "$PATH_DIAGRAMS" =~ ^[\.\/] ]]; then
    echo "PATH_DIAGRAMS must follow the established pattern"
    exit 1 
fi
if [ -z "$OUTPUT_DIAGRAMS" ] || [[ "$OUTPUT_DIAGRAMS" =~ ^[\.\/] ]]; then
    echo "OUTPUT_DIAGRAMS must follow the established pattern"
    exit 1
fi

cd .. # return to root path

# Put "/"
OUTPUT_DIAGRAMS="/$OUTPUT_DIAGRAMS"
PATH_DIAGRAMS="/$PATH_DIAGRAMS"

# get absolute path
ROOT_OUTPUT_DIAGRAMS="$(pwd)"$OUTPUT_DIAGRAMS

# The default name for the wiki repository.
TEMP_REPO_NAME="wiki-repo" 

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
    local email=`git log -1 --format="%ae"`
    echo $email
    # set configs git
    git config --global user.email "$email"
    hasError "Could not config git"
    git config --global user.name "$GITHUB_ACTOR"
    hasError "Could not config git"
}

# Clone the wiki repository and change working directory to wiki repository
function getWikiRepository() {
    hasError "Could not return to root path"

    git clone "https://$GITHUB_ACTOR:$TOKEN@github.com/$GITHUB_REPOSITORY.wiki.git" "$TEMP_REPO_NAME"
    hasError "Could not clone repo"
    
    # move to wiki repository
    cd "$TEMP_REPO_NAME"
    hasError "Could not move to wiki repository"
}

# Function to generate the png files
function pumlToPng() {
    # default name to plantuml
    FILE_JAR="plantuml.jar" 

    wget -q -O $FILE_JAR https://github.com/plantuml/plantuml/releases/download/v1.2022.8/plantuml-1.2022.8.jar
    hasError "Could not get plantuml.jar"

    java -q -jar $FILE_JAR -charset UTF-8 -output $ROOT_OUTPUT_DIAGRAMS "${GITHUB_WORKSPACE}${PATH_PUML}/**.puml"
    hasError "Could not generate png files"
    
    rm $FILE_JAR
}

# for each in png files and put in markdown
function putEachPngFile() {
    # remove old file
    rm "$PAGE_NAME"

    # get png files that are using snake_case 
    local files_png
    files_png=$(ls "$ROOT_OUTPUT_DIAGRAMS" -t -U | grep -oP "^[a-z]+(_[a-z]+)*\.png$")
    hasError "Could not get png files"
    
    for file in $files_png; do
        doMarkdown $file
    done
}
# build markdown for a file
function doMarkdown() {
    local file="[[$OUTPUT_DIAGRAMS/$1|alt=$1]]"
    local title=$(getTitle $1)

    echo "## $title" >> "$PAGE_NAME"
    echo "$file" >> "$PAGE_NAME" # image
}
# build the title for each diagram
function getTitle() {
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
echo "configuring git..."
SetConfigsGit
echo "generating png files..."
pumlToPng
echo "generating the markdown file..."
putEachPngFile
echo "starting the function doPush..."
doPush