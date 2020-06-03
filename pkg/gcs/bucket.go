package gcs

import (
	"context"
	"fmt"
	"io"
	"time"

	// "github.com/spf13/viper"
	"cloud.google.com/go/storage"
)

// getBucketMetadata gets the bucket metadata.
func getBucketMetadata(w io.Writer, bucketName string) (*storage.BucketAttrs, error) {
	// bucketName := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	attrs, err := client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Bucket(%q).Attrs: %v", bucketName, err)
	}
	fmt.Fprintf(w, "BucketName: %v\n", attrs.Name)
	fmt.Fprintf(w, "Location: %v\n", attrs.Location)
	fmt.Fprintf(w, "LocationType: %v\n", attrs.LocationType)
	fmt.Fprintf(w, "StorageClass: %v\n", attrs.StorageClass)
	fmt.Fprintf(w, "TimeCreated: %v\n", attrs.Created)
	fmt.Fprintf(w, "Metageneration: %v\n", attrs.MetaGeneration)
	fmt.Fprintf(w, "PredefinedACL: %v\n", attrs.PredefinedACL)
	if attrs.Encryption != nil {
		fmt.Fprintf(w, "DefaultKmsKeyName: %v\n", attrs.Encryption.DefaultKMSKeyName)
	}
	if attrs.Website != nil {
		fmt.Fprintf(w, "IndexPage: %v\n", attrs.Website.MainPageSuffix)
		fmt.Fprintf(w, "NotFoundPage: %v\n", attrs.Website.NotFoundPage)
	}
	fmt.Fprintf(w, "DefaultEventBasedHold: %v\n", attrs.DefaultEventBasedHold)
	if attrs.RetentionPolicy != nil {
		fmt.Fprintf(w, "RetentionEffectiveTime: %v\n", attrs.RetentionPolicy.EffectiveTime)
		fmt.Fprintf(w, "RetentionPeriod: %v\n", attrs.RetentionPolicy.RetentionPeriod)
		fmt.Fprintf(w, "RetentionPolicyIsLocked: %v\n", attrs.RetentionPolicy.IsLocked)
	}
	fmt.Fprintf(w, "RequesterPays: %v\n", attrs.RequesterPays)
	fmt.Fprintf(w, "VersioningEnabled: %v\n", attrs.VersioningEnabled)
	if attrs.Logging != nil {
		fmt.Fprintf(w, "LogBucket: %v\n", attrs.Logging.LogBucket)
		fmt.Fprintf(w, "LogObjectPrefix: %v\n", attrs.Logging.LogObjectPrefix)
	}
	if attrs.Labels != nil {
		fmt.Fprintf(w, "\n\n\nLabels:")
		for key, value := range attrs.Labels {
			fmt.Fprintf(w, "\t%v = %v\n", key, value)
		}
	}
	return attrs, nil
}
