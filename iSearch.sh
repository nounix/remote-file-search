#!/bin/bash

IP="192.168.200.2"
MOUNT_DIR="/run/user/$(id -u)/gvfs/smb-share:server=$IP,share="
SERVER_DIR="/volume1/"
EXCLUDE_DIRS='/.;/@'
FILEMANAGER="nemo"

iSearch-frontend $IP $MOUNT_DIR $SERVER_DIR $EXCLUDE_DIRS $FILEMANAGER
