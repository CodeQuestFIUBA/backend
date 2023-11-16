FROM golang:1.20

# Set destination for COPY.
WORKDIR /app

# Copy the source code.
COPY ./ ./

# Download Go modules.
RUN go get

# Build.
RUN go build -o /codequest .

# Bind port.
EXPOSE 8080

# Run app.
ENTRYPOINT [ "/codequest" ]
