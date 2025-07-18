---
trigger: manual
description: create scaffolding for new proyects
---

## Scaffolding template for backend with Go

.
├─ cmd/
│  └─ api/
│     └─ main.go
└─ internal/
   ├─ app/
   │  └─ config/
   │     ├─ config.go
   │     ├─ environment_config.go
   │     ├─ environment_config_dev.go
   │     ├─ environment_config_beta.go
   │     └─ environment_config_prod.go
   ├─ core/
   │  ├─ domain/
   │  │  └─ <Entity>.go
   │  ├─ errors/
   │  │  └─ errors.go
   │  └─ usecase/
   │     └─ <entity>_usecase.go
   └─ infrastructure/
      ├─ adapters/
      │  └─ repositories/mysql/
      │     ├─ mysql_client.go
      │     └─ <entity>_mysql_repository.go
      └─ entrypoints/
         ├─ handler/
         │  └─ <entity>_handler.go
         └─ router/
            ├─ handlers_container.go
            ├─ url_mappings.go
            └─ web_application.go