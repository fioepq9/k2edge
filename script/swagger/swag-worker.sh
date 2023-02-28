  docker run --rm -v "$(pwd)/worker:/app"  swaggerapi/swagger-codegen-cli generate \
    -i "/app/swagger/worker.json" \
    -l "go" \
    -o "/app/client"