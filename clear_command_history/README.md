Этот скрипт предназначен для удаления файла .bash_history, содержащего историю командной оболочки bash.
ATT&CK: T1070.003 Clear Command History
Команда для сборки - go build -o clear_command_history ./main.go
Команда для запуска -  ./clear_command_history
Требования: Golang 1.24.2 и выше, запуск от имени суперпользователя