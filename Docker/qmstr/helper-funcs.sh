#!/bin/bash

function multiecho(){
    str=$1
    num=$2
    v=$(printf "%-${num}s" "$str")
    echo "${v// /${str}}"
}

function printheader(){
    str=$1
    echo
    echo "$str"
    multiecho "=" ${#str}
}

function quit(){
    curl http://localhost:9000/quit
}

function dump(){
    curl http://localhost:9000/dump
}

function printtargets(){
    curl http://localhost:9000/linkedtargets
}

function build_cmake(){
    cd /build
    rm -fr build
    mkdir build
    cd build
    cmake ..
    make
}

function build_qmstr(){
    PROGS="qmstr-prototype/qmstr/qmstr-master qmstr-prototype/qmstr/qmstr-wrapper"
    for p in $PROGS; do
        printheader "Building $p"
        go build -i -v $p
        printheader "Testing $p"
        go test -i -v $p
        printheader "Installing $p"
        go install -v $p
    done
}
