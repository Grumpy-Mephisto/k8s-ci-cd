<div align="center">
    <h1><code>💻</code> Operating Systems Project</h1>
    <h3>📚 A project for the Operating Systems course at the King Mongkut's University of Technology Ladkrabang.</h3>
</div>

## `⭐` Introduction

โปรเจคนี้เป็นส่วนหนึ่งของวิชา **Operating System** ที่ให้สร้าง Container ของ Docker ขึ้นมา และอธิบายการทำงานของ Docker และ Container ต่างๆ ที่เราสร้างขึ้นมา

## `📝` Requirements

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Kubernetes](https://kubernetes.io/)

## `📚` How to use

1. โคลน Repository นี้ลงมา

```bash
git clone
```

2. สั่งรัน Docker

   2.1 ใช้คำสั่ง `docker compose` ในการรัน

   ```bash
   cd web-server
   docker-compose up -d # docker compose up -d
   ```

   2.2 ใช้คำสั่ง `docker` ในการรัน

   ```bash
   cd web-server
   docker build -t web-server .
   docker run -d -p 3000:3000 web-server
   ```

3. เปิดเว็บบราวเซอร์และไปที่ `http://localhost:3000`

4. เปิดใช้งาน **[ArgoCD](https://argo-cd.readthedocs.io/en/stable/)**

   4.1 สร้าง Namespace ของ ArgoCD

   ```bash
    kubectl create namespace argocd
   ```

   4.2 ติดตั้ง ArgoCD โดยใช้คำสั่ง

   ```bash
    kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
   ```

   4.3 ตรวจสอบว่า ArgoCD ทำงานอยู่หรือไม่

   ```bash
    kubectl get pods -n argocd
   ```

   4.4 เข้าถึงเซิฟเวอร์ ArgoCD โดยใช้คำสั่ง

   ```bash
    kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
    kubectl port-forward svc/argocd-server -n argocd 8080:443
   ```

   4.5 เข้าสู่ระบบโดยใช้ Username และ Password ที่ได้รับจากคำสั่ง

   ```bash
    kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
   ```

## `📝` Description

### `🐳` Docker

ใช้ Docker ในการสร้าง Image ของ Web Server โดย Web Server เราใช้ **Go** ในการทำ Web Server และใช้ **Fiber** เป็น Web Framework ในการทำ Web Server

### `🦑` Docker Compose

ใช้ Docker Compose ในการสร้าง Container ของ Web Server (สำหรับการทดสอบ)

### `🎬` GitHub Actions

ใช้ GitHub Actions ในการ CI ตัว Web Server โดยใช้ **Docker** ในการทำ CI ของ Web Server และทำการ Push Image ของ Web Server ขึ้นไปบน **Docker Hub**

### `☸️` Kubernetes

ใช้ Kubernetes ในการสร้าง Cluster ของ Web Server โดยใช้ **Minikube** ในการสร้าง Cluster ของ Web Server

### `🐙` ArgoCD

ใช้ ArgoCD ในการ Deploy และทดสอบ Web Server โดยใช้ **GitOps** ในการ Deploy และทดสอบ Web Server
