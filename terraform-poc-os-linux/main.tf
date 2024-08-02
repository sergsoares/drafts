terraform {
  required_providers {
    system = {
      source = "neuspaces/system"
      version = "0.5.0"
    }
    remote = {
      source = "tenstad/remote"
      version = "0.1.3"
    }
  }
}

locals {
  host = "mytest.ourlab.cc"
  user = "root"
  port = 22
  ssh_key_file = "./key"
}

provider "system" {
  ssh {
    host        = local.host
    port        = local.port
    user        = local.user
    private_key = file(ssh_key_file)
  }
}

resource "system_packages_apt" "packages" {
  package {
    name = "curl"
  }

  package {
    name = "docker.io"
  }

  package {
    name = "caddy"
  }
}


resource "terraform_data" "vscode" {
  depends_on = [ system_packages_apt.packages ]
  connection {
    type        = "ssh"
    user        = local.user
    private_key = file(local.ssh_key_file)
    host        = local.host
  }
  provisioner "remote-exec" {
    inline = [
      "curl -fsSL https://code-server.dev/install.sh | sh"
    ]
  }
}

resource "system_service_systemd" "vscode" {
  depends_on = [ terraform_data.vscode ]
  name    = "code-server@${local.user}"
  enabled = true
  status = "started"
}

resource "remote_file" "caddyfile" {
  depends_on = [ system_packages_apt.caddy ]
  conn {
    host        = local.host
    port        = 22
    user        = local.user
    private_key = file(local.ssh_key_file)
  }

  path        = "/etc/caddy/Caddyfile"
  content = file("${path.module}/caddy/Caddyfile")
}

resource "system_service_systemd" "caddy" {
  depends_on = [ system_packages_apt.caddy, remote_file.caddyfile]
  name    = "caddy"
  enabled = true
  status = "started"
  restart_on  = ["v1"]
}