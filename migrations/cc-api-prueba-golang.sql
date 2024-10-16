--liquibase formatted sql

--changeset dfrd:1 labels:initial context:cc-api-prueba-golang:v1.0
--comment: Primer commit para permitir hacer rollback
select 1 ;
--rollback select 1;