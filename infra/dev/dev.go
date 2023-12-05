// Package dev deploys the hello world app to GCP using Cloud Run and Pulumi
package dev

import (
	cloudrun "github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const IMAGE_URI = "docker.io/nullstring/cicd-pulumi-demo:latest"

func Go() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a new Cloud Run Service
		helloWorldService, err := cloudrun.NewService(ctx, "helloworld-service", &cloudrun.ServiceArgs{
			Location: pulumi.String("us-central1"),
			Template: &cloudrun.ServiceTemplateArgs{
				Spec: &cloudrun.ServiceTemplateSpecArgs{
					Containers: cloudrun.ServiceTemplateSpecContainerArray{
						&cloudrun.ServiceTemplateSpecContainerArgs{
							Image: pulumi.String(IMAGE_URI),
						},
					},
				},
			},
		})

		if err != nil {
			return err
		}

		// Create an IAM member to make the service publicly accessible.
		_, err = cloudrun.NewIamMember(ctx, "helloworld-invoker", &cloudrun.IamMemberArgs{
			Service:  helloWorldService.Name,
			Location: helloWorldService.Location,
			Role:     pulumi.String("roles/run.invoker"),
			Member:   pulumi.String("allUsers"),
		})

		if err != nil {
			return err
		}

		// Exports
		ctx.Export("srvurl", helloWorldService.Statuses.Index(pulumi.Int(0)).Url())

		return nil
	})
}
