# single_cases

Этот проект содержит набор скриптов, реализующих отдельные техники из базы знаний MITRE ATT&CK.
К каждому скрипту прилагается README.md файл с указанием идентификатора техники и описанием скрпта.

## Список скриптов

1. clear_command_history **(Собирается и запускается вручную)**
2. delete_files
3. hidden_file_and_dir
4. http_get_request
5. log_clear
6. masquerade_task
7. path_interception
8. preload_injection
9. process_extension_anomalies
10. pw_search **(Собирается и запускается вручную)**
11. unusual_process_path

## Сборка

Команда для сборки образов:
docker buildx bake

## Запуск

Команда для запуска скриптов (имя скрипта соответсввует названию директории):
docker compose up <имя скрипта>

## Требования

1. Golang 1.20 и выше.
2. Docker Compose