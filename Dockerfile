## The following are used for the "big" version
# for now we'll simply go ahead, download the file, and build.

# # Build the manager binary
# FROM golang:1.22-alpine as builder

# WORKDIR /workspace
# # Copy the Go Modules manifests
# COPY siteclean/go.mod go.mod
# # COPY siteclean/go.sum go.sum
# # cache deps before building and copying source so that we don't need to re-download as much
# # and so that source changes don't invalidate our downloaded layer
# RUN go mod download

# # Copy the go source
# COPY siteclean/main.go main.go


# # Build
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -o siteclean main.go

# FROM uselagoon/nginx:24.5.0

# COPY --from=builder /workspace/siteclean /siteclean

# COPY siteclean/downloader/downloader.sh /downloader.sh

# "little" version for demo here

FROM uselagoon/nginx:latest
ARG STATIC_FILE_LOCATION=https://github.com/bomoko/lagoon-mini-site/raw/main/default.zip
RUN apk add wget
ENV STATIC_FILE_LOCATION=${STATIC_FILE_LOCATION}

RUN wget -O /tmp/site.zip ${STATIC_FILE_LOCATION} && unzip /tmp/site.zip -d /app
