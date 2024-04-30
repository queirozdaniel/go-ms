# API para autenticação

API para guardar e servir recursos para Autenticação e autorização.

1. API REST para Autenticação
2. API gRPC para Autorização


## Modelagem

Todas as tabelas e mapeamentos serão baseadas nas entidades.

Entidade:
- Colunas (`created_date`, `updated_date`, `is_active`, `is_deleted`)

1. Empresa "Tenant": (`id`, `name`, `subname`, `cnpj`)
2. Usuario "User": (`id`, `access_code`, `password`, `is_checked`, `access_mode`, `id_person`, `id_tenant`)
3. Pessoa "Person": (`id`, `email`, `cellphone`, `address`, `id_tenant`)
4. Permissão "Role": (`id`, `key`, `description`)
5. Usuario Permissao "User Permissio": (`id`, `user_id`, `role_id`)
