tag=$(git rev-parse --short HEAD)
docker build . -t gin-example:${tag} -f ./docker/Dockerfile