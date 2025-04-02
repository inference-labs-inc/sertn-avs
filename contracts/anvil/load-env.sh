#!/bin/bash

# Get the directory of the current script
SCRIPT_DIR=$(dirname "$(realpath "$0")")

# Check if .env file exists
if [ -f "$SCRIPT_DIR/../.env" ]; then
  # Export all variables from .env
  export $(grep -v '^#' $SCRIPT_DIR/../.env | xargs)
else
  echo ".env file not found. Please create one."
  exit 1
fi
