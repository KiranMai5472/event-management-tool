FROM golang:alpine

RUN mkdir /cms
# Set destination for COPY
WORKDIR /cms

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . /cms

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8000

# Run
CMD ["/cms/main"]















# # FROM golang:1.20

# # WORKDIR /cms

# # COPY go.mod .
# # COPY main.go .

# # RUN go build -o bin .

# # ENTRYPOINT [ "cms/bin" ]







# FROM golang:1.21.3

# # Set destination for COPY
# WORKDIR /cms/

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /cms

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8000

# # Run
# CMD ["go","run","main.go"]















# # # FROM golang:1.20

# # # WORKDIR /cms

# # # COPY go.mod .
# # # COPY main.go .

# # # RUN go build -o bin .

# # # ENTRYPOINT [ "cms/bin" ]


