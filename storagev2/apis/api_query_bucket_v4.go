// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	uplog "github.com/qiniu/go-sdk/v7/internal/uplog"
	querybucketv4 "github.com/qiniu/go-sdk/v7/storagev2/apis/query_bucket_v4"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
	httpclient "github.com/qiniu/go-sdk/v7/storagev2/http_client"
	region "github.com/qiniu/go-sdk/v7/storagev2/region"
	"net/url"
	"strings"
)

type innerQueryBucketV4Request querybucketv4.Request

func (query *innerQueryBucketV4Request) getBucketName(ctx context.Context) (string, error) {
	return query.Bucket, nil
}
func (query *innerQueryBucketV4Request) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Bucket != "" {
		allQuery.Set("bucket", query.Bucket)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Bucket"}
	}
	if query.AccessKey != "" {
		allQuery.Set("ak", query.AccessKey)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "AccessKey"}
	}
	return allQuery, nil
}

type QueryBucketV4Request = querybucketv4.Request
type QueryBucketV4Response = querybucketv4.Response

// 查询存储空间服务域名
func (storage *Storage) QueryBucketV4(ctx context.Context, request *QueryBucketV4Request, options *Options) (*QueryBucketV4Response, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerQueryBucketV4Request)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	pathSegments := make([]string, 0, 2)
	pathSegments = append(pathSegments, "v4", "query")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	bucketName := options.OverwrittenBucketName
	if bucketName == "" {
		var err error
		if bucketName, err = innerRequest.getBucketName(ctx); err != nil {
			return nil, err
		}
	}
	uplogInterceptor, err := uplog.NewRequestUplog("queryBucketV4", bucketName, "", nil)
	if err != nil {
		return nil, err
	}
	req := httpclient.Request{Method: "GET", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, BufferResponse: true, OnRequestProgress: options.OnRequestProgress}
	if options.OverwrittenEndpoints == nil && options.OverwrittenRegion == nil && storage.client.GetRegions() == nil {
		bucketHosts := httpclient.DefaultBucketHosts()
		if options.OverwrittenBucketHosts != nil {
			req.Endpoints = options.OverwrittenBucketHosts
		} else {
			req.Endpoints = bucketHosts
		}
	}
	ctx = httpclient.WithoutSignature(ctx)
	var respBody QueryBucketV4Response
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}