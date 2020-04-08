FROM golang:latest
# Create WORKDIR (working directory) for app
EXPOSE 8000
WORKDIR /app
# Copy the local package files to the container's workspace
# (in the above WORKDIR)
ADD . .
# Switch WORKDIR to directory where server main.go lives
WORKDIR /app
# Build the go-API-template userServer command inside the container
# at the most recent WORKDIR
RUN go build -o main


# Run the userServer command by default when the container starts.
# runs command at most recent WORKDIR
ENTRYPOINT ./main
