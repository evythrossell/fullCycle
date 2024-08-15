
resource "local_file" "exemplo" {
    filename = "exemplo.txt"
    content = var.content
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