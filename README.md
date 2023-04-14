---
title: Terraform State Inspector
tags: terraform, state, inspector, tool, json, parse, extract, metadata, outputs, resources
author: Dan Fedick
---

Terraform State Inspector (TSI) is a command-line tool that helps you parse and extract useful information from your Terraform state files. It allows you to retrieve state file metadata, outputs, and the number of resources under management.

## Download Instructions

To download the TSI tool, you can either clone the repository and build or download the binary from the releases page.  

## Release: 
The [releases page](https://github.com/demoland/tsi/releases) has the latest version of the binaries, build for darwin  and linux.

**Example:**
* Browse to: https://github.com/demoland/tsi/releases
* Download the tsi-darwin binary: 

```bash
cp tsi-darwin /usr/local/bin/
chmod +x /usr/local/bin/tsi
```

This will install the tsi binary  executable  called tsi in the current directory, then copy to your executable path and enable execution permission.

## Run Commands

To run the TSI tool, use the following syntax:

```bash
./tsi [flags]
```

Here's a list of available flags:

* `-file (string):` Specify the JSON file to parse.

* `-meta` Print statefile metadata.

* `-outputs` Print outputs.

* `-resources` Print the number of resources under management.

**Examples**
To print the statefile metadata:

```bash
./tsi -file my-state-file.json -meta
```

To print the state outputs:

```bash
./tsi -file my-state-file.json -outputs
```

To print the data and resources and the current number of resources under management that are managed in statefile:

```bash
./tsi -file my-state-file.json -resources
```

**Getting Help**
To display the help message, run:

```bash
./tsi -h
```

This will show the usage instructions and the available flags.
