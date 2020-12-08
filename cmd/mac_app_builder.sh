#!/bin/bash

# This script can be used to create a "Mac App" complete with icons
# Then, they can sit nicely on the Mac Desktop or on the Dock
#
# Cmd-line argument should be one of the following:
# sqscopy, sqspaste, or sqspurge

PGM=$1

cd $1
rm -rf ~/Desktop/${PGM}.app/
rm -rf ./assets/
rm -rf ${PGM}.app/
mkdir ./assets
rm -f ${PGM}
go build -ldflags="-s -w"
sync ; sleep 3 ; sync
cp ${PGM} ./assets/
sync ; sleep 3 ; sync
# macapp:https://gist.github.com/jftuga/b3ec5a66472c0aec5676bfd7b90a1909
macapp -bin ${PGM} -icon ./appicon1024.png -identifier github.jftuga.sqs_clipboard -name ${PGM} -assets ./assets/
sync ; sleep 3 ; sync
mv ${PGM}.app/ ~/Desktop/
