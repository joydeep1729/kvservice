# Project Title

This is a web application which loads a CSV on startup and then exposes an API to explore the file. 
The API accepts a "key" and returns the corresponding "key, value" pair as JSON

## Getting Started

Simply clone the repository into your local machine and do 'make' and 'make runserver'

### Prerequisites

Tested in Linux having GO (1.10.3 version) and Docker (18.06 version)

### Installing

Install Go 1.10 from the link below:
    https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

Install Docker 18.06 from the link below:
    https://docs.docker.com/install/linux/docker-ce/ubuntu/

### Example

For the following request:
    curl -X GET http://localhost:8080/api/app/v1/kv/one

The Response will be:
    {"key":"one","value":1}

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
