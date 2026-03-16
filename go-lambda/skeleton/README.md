# ${{ values.name }}

${{ values.description }}

## Desarrollo

```bash
just test    # tests con race detector
just deploy  # build + package + terraform apply
just destroy # terraform destroy
```

## Stack

Go · AWS Lambda (`provided.al2023` · `${{ values.architecture }}`) · API Gateway HTTP v2 · Terraform
