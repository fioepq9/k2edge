docker run --privileged=true --rm -v "$(pwd)/master:/app" swaggerapi/swagger-codegen-cli generate \
    -i "/app/swagger/master.json" \
    -l "go" \
    -o "/app/client"