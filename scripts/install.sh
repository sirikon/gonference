#!/bin/sh

GO111MODULE=off go get -u github.com/gobuffalo/packr/v2/packr2
go mod download
npm i -g yarn
cd $FRONT_STYLE_PATH
yarn
cd ../../
cd $BACKOFFICE_UI_PATH
yarn
