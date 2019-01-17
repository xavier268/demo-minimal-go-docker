# Build a minimal docker image with go

This demonstrates a minimal docker image for a small webserver that
run in a container with an image size of **less than 4,9M !**

It uses statically, sized-optimized go compilation, and an empty docker image with no OS included (scratch).

Run *runDocker.sh* launch and test the minimal container.

It will compile a static version of the file, if needed using *buildStatic.sh*. I

It will then create a temporary docker image, dispaly its size, run a container, and cleanup everything on exiting.

# Static compilation for Linux

Use *buildStatic.sh* to compile a static, stripped-down binary for linux (and Docker).

# Cross-compiling for windows

Use *buildWindows.sh* to cross compile a windows binary.

# The demo GO program

The demo program iself is a simple web server, running on localhost:8080.

To stop the server (and the demo), connect to *localhost:8080/quit*
