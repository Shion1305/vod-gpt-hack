include {
  path = find_in_parent_folders()
}

inputs = {
  aws_access_key = get_env("AWS_ACCESS_KEY_ID")
  aws_secret_key = get_env("AWS_SECRET_ACCESS_KEY")
  aws_session_token = get_env("AWS_SESSION_TOKEN")
}
