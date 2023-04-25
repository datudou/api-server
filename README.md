# Getting Started
These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites
Before you begin, ensure that you have the following software installed on your system:

* Go (version 1.19 or higher)
* Git

### Installation
To set up the project on your local machine, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/datudou/api-server.git
```

2. Change to the project directory:

```bash
cd api-server
```
3. Change the config-dev.yaml 

```bash
vim config-dev.yaml
```
   Modify the database config in the config-dev.yaml

4. Build the project::

```bash
go build -o server ./cmd/server
```

5. Run the application:

```bash
./server
```
Now, the server should be running on localhost:8080.
