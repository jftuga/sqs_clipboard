#!/bin/bash

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
# convert Button-Download-icon.png -resize 1024x1024 appicon1024.png
# macapp: https://gist.github.com/mholt/11008646c95d787c30806d3f24b2c844
macapp -bin ${PGM} -icon ./appicon1024.png -identifier github.jftuga.sqs_clipboard -name ${PGM} -assets ./assets/
sync ; sleep 3 ; sync
mv ${PGM}.app/ ~/Desktop/
