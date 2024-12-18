# Whoami-go: A Simple Public IP Lookup Tool

Whoami-go helps you easily retrieve the public IP address of your device, either in plain text or JSON format.

## Features
- Quickly display the public IP address of your current device.
- Output available in plain text or JSON format.
- Lightweight, runs efficiently in a Docker container.

## Prerequisite
- A device or server with a public IP address (e.g., a VPC or any other Linux device connected to the internet).  
- **Docker** installed on the device to run the application.

## Usage

### 1. Install via Docker

```bash
docker pull nealhallyoung/whoami-go:0.0.1
docker run -d -p 80:80 --name=whoami-1 nealhallyoung/whoami-go:0.0.1
```

### 2. Lookup

```bash
curl http://[SERVER_IP]:80
curl http://[SERVER_IP]:80/json
```

Example Output:

```
# text
123.45.67.89

# json
{
  "ip": "123.45.67.89"
}

```

You can also open the following URL in your browser to view the IP address:
```
http://[SERVER_IP]:80
```

## License

This project is licensed under the [MIT](LICENSE) License.