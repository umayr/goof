# `goof`
> extracts blogs from top web magazines.


## Install
```bash
λ go get github.com/umayr/goof
```

## Usage

```
λ goof -h                                                                                                                

NAME:
   goof - Extracts blogs from top web magazines

USAGE:
   goof [global options] command [command options] [arguments...]
   
VERSION:
   0.0.0
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --page, -p "1"	Tells how many pages need to be extracted
   --plug, -u 		Tells which plug needs to be invoked
   --debug, -d		Turns on the debug mode
   --help, -h		show help
   --version, -v	print the version
```

## Example 
```bash
λ DEBUG=* && MONGO_URI="mongo://localhost" && DB_NAME="goofy-db" && goof --plug tech-crunch --page 10
```

## Setup
```bash
# Make sure you have go 1.5+ and mongodb 3.0+ installed and `GO15VENDOREXPERIMENT` enabled.
# Download glide via brew
λ brew install glide

# Clone this project in <GOPATH>/src directory
λ cd $GOPATH/src && git clone https://github.com/umayr/goof/ && cd $_

# Install dependencies
λ glide install


```
