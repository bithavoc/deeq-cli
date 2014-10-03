#!/usr/bin/env bash

shopt -s extglob
set -o errtrace
set -o errexit

DEEQ_DIST="unknown"
DEEQ_ROOT_DIR="$HOME/.deeq"
DEEQ_BIN_DIR="$DEEQ_ROOT_DIR/bin"
DEEQ_BIN_PATH="$DEEQ_BIN_DIR/deeq"

deeq_step()
{
    echo "-> $1"
}

deeq_error()
{
    echo "!!! ERROR: $1"
    echo "Go to http://deeqapp.com for help"
    exit 1
}

deeq_ensure_dir()
{
    deeq_step "Ensuring $1"
    mkdir -p $1
}

deeq_file_exectuable()
{
    deeq_step "$1 is a new executable"
    chmod +x $1
}

deeq_tool_available() {
    hash "$1" > /dev/null 2>&1
    return $?
}

#
# Downloads a file and puts it in the file system
#
# Arguments:
# 1. First argument is description of operation.
# 2. URL of the file to download
# 3. fs destination path
deeq_download()
{
    deeq_step "Downloading Deeq executable for $DEEQ_DIST"
    DEEQ_BIN_URL="http://dl.bithavoc.io/deeq/releases/deeq-$DEEQ_DIST-latest"
    if deeq_tool_available "curl"; then
        # is curl available(OSX)?
        deeq_step "Using curl to download $DEEQ_BIN_URL"
        curl -SL# $DEEQ_BIN_URL -o $DEEQ_BIN_PATH
    else
        # maybe WGET(Linux)?
        if deeq_tool_available "wget"; then
            deeq_step "Using wget to download $DEEQ_BIN_URL"
            wget --progress=bar $DEEQ_BIN_URL -O $DEEQ_BIN_PATH
        else
            # we're fucked
            deeq_error "Neither curl or wget were found"
        fi
    fi
    chmod +x $DEEQ_BIN_PATH
}

deeq_determinate_OS()
{
    if [[ "$OSTYPE" == "darwin"* ]]; then
        DEEQ_DIST="osx"
    elif [[ "$OSTYPE" == "linux-gnu" ]]; then
        DEEQ_DIST="linux"
    else
        deeq_error "Unable to determinate Operating System, it's very unlikely that Deeq is supposed in this OS"
    fi
}

deeq_detect_profile()
{
    # Detect profile file if not specified as environment variable (eg: PROFILE=~/.myprofile).
    if [ -z "$PROFILE" ]; then
      if [ -f "$HOME/.zshrc" ]; then
        PROFILE="$HOME/.zshrc"
      elif [ -f "$HOME/.bashrc" ]; then
        PROFILE="$HOME/.bashrc"
      elif [ -f "$HOME/.bash_profile" ]; then
        PROFILE="$HOME/.bash_profile"
      elif [ -f "$HOME/.profile" ]; then
        PROFILE="$HOME/.profile"
      fi
    fi
}

deeq_ensure_PATH()
{
    deeq_detect_profile
    if deeq_tool_available "deeq"; then
        #already in path, do nothing
        deeq_step "Deeq executable already in PATH (detected from $PROFILE)"
    else
        # deeq is not in path, let's add it
        deeq_step "Adding executable to PATH in $PROFILE"
        echo "export PATH=$PATH:$DEEQ_BIN_DIR" >> $PROFILE
        deeq_step "Reloading $PROFILE"
        source $PROFILE
    fi
}

deeq_show_version()
{
    deeq_step "Deeq $(deeq version -s) installed successfully"
    echo
}

deeq_bye() {
    deeq_step "Deeq is now correctly installed, run this command again to get the latest version"
    deeq about
}

deeq_step "Installing Deeq"
deeq_determinate_OS
deeq_ensure_dir $DEEQ_ROOT_DIR
deeq_ensure_dir $DEEQ_BIN_DIR
deeq_download
deeq_file_exectuable $DEEQ_BIN_PATH
deeq_ensure_PATH
deeq_show_version
deeq_bye
