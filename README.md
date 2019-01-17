# Build a minimal docker image with go

This demonstrates a minimal docker image for a small webserver that
run in a container with an image size of **less than 4,9M !**

It uses statically, sized-optimized go compilation, and an empty docker image
with no OS included (scratch).

Run *runDocker.sh* to compile (by calling *buildStatic.sh* ) and demonstrates the minimal container.

# Cross-compiling for windows

Use *buildWindows.sh* to cross compile a windows binary.

# The demo GO program

The demo program iself is a simple web server, running on localhost:8080.

To stop the server (and the demo), connect to *localhost:8080/quit*
