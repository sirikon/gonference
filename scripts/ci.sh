#!/bin/sh

cd src/backoffice-ui
npm i && npm run lint && npm run build
cd ../..
tsk build
