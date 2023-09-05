output "instance_1_ip_addr" {
  value = aws_instance.instance_trkd.public_ip
}

output "instance_2_ip_addr" {
  value = aws_instance.instance_trkd.private_ip
}
