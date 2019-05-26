#!/bin/sh

go get -u github.com/gobuffalo/packr/v2/packr2
go get
cd $FRONT_STYLE_PATH
npm i
cd ../../
cd $BACKOFFICE_UI_PATH
npm i
