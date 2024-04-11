resource "aws_ssm_parameter" "this" {
  name  = "foo"
  type  = "String"
  value = "barr"
}