Docker build file to create base image file with all tools installed.  
Then you can use `../Dockerfile.fromBaseImage` (or `build.fromBaseImage.sh`) to quickly create grpcpingpong image many times, especially during dev.

Build with command  
```
docker build -t godev .
```
