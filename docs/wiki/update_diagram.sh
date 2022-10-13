#!/bin/bash 

if [ -z "$PATH_DIAGRAMS" ]; then
    echo "PATH_DIAGRAMS is not specified"
    exit 1
fi
if [ -z "$TOKEN" ]; then
    echo "TOKEN is not specified"
    exit 1
fi

# The pattern for PATH_DIAGRAMS is "/folder/folder"
if [[ "$PATH_DIAGRAMS" =~ ^\.\/ ]]; then
    PATH_DIAGRAMS=${PATH_DIAGRAMS:1} # remove the first char
elif [[ ! "$PATH_DIAGRAMS" =~ ^\/ ]]; then
    PATH_DIAGRAMS="/$PATH_DIAGRAMS"
fi

PATH_DIAGRAMS_LOCAL="$PATH_DIAGRAMS"
PATH_DIAGRAMS_GIT="https://github.com/${GITHUB_REPOSITORY}/blob/main$PATH_DIAGRAMS"
MESSAGE_COMMIT=`git log -1 --format="%s"`

# Stop the execution if has an error
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
    
    #Update path local, remove the username in GITHUB_REPOSITORY
    PATH_DIAGRAMS_LOCAL=..$(echo "$GITHUB_REPOSITORY" | grep '/[[:alnum:]_\.\-]\+$' --only-matching)"$PATH_DIAGRAMS"
    
    # remove old file
    rm Diagrams.md
}

# for each in png file and put in markdown
function putEachPngFile() {
    local files_png
    files_png=$(ls "$PATH_DIAGRAMS_LOCAL" -t -U | grep -oP "^[a-z]+(_[a-z]+)*\.png$")
    hasError "Could not get png files"
    
    for i in $files_png; do
        doMarkdown $i
    done
}

# build markdown
function doMarkdown() {
    local file_path="$PATH_DIAGRAMS_GIT/$1"
    local title=$(getNameToNewFile $1)

    echo "## $title" >> Diagrams.md
    echo "![$1]($file_path)" >> Diagrams.md # image
}

# build the header for each diagram, 
function getNameToTitle() {
    local title=`expr match "$1" '\([a-z_]*\)'` # remove .png
    title=${title//_/ } # replace _ to blank space
    title=${title^} # first letter uppercase
    echo title
}

function doPush() {
    if [ -z "$(git status --porcelain)" ]; then 
        echo "there no files to changed"
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
echo "generating the markdown file..."
putEachPngFile
echo "configuring git..."
SetConfigsGit
echo "starting the function doPush..."
doPush