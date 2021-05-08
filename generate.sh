#!/usr/bin/env bash

cd sample

# cat $FILE | sanctify -p model -t $filename > $output
for FILE in *; 
do
  filename="${FILE%.*}"
  output="../model/$filename.go";
  echo $output; 
  # generate
  cat $FILE | sanctify -p model -t $filename > $output
done

cd ..