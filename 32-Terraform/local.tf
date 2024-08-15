
resource "local_file" "exemplo" {
    filename = "exemplo.txt"
    content = var.content
}

data "local_file" "content-sample" {
    filename = "exemplo.txt"
}

output "data-source-result" {
    value = data.local_file.content-sample.content_base64
}

variable "content" {
    type = string
}

output "file-id" {
    value = resource.local_file.exemplo.id
}

output "content" {
    value = var.content
}