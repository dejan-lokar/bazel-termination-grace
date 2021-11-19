#!/bin/bash

trap 'echo "Trapped: INT or TERM"' INT TERM
echo "Waiting for a signal... PID: $$"

while :
do
  sleep 1
done
