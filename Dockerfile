FROM golang:1.25.5-alpine AS builder
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/dragonctl ./cmd/dragonctl
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/k8s-kubeadm-token-server ./cmd/k8s-kubeadm-token-server
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/scheduler ./cmd/scheduler

FROM scratch AS dragonctl
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/dragonctl /dragonctl
ENTRYPOINT ["/dragonctl"]

FROM scratch AS k8s-kubeadm-token-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/k8s-kubeadm-token-server /k8s-kubeadm-token-server
ENTRYPOINT ["/k8s-kubeadm-token-server"]

FROM scratch AS scheduler
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/scheduler /scheduler
ENTRYPOINT ["/scheduler"]
