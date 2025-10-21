# Week 3

Week 3's main objective is:

- API Logging (to files, or using a tech stack ELK)

## Technology Stack

- Golang 1.25:
  - GIN Web Framework (routing, middlewares, validation)
  - GORM + MySQL Driver (modeling)
- OpenSearch (equivalent to Elasticsearch v7.1)
- OpenSearch Dashboard (equivalent to Kibana v7.1)
- Logstash

Ports used:

- `3002` for GIN Web Framework.
- `9200` and `9400` for OpenSearch (Elasticsearch).
- `3306` for MySQL database.
- `5601` for OpenSearch Dashboards (Kibana).

## Installation

### Docker (Recommended)

Stay at root, and run:

```bash
docker compose --profile week3 up
```

This should setup all necessary services. Make sure to fill in the environment
values as requested with an `.env` file at root directory.

### Native

1. Install OpenSearch (by yourself)
2. Install OpenSearch Dashboard (by yourself)
3. Install Logstash (of course, by yourself)
4. Setup all environment variables for your shell or command prompt. An `.env`
   inside the `week3` folder will also work.
5. Run `go run` in `week3`

## Trace

Setting up with Docker Compose:

```yml
# Taken from OpenSearch's Docker example file
opensearch:
  image: opensearchproject/opensearch
  container_name: opensearch
  hostname: opensearch
  environment:
    node.name: opensearch-node1
    discovery.seed_hosts: opensearch-node1
    OPENSEARCH_JAVA_OPTS: -Xmx512m -Xmx512m # Set max memory usages
    bootstrap.memory_lock: true # Don't allow swapfile
    OPENSEARCH_INITIAL_ADMIN_PASSWORD: ${OPENSEARCH_ADMIN_PASSWORD}
    discovery.type: single-node
  ulimits:
    memlock:
      soft: -1
      hard: -1
    nofile:
      soft: 65536 # maximum number of open files for the OpenSearch user, set to at least 65536 on modern systems
      hard: 65536
  volumes:
    - opensearch-data:/usr/share/opensearch/data
  ports:
    - 9200:9200
    - 9600:9600

# Taken from OpenSearch's Docker example file
opensearch-dashboards:
  image: opensearchproject/opensearch-dashboards
  container_name: opensearch-dashboard
  ports:
    - 5601:5601
  expose:
    - "5601"
  environment:
    OPENSEARCH_HOSTS: '["https://opensearch:9200"]'
```

### FluentBit setup

## References

- [FluentBit Manual](https://docs.fluentbit.io/manual/concepts/data-pipeline)
