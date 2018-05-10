#!/bin/bash

node-sass --output-style compressed "`dirname "$0"`/skeleton.scss" > "`dirname "$0"`/../skeleton.css"