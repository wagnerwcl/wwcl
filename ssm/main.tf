resource "aws_ssm_parameter" "this" {
  name  = "foo-01"
  type  = "String"
  value = "barr"
}   
