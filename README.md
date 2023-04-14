---
title: Terraform State Inspector
tags: terraform, state, inspector, tool, json, parse, extract, metadata, outputs, resources
author: Dan Fedick
---

Terraform State Inspector (TSI) is a command-line tool that helps you parse and extract useful information from your Terraform state files. It allows you to retrieve state file metadata, outputs, and the number of resources under management.

## Download Instructions

To download the TSI tool, you can either clone the repository or download the zip file from the repository page.

## Cloning the repository

If you have git installed, you can clone the repository using the following command:

```bash
git clone https://github.com/yourusername/terraform-state-inspector.git
```

## Downloading the zip file

Visit the TSI [repository page](https://github.com/demoland/tsi.git) and click on the green "Code" button. Select "Download ZIP" from the dropdown menu and extract the contents to your desired location.

## Build Instructions

To build the TSI tool, you need to have the Go language installed on your system. Follow the instructions below to build the tool:

**Navigate to the TSI directory:**

```bash
cd terraform-state-inspector
```

**Build the binary:**

```bash
go build *.go -o tsi
cp tsi /usr/local/bin
chmod +x /usr/local/bin/tsi
```

This will create an executable binary called tsi in the current directory, then copy to your executable path and enable execution permission.

## Run Commands

To run the TSI tool, use the following syntax:

```bash
./tsi [flags]
```

Here's a list of available flags:

`-file (string):` Specify the JSON file to parse.
`-meta` Print statefile metadata.
`-outputs` Print outputs.
`-resources` Print the number of resources under management.

**Examples**
To print the statefile metadata:

```bash
./tsi -file my-state-file.json -meta
```

To print the outputs:

```bash
./tsi -file my-state-file.json -outputs
```

To print the number of resources under management:

```bash
./tsi -file my-state-file.json -resources
```

**Getting Help**
To display the help message, run:

```bash
./tsi -h
```

This will show the usage instructions and the available flags.
