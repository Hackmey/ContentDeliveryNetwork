# Content Delivery Network (CDN) Implementation

![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)
![Redis](https://img.shields.io/badge/Redis-7.0+-red.svg)
![Docker](https://img.shields.io/badge/Docker-24.0+-blue.svg)

A distributed CDN solution designed to optimize web content delivery through geographic edge servers and caching.

## Project Overview

This CDN implementation enhances website performance by:
- Reducing latency through geographic distribution
- Improving load times by 50-70%
- Providing redundancy and DDoS protection
- Utilizing Redis caching with automatic content expiration
- Implementing load balancing across edge servers

## Key Components

### 🖥 Origin Server (Port 9090)
- Serves core HTML content via Go + Gorilla Mux
- Pages: 
  - `/main.html` - Home page
  - `/arch.html` - System architecture
  - `/desc.html` - CDN benefits documentation
- Acts as content source for edge servers

### 🌐 Edge Servers
| Location | Port | Features |
|----------|------|----------|
| **US**   | 8080 | - Redis caching<br>- Latency monitoring<br>- Auto-redirect from `/` |
| **EU**   | 8090 | - Cache synchronization<br>- TTL management (10min)<br>- Request metrics |

## Getting Started

### Prerequisites
- Install [Go](https://go.dev/doc/install) (v1.21+)
- Install [Redis](https://redis.io/docs/getting-started/) (v7.0+)
- Install [Docker](https://docs.docker.com/get-docker/) (v24.0+)

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/JasminPradhan/ContentDeliveryNetwork.git
    cd ContentDeliveryNetwork
    ```

2. Start the origin server:
    ```bash
    go run origin/main.go
    ```

3. Start the edge servers:
    ```bash
    cd edge-eu
    docker build -t cdn-edge-eu .
    docker run -p 8080:8090 cdn-edge-eu
    ```
    ```bash
    cd edge-us
    docker build -t cdn-edge-us .
    docker run -p 8060:8080 cdn-edge-us
    ```

4. Access the CDN:
    - US Edge Server: `http://localhost:8080`
    - EU Edge Server: `http://localhost:8090`


## System Architecture

```mermaid
graph TD
    User -->|Request| Edge
    Edge -->|Cache Hit| User
    Edge -->|Cache Miss| Origin
    Origin -->|Content + Headers| Edge
    Edge -->|Cache & Deliver| User

```

## Demonstration

### Cache-Aware Delivery System
- **Automatic Source Selection**: 
  ```http
  GET /main.html HTTP/1.1
  X-Cache-Status: hit (US Edge) | miss (Origin)


![Screenshot (148)](https://github.com/user-attachments/assets/c7b61db1-4e49-47d6-9df5-49790cf451eb)
![Screenshot (147)](https://github.com/user-attachments/assets/20e24c3e-2289-445f-967e-34ff7e7f1be2)
![Screenshot (146)](https://github.com/user-attachments/assets/4f5ddb58-3eb7-4c4a-a896-dda19e6dbbf0)
![Screenshot (145)](https://github.com/user-attachments/assets/fcdf32cf-b2a0-472e-a81e-918af2c5edb1)
![Screenshot (144)](https://github.com/user-attachments/assets/20759890-8b72-4418-a794-50a268baf753)
![Screenshot (143)](https://github.com/user-attachments/assets/5a28d095-ebd2-4eff-a877-9644e0ff3e06)
![Screenshot (142)](https://github.com/user-attachments/assets/0960d495-6efe-4d21-8579-d72783debd54)
![Screenshot (141)](https://github.com/user-attachments/assets/216d46c1-7f2f-4568-ae11-3193d5a8594a)
![Screenshot (140)](https://github.com/user-attachments/assets/61c71a5f-931f-4f91-8e7f-68edfdb655dd)
![Screenshot (139)](https://github.com/user-attachments/assets/5f986e6a-5394-48de-8d86-0e573588029a)
![Screenshot (138)](https://github.com/user-attachments/assets/df498f88-4e5b-4ae6-8e46-c56da1857bc0)
![Screenshot (137)](https://github.com/user-attachments/assets/434be16d-35ec-4662-a720-f6900758569b)
![Screenshot (136)](https://github.com/user-attachments/assets/cca73e2b-d11b-4fb4-b6d0-c1e07d0d2a89)
![Screenshot (135)](https://github.com/user-attachments/assets/14c40bf2-9c7f-4fd2-a16c-3f2f02651ec0)

