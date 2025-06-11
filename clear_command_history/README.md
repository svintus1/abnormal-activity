# clear_command_history

Этот скрипт предназначен для удаления файла .bash_history, содержащего историю командной оболочки bash.

MITRE ATT&CK - T1070.003 Clear Command History

## Что делает скрипт

1. Проверяет, что запущен от root.
2. Копирует файл assets/.bash_history в /root/.bash_history.
3. Удаляет файл /root/.bash_history.
4. Логирует основные действия и ошибки.

## Сборка

Команда для сборки - go build -o clear_command_history ./main.go.

## Запуск

Команда для запуска -  ./clear_command_history.

## Требования

1. Golang 1.24.2 и выше.
2. Запуск от имени суперпользователя.