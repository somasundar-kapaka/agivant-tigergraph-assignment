# TigerGraph 4.1.3 — Docker Setup

Run TigerGraph 4.1.3 locally using Docker Compose with persistent storage, health checks, and full port exposure.

---

## Prerequisites

| Tool | Minimum Version |
|---|---|
| Docker | 29.4.0 |
| Docker Compose | v5.1.2 |
| RAM | 8 GB available |
| Disk | 20 GB free |



---

## Quick Start

### 1. Create your `.env` file

```bash
export TG_LICENSE=<license> 
```


### 2. Start TigerGraph

```bash
docker compose -f deploy/compose.yaml up -d
```

### 3. Wait for startup 

```bash
docker compose -f deploy/compose.yaml  logs -f 
```

Wait until you see TigerGraph services fully initialized. Then verify:

```bash
sudo docker exec -it tigergraph-4.1.3 bash
gadmin status
```

### 4. Open GraphStudio

```
# Note: check the inbound rules of cloud
http://<instance-ip>:14240
```

Default credentials: `tigergraph / tigergraph`

---

## Port Reference

| Port | Purpose |
|---|---|
| `9000` | REST++ API endpoint |
| `14240` | GraphStudio & Admin Portal (browser UI) |
| `2222` | SSH into the container |
<!-- | `8123` | GSQL client connection |
| `7070` | Nginx reverse proxy | -->

---

## Docker Commands

### Lifecycle

```bash
# Start (detached)
docker compose -f deploy/compose.yaml up -d

# Stop (keeps data)
docker compose -f deploy/compose.yaml stop

# Stop and remove containers (keeps volumes)
docker compose -f deploy/compose.yaml down

# Stop and remove containers + ALL data volumes (destructive)
docker compose -f deploy/compose.yaml down -v
```

### Logs & Monitoring

```bash
# Stream live logs
docker compose -f deploy/compose.yaml logs -f

# Last 100 lines
docker compose -f deploy/compose.yaml logs --tail=100 

# Check container health status
docker inspect --format='{{.State.Health.Status}}' tigergraph-4.1.3
```

### Shell Access

```bash
# Exec into running container
docker exec -it tigergraph-4.1.3 bash

# SSH into container (password: tigergraph)
ssh -p 2222 tigergraph@localhost
```

### GSQL

```bash
# Run exec command
sudo docker exec -it tigergraph-4.1.3 bash
gsql   # Run a GSQL command directly
```

### Data 

```bash
# List named volumes
docker volume ls | grep tg_
```

---

## Persistent Storage

All data survives container restarts via named Docker volumes:

| Volume | Path inside container | Contents |
|---|---|---|
| `tg_data` | `/home/tigergraph/tigergraph/data` | Graph data |
| `tg_app` | `/home/tigergraph/tigergraph/app` | App binaries |

---

## Resource Defaults

| Setting | Value |
|---|---|
| CPU limit | 4 cores |
| CPU reservation | 2 cores |
| Memory limit | 8 GB |
| Memory reservation | 4 GB |

Adjust `deploy.resources` in `deploy/compose.yaml` to match your host machine.

---

## Troubleshooting

**Container exits immediately**
```bash
docker compose -f deploy/compose.yaml logs 
# Check for "license" or "memory" errors in output
```

**GraphStudio not loading**
```bash
# Confirm port is bound
docker port tigergraph-4.1.3
# Wait the full 2 minutes — UI starts after REST++ is ready
```

**GSQL connection refused**
```bash
# Check TigerGraph internal services
docker exec -it tigergraph-4.1.3 gadmin status
```

**Out of memory / OOM killed**
```bash
# Increase memory limit in docker-compose.yaml
# Or free up host RAM and restart
docker compose restart
```

---

## Version Info

| Item | Value |
|---|---|
| TigerGraph version | `4.1.3` |
| Docker image | `tigergraph/tigergraph:4.1.3` |
| Base OS | Ubuntu (amd64) |
| GraphStudio | Bundled |