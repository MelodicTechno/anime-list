# Redis container management
param(
    [ValidateSet("start", "stop", "restart", "init")]
    [string]$Action = "start"
)

$ErrorActionPreference = "Stop"

function Test-ContainerExists {
    $null -ne (docker ps -a --format "{{.Names}}" | Select-String -Pattern "^redis$")
}

switch ($Action) {
    "init" {
        if (Test-ContainerExists) {
            docker rm -f redis
        }
        docker run -d --name redis -p 6379:6379 redis
        Write-Host "Redis container created and started"
    }
    "start" {
        if (Test-ContainerExists) {
            docker start redis
            Write-Host "Redis container started"
        } else {
            docker run -d --name redis -p 6379:6379 redis
            Write-Host "Redis container created and started"
        }
    }
    "stop" {
        docker stop redis
        Write-Host "Redis container stopped"
    }
    "restart" {
        docker restart redis
        Write-Host "Redis container restarted"
    }
}
