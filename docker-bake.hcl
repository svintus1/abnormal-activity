group "default" {
  targets = [
    "clear_command_history",
    "delete_files",
    "hidden_file_and_dir",
    "http_get_request",
    "log_clear",
    "path_interception",
    "preload_injection",
    "process_extension_anomalies",
    "pw_search",
    "masquerade_task",
    "unusual_process_path"
  ]
}

target "clear_command_history" {
  context = "./clear_command_history"
  dockerfile = "Dockerfile.clear_command_history"
  tags = ["clear_command_history:latest"]
}

target "delete_files" {
  context = "./delete_files"
  dockerfile = "Dockerfile.delete_files"
  tags = ["delete_files:latest"]
}

target "hidden_file_and_dir" {
  context = "./hidden_file_and_dir"
  dockerfile = "Dockerfile.hidden_file_and_dir"
  tags = ["hidden_file_and_dir:latest"]
}

target "http_get_request" {
  context = "./http_get_request"
  dockerfile = "Dockerfile.http_get_request"
  tags = ["http_get_request:latest"]
}

target "log_clear" {
  context = "./log_clear"
  dockerfile = "Dockerfile.log_clear"
  tags = ["log_clear:latest"]
}

target "path_interception" {
  context = "./path_interception"
  dockerfile = "Dockerfile.path_interception"
  tags = ["path_interception:latest"]
}

target "preload_injection" {
  context = "./preload_injection"
  dockerfile = "Dockerfile.preload_injection"
  tags = ["preload_injection:latest"]
}

target "process_extension_anomalies" {
  context = "./process_extension_anomalies"
  dockerfile = "Dockerfile.process_extension_anomalies"
  tags = ["process_extension_anomalies:latest"]
}

target "pw_search" {
  context = "./pw_search"
  dockerfile = "Dockerfile.pw_search"
  tags = ["pw_search:latest"]
}

target "masquerade_task" {
  context = "./masquerade_task"
  dockerfile = "Dockerfile.masquerade_task"
  tags = ["masquerade_task:latest"]
}

target "unusual_process_path" {
  context = "./unusual_process_path"
  dockerfile = "Dockerfile.unusual_process_path"
  tags = ["unusual_process_path:latest"]
}
