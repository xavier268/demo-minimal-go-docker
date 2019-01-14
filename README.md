# Build a minimal docker image with go

This demonstrates a minimal docker image for a small webserver that
run in a container with an image size of **less than 4,9M !**

It uses statically, sized-optimized go compilation, and an empty docker image
with no OS included (scratch).

To run the demo, run run.sh and connect to localhost:8080/something.
To stop the server, connect to localhost:8080/quit
