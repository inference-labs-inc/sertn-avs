#!/usr/bin/env bash
set -e


BASE_DIR=${XDG_CONFIG_HOME:-$HOME}
EZKL_DIR=${EZKL_DIR-"$BASE_DIR/.ezkl"}

# Create the .ezkl bin directory if it doesn't exit
mkdir -p $EZKL_DIR

# Store the correct profile file (i.e. .profile for bash or .zshenv for ZSH).
case $SHELL in
*/zsh)
    PROFILE=${ZDOTDIR-"$HOME"}/.zshenv
    PREF_SHELL=zsh
    ;;
*/bash)
    PROFILE=$HOME/.bashrc
    PREF_SHELL=bash
    ;;
*/fish)
    PROFILE=$HOME/.config/fish/config.fish
    PREF_SHELL=fish
    ;;
*/ash)
    PROFILE=$HOME/.profile
    PREF_SHELL=ash
    ;;
*)
    echo "NOTICE: Shell could not be detected, you will need to manually add ${EZKL_DIR} to your PATH."
esac

# Check for non standard installation of ezkl
if  [ "$(which ezkl)s" != "s" ] && [ "$(which ezkl)" != "$EZKL_DIR/ezkl" ] ; then
    echo "ezkl is installed in a non-standard directory, $(which ezkl). To use the automated installer, remove the existing ezkl from path to prevent conflicts"
    exit 1
fi

if [[ ":$PATH:" != *":${EZKL_DIR}:"* ]]; then
    # Add the ezkl directory to the path and ensure the old PATH variables remain.
    echo >> $PROFILE && echo "export PATH=\"\$PATH:$EZKL_DIR\"" >> $PROFILE
fi

# Install latest ezkl version
# Get the right release URL
if [ -z "$1" ]
then
    RELEASE_URL="https://api.github.com/repos/zkonduit/ezkl/releases/latest"
    echo "No version tags provided, installing the latest ezkl version"
else
    RELEASE_URL="https://api.github.com/repos/zkonduit/ezkl/releases/tags/$1"
    echo "Installing ezkl version $1"
fi

PLATFORM=""
case "$(uname -s)" in

    Darwin*)
        PLATFORM="macos"
        ;;

    Linux*Microsoft*)
        PLATFORM="linux"
        ;;

    Linux*)
        PLATFORM="linux"
        ;;

    CYGWIN*|MINGW*|MINGW32*|MSYS*)
        PLATFORM="windows-msvc"
        ;;

    *)
        echo "Platform is not supported. If you would need support for the platform please submit an issue https://github.com/zkonduit/ezkl/issues/new/choose"
        exit 1
        ;;
esac

# Check arch
ARCHITECTURE="$(uname -m)"
if [ "${ARCHITECTURE}" = "x86_64" ]; then
    # Redirect stderr to /dev/null to avoid printing errors if non Rosetta.
    if [ "$(sysctl -n sysctl.proc_translated 2>/dev/null)" = "1" ]; then
        ARCHITECTURE="arm64" # Rosetta.
    else
        ARCHITECTURE="amd64" # Intel.
    fi
elif [ "${ARCHITECTURE}" = "arm64" ] ||[ "${ARCHITECTURE}" = "aarch64" ]; then
    ARCHITECTURE="aarch64" # Arm.
elif [ "${ARCHITECTURE}" = "amd64" ]; then
    ARCHITECTURE="amd64" # Amd
else
    echo "Architecture is not supported. If you would need support for the architecture please submit an issue https://github.com/zkonduit/ezkl/issues/new/choose"
    exit 1
fi

# Remove existing ezkl
echo "Removing old ezkl binary if it exists"
[ -e file ] && rm file

# echo platform and architecture
echo "Platform: $PLATFORM"
echo "Architecture: $ARCHITECTURE"

# download the release and unpack the right tarball
if [ "$PLATFORM" == "windows-msvc" ]; then
    JSON_RESPONSE=$(curl -s "$RELEASE_URL")
    FILE_URL=$(echo "$JSON_RESPONSE" | grep -o 'https://github.com[^"]*' | grep "build-artifacts.ezkl-windows-msvc.tar.gz")

    echo "Downloading package"
    curl -L "$FILE_URL" -o "$EZKL_DIR/build-artifacts.ezkl-windows-msvc.tar.gz"

    echo "Unpacking package"
    tar -xzf "$EZKL_DIR/build-artifacts.ezkl-windows-msvc.tar.gz" -C "$EZKL_DIR"

    echo "Cleaning up"
    rm "$EZKL_DIR/build-artifacts.ezkl-windows-msvc.tar.gz"

elif [ "$PLATFORM" == "macos" ]; then
    if [ "$ARCHITECTURE" == "aarch64" ] || [ "$ARCHITECTURE" == "arm64" ]; then
        JSON_RESPONSE=$(curl -s "$RELEASE_URL")
        FILE_URL=$(echo "$JSON_RESPONSE" | grep -o 'https://github.com[^"]*' | grep "build-artifacts.ezkl-macos-aarch64.tar.gz")

        echo "Downloading package"
        curl -L "$FILE_URL" -o "$EZKL_DIR/build-artifacts.ezkl-macos-aarch64.tar.gz"

        echo "Unpacking package"
        tar -xzf "$EZKL_DIR/build-artifacts.ezkl-macos-aarch64.tar.gz" -C "$EZKL_DIR"

        echo "Cleaning up"
        rm "$EZKL_DIR/build-artifacts.ezkl-macos-aarch64.tar.gz"
    else
        JSON_RESPONSE=$(curl -s "$RELEASE_URL")
        FILE_URL=$(echo "$JSON_RESPONSE" | grep -o 'https://github.com[^"]*' | grep "build-artifacts.ezkl-macos.tar.gz")

        echo "Downloading package"
        curl -L "$FILE_URL" -o "$EZKL_DIR/build-artifacts.ezkl-macos.tar.gz"

        echo "Unpacking package"
        tar -xzf "$EZKL_DIR/build-artifacts.ezkl-macos.tar.gz" -C "$EZKL_DIR"

        echo "Cleaning up"
        rm "$EZKL_DIR/build-artifacts.ezkl-macos.tar.gz"

    fi

elif [ "$PLATFORM" == "linux" ]; then
    if [ "$ARCHITECTURE" == "amd64" ]; then
        JSON_RESPONSE=$(curl -s "$RELEASE_URL")
        FILE_URL=$(echo "$JSON_RESPONSE" | grep -o 'https://github.com[^"]*' | grep "build-artifacts.ezkl-linux-gnu.tar.gz")

        echo "Downloading package"
        curl -L "$FILE_URL" -o "$EZKL_DIR/build-artifacts.ezkl-linux-gnu.tar.gz"

        echo "Unpacking package"
        tar -xzf "$EZKL_DIR/build-artifacts.ezkl-linux-gnu.tar.gz" -C "$EZKL_DIR"

        echo "Cleaning up"
        rm "$EZKL_DIR/build-artifacts.ezkl-linux-gnu.tar.gz"
    elif [ "$ARCHITECTURE" == "aarch64" ]; then
        JSON_RESPONSE=$(curl -s "$RELEASE_URL")
        FILE_URL=$(echo "$JSON_RESPONSE" | grep -o 'https://github.com[^"]*' | grep "build-artifacts.ezkl-linux-aarch64.tar.gz")

        echo "Downloading package"
        curl -L "$FILE_URL" -o "$EZKL_DIR/build-artifacts.ezkl-linux-aarch64.tar.gz"

        echo "Unpacking package"
        tar -xzf "$EZKL_DIR/build-artifacts.ezkl-linux-aarch64.tar.gz" -C "$EZKL_DIR"

        echo "Cleaning up"
        rm "$EZKL_DIR/build-artifacts.ezkl-linux-aarch64.tar.gz"
    else
        echo "Non aarch ARM architectures are not supported for Linux at the moment. If you would need support for the ARM architectures on linux please submit an issue https://github.com/zkonduit/ezkl/issues/new/choose"
        exit 1
    fi
else
    echo "Platform and Architecture is not supported. If you would need support for the platform and architecture please submit an issue https://github.com/zkonduit/ezkl/issues/new/choose"
    exit 1
fi


echo && echo "Successfully downloaded ezkl at ${EZKL_DIR}"
echo "We detected that your preferred shell is ${PREF_SHELL} and added ezkl to PATH. Run 'source ${PROFILE}' or start a new terminal session to use ezkl."
