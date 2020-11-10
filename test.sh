#!/bin/bash
go build
for bad in ./examples/bad*
do
    echo "$bad" - $(./tetris-optimizer "$bad")
   
done
 echo
for good in ./examples/inp*
do
    echo -e "$good" : 
    ./tetris-optimizer "$good"
    echo
done
echo "./hardinput :" 
time ./tetris-optimizer ./examples/hardinput