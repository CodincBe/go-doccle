# README

Originally forked from [jonaswouters/go-doccle](https://github.com/jonaswouters/go-doccle), all kudos to him
for having the information on the API already identified.

A simple set up that potentially will not need to be altered much more unless new use cases are being added.

## Simple usage example

Build the package and make sure that there's a config.json file located in the same directory containing
your username and password for Doccle.

The following options are supported:
 - **save-path:** absolute path to the directory you want to save the listed documents.
 - **only-new:** only retrieve the documents from Doccle that are _new_
 - **archive:** the retrieved documents will be archived, if they support this (already archived will not have this option)

## Contributions

Contributions are always welcomed.
