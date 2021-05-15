FROM alpine
WORKDIR /
COPY mutation-linux /mutation

# Document that the service listens on port 80.
EXPOSE 80

# Run the outyet command by default when the container starts.
ENTRYPOINT /mutation
