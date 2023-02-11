resource "aws_iam_user" "dmanga_admin" {
  name = "dmanga_admin"
}

resource "aws_iam_access_key" "dmanga_admin" {
  user = aws_iam_user.dmanga_admin.name
}

data "aws_iam_policy_document" "backend" {
  statement {
    actions = [
      "s3:PutObject"
    ]
    resources = [
      aws_s3_bucket.dmanga_resource.arn
    ]
  }
}

resource "aws_iam_user_policy" "create_presign_client" {
  name   = "dmanga-create-presign-client"
  user   = aws_iam_user.dmanga_admin.name
  policy = data.aws_iam_policy_document.backend.json
}
