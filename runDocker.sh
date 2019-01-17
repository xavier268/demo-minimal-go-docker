#!/bin/bash
# (c) Xavier Gandillot - 2019
echo "Running a docker container from the static binary"


# Ensure Docker is available
docker info >> /dev/null
if [ $? -ne 0 ] ; then
  echo "Docker does not seem to be available on your system ?"
  exit 
fi
#Check if static binary already exists. Compile if necessary.
if [ ! -f bin/main.static ] ; then
  ./buildStatic.sh
fi

## Building the docker image
IMAGE=$( docker build -q . )
SIZE=$( docker image inspect $IMAGE --format='{{.Size}}' )
SIZE=$(( $SIZE / 1000 ))
echo "Just build docker image : $IMAGE"
echo "Docker image size is : $SIZE Kb"

## Running the container
echo "Running a docker container"
docker run --rm -it -p 8080:8080 $IMAGE

## remove IMAGE
echo "Cleaning up"
docker rmi $IMAGE
echo "Done."
