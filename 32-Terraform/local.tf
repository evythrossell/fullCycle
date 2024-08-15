
resource "local_file" "exemplo" {
    filename = "exemplo.txt"
    content = var.content
}

variable "content" {
    type = string
}