locals {
  function_name = reverse(split("/", path.cwd))[0]
  archive_path = "${path.module}/../../../lambda-handler-${function_name}.zip"
}
