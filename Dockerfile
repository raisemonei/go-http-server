FROM bitnami/minideb:stretch
RUN echo "deb http://archive.debian.org/debian stretch main" > /etc/apt/sources.list

WORKDIR /app

COPY apis /app/apis

RUN install_packages ca-certificates wget curl && groupadd -r appgroup && useradd -r -g appgroup appuser

# Change ownership of the application files to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER appuser

# Ensure the binary has execution permissions
RUN chmod +x /app/apis

CMD ["/app/apis"]