remote_state {
  backend = "s3"
  config = {
    bucket  = "progate-vod-gpt-terraform"
    key     = "${path_relative_to_include()}.tfstate"
    region  = "us-west-2"
    encrypt = true
  }
}
