# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

# Download required GO modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download 

# Copy all GO files to image and build image
COPY . ./
RUN go build -o /docker-social-golang 

# Export port 
EXPOSE 8080

# Run built image
CMD [ "/docker-social-golang" ]
