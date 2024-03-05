# Dock-ECR

Dock-ECR is a lightweight command-line tool that allows you to interface with Docker and AWS ECR. It simplifies the process of building, pulling, and pushing Docker images to the cloud repository.
Available Commands
```
build - Builds a Docker image and, if specified, pushes it to a cloud repository.
pull - Pulls a Docker image using its resource URI.
push - Pushes a Docker image tagged with a repository name to a cloud repository.
```

## Usage
### `dock-ecr build`

Builds a Docker image and, if specified, pushes it to a cloud repository.

```
Usage:
  dock-ecr build [flags]

Flags:
  -d, --dockerfile string                Path to Dockerfile (default "Dockerfile")
  -h, --help                             help for build
  -i, --image-tags stringArray           Docker tags to be assigned to image
      --push true                        If true, pushes the image to the specified repository
  -r, --repository-name string           Repository of image
  -t, --repository-tags stringToString   Repository resource tags to be assigned (default [])
```
Example of building a local dockerfile, tagging it as `v1.0.0`, and pushing it to the `common/dock-ecr` repository in AWS ECR:
```shell
dock-ecr build --dockerfile Dockerfile \
  --repository-name common/dock-ecr \
  --image-tags v1.0.0 \
  --repository-tags deployer=gregfurman \
  --push
```

### `dock-ecr pull`

Pulls a Docker image using its resource URI.

```
Usage:
  dock-ecr pull [flags]

Flags:
  -h, --help                help for pull
  -i, --image-name string   URI of image
```

Example pulling an image from the `test/repo` repository in with the tag `1.0.0`:
```shell
dock-ecr --image-name 012345678910.dkr.ecr.<region>.amazonaws.com/test/repo:1.0.0

```
### `dock-ecr push`

```
Usage:
  dock-ecr push [flags]

Flags:
  -h, --help                             help for push
  -i, --image-tags stringArray           Docker tags to be assigned to image
  -r, --repository-name string           Repository of image
  -t, --repository-tags stringToString   Repository resource tags to be assigned (default [])
```

Example of pushing a local Docker image, tagged as `v1.0.0` and `staging`, to the `common/dock-ecr` repository in AWS ECR:
```shell
dock-ecr push \
  --repository-name common/dock-ecr \
  --image-tags v1.0.0,staging \
  --repository-tags deployer=gregfurman
```

## Installation

To install Dock-ECR, you need to have Golang installed on your system. You can install Dock-ECR by cloning this repository and running the `make build` command.

## License

This project is licensed under the MIT License - see the LICENSE file for details.