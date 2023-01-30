docker build -t multithread-test-go . && \
    docker image rm -f $(docker images -f dangling=true -q)