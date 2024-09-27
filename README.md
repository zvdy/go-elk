# ELK Golang 

## Overview

This project sets up an ELK (Elasticsearch, Logstash, Kibana) stack using Docker Compose. It includes a Go application `log_generator` that generates random logs and sends them to Elasticsearch. The logs are then visualized using Kibana.

## Prerequisites

- Docker
- Docker Compose

## Project Structure

- [`docker-compose.yml`]: Defines the services for Elasticsearch, Logstash, Kibana, and the log generator.
- [`Dockerfile`]: Builds the Docker image for the Go log generator.
- [`log_generator.go`]: Go application that generates random logs and sends them to Elasticsearch.
- [`logstash.conf`]: Configuration file for Logstash to process and forward logs to Elasticsearch.

## Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/zvdy/go-ELK
    cd go-ELK
    ```

2. **Build and start the services**:
    ```sh
    docker-compose up --build
    ```

3. **Verify Elasticsearch**:
    Access Elasticsearch at [`http://localhost:9200`].

4. **Access Kibana**:
    Access Kibana at [`http://localhost:5601`].

## Files

Sure, here's an explanation of the functionality of each file mentioned in the [`README.md`]:

### [`docker-compose.yml`]
This file defines the services required to set up the ELK stack and the Go log generator using Docker Compose. It includes configurations for:

- **Elasticsearch**: Runs an Elasticsearch container, exposing port 9200.
- **Kibana**: Runs a Kibana container, exposing port 5601 and connecting to Elasticsearch.
- **Logstash**: Runs a Logstash container, configured to process logs and forward them to Elasticsearch.
- **log_generator**: Builds and runs the Go application that generates random logs and sends them to Elasticsearch.

### [`Dockerfile`]
This file is used to build the Docker image for the Go log generator application. It specifies:

- The base image (`golang:1.23`).
- The working directory inside the container (`/app`).
- Copying the application files into the container.
- Building the Go application.
- The command to run the application (`./log_generator`).

### [`log_generator.go`]
This Go application generates random logs and sends them to Elasticsearch. Key functionalities include:

- **Generating Random Logs**: Creates logs with random levels (INFO, WARN, ERROR) and messages.
- **Retry Mechanism**: Waits for Elasticsearch to be ready before sending logs.
- **Sending Logs**: Sends the generated logs to Elasticsearch in JSON format at regular intervals (every second).

### [`logstash.conf`]
This configuration file for Logstash specifies:

- **Input**: Reads logs from a file ([`/logs/random_log.yaml`]).
- **Filter**: Uses Grok patterns to parse the log messages.
- **Output**: Forwards the parsed logs to Elasticsearch and prints them to the console.

### Usage
- **Generate Logs**: The [`log_generator`] service starts automatically and generates logs.
- **View Logs in Kibana**: Access Kibana to visualize the logs by configuring an index pattern for [`random_logs`].

### Troubleshooting
- **Elasticsearch Connection Issues**: Ensure Elasticsearch is running and accessible. The [`log_generator`] includes a retry mechanism.
- **Logstash Configuration**: Ensure the [`logstash.conf`] file is correctly mapped and configured.

### License
The project is licensed under the MIT License.

## Usage

1. **Generate Logs**:
    The [`log_generator`] service will automatically start generating logs and sending them to Elasticsearch.

2. **View Logs in Kibana**:
    Open Kibana at [`http://localhost:5601`] and configure an index pattern for [`random_logs`] to visualize the logs.

## Troubleshooting

- **Elasticsearch Connection Issues**:
    Ensure Elasticsearch is running and accessible. The [`log_generator`] includes a retry mechanism to wait for Elasticsearch to be ready.

- **Logstash Configuration**:
    Ensure the [`logstash.conf`] file is correctly mapped and configured to process logs.

## License

This project is licensed under the MIT License. See the LICENSE file for details.