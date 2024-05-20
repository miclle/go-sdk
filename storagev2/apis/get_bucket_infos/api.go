// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 获取用户所有存储空间信息
package get_bucket_infos

import (
	"encoding/json"
	credentials "github.com/qiniu/go-sdk/v7/storagev2/credentials"
	errors "github.com/qiniu/go-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Region      string                          // 区域 ID
	Statistics  bool                            // 是否返回存储空间的实时统计信息
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	AllBucketInfosV2 AllBucketInfosV2 // 所有存储空间信息
}

// 防盗链 Referer 白名单
type AllowedReferer = []string

// 防盗链 Referer 黑名单
type BlockedReferer = []string
type BucketInfo struct {
	Source         string            // 镜像回源地址，可以有多个，以 `;` 分隔
	Host           string            // 请求镜像地址时携带的 `Host` 头部
	Expires        int64             // 镜像回源地址过期时长
	Protected      int64             // 是否开启了原图保护
	Private        int64             // 是否是私有空间
	NoIndexPage    int64             // 是否禁用 index.html（或 index.htm） 页面
	MaxAge         int64             // 客户端缓存过期时长
	Separator      string            // 图片样式分隔符，可能返回多个
	Styles         map[string]string // 图片样式，键表示图片样式命令名字，值表示图片样式命令内容
	AntiLeechMode  int64             // 防盗链模式，1：表示设置 Referer 白名单; 2：表示设置 Referer 黑名单
	TokenAntiLeech int64             // Token 防盗链模式, 0：表示关闭，1：表示打开
	ReferWl        AllowedReferer    // 防盗链 Referer 白名单列表
	ReferBl        BlockedReferer    // 防盗链 Referer 黑名单列表
	SourceEnabled  bool              // 在源站支持的情况下是否开启源站的 Referer 防盗链
	NoReferer      bool              // 0：表示不允许空 Refer 访问; 1：表示允许空 Refer 访问
	MacKey         string            // 第一个 MacKey，Index 为 1，用于防盗链 Token 的生成
	MacKey2        string            // 第二个 MacKey，Index 为 2，用于防盗链 Token 的生成
	Zone           string            // 存储区域，兼容保留
	Region         string            // 存储区域
	Remark         string            // 空间备注信息
	CreatedAt      string            // 空间创建时间
}

// 存储空间信息
type BucketInfoV2 = BucketInfo
type jsonBucketInfo struct {
	Source         string            `json:"source"`                // 镜像回源地址，可以有多个，以 `;` 分隔
	Host           string            `json:"host"`                  // 请求镜像地址时携带的 `Host` 头部
	Expires        int64             `json:"expires"`               // 镜像回源地址过期时长
	Protected      int64             `json:"protected"`             // 是否开启了原图保护
	Private        int64             `json:"private"`               // 是否是私有空间
	NoIndexPage    int64             `json:"no_index_page"`         // 是否禁用 index.html（或 index.htm） 页面
	MaxAge         int64             `json:"max_age"`               // 客户端缓存过期时长
	Separator      string            `json:"separator"`             // 图片样式分隔符，可能返回多个
	Styles         map[string]string `json:"styles"`                // 图片样式，键表示图片样式命令名字，值表示图片样式命令内容
	AntiLeechMode  int64             `json:"anti_leech_mode"`       // 防盗链模式，1：表示设置 Referer 白名单; 2：表示设置 Referer 黑名单
	TokenAntiLeech int64             `json:"token_anti_leech_mode"` // Token 防盗链模式, 0：表示关闭，1：表示打开
	ReferWl        AllowedReferer    `json:"refer_wl"`              // 防盗链 Referer 白名单列表
	ReferBl        BlockedReferer    `json:"refer_bl"`              // 防盗链 Referer 黑名单列表
	SourceEnabled  bool              `json:"source_enabled"`        // 在源站支持的情况下是否开启源站的 Referer 防盗链
	NoReferer      bool              `json:"allow_empty_referer"`   // 0：表示不允许空 Refer 访问; 1：表示允许空 Refer 访问
	MacKey         string            `json:"mac_key"`               // 第一个 MacKey，Index 为 1，用于防盗链 Token 的生成
	MacKey2        string            `json:"mac_key2"`              // 第二个 MacKey，Index 为 2，用于防盗链 Token 的生成
	Zone           string            `json:"zone"`                  // 存储区域，兼容保留
	Region         string            `json:"region"`                // 存储区域
	Remark         string            `json:"remark"`                // 空间备注信息
	CreatedAt      string            `json:"ctime"`                 // 空间创建时间
}

func (j *BucketInfo) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonBucketInfo{Source: j.Source, Host: j.Host, Expires: j.Expires, Protected: j.Protected, Private: j.Private, NoIndexPage: j.NoIndexPage, MaxAge: j.MaxAge, Separator: j.Separator, Styles: j.Styles, AntiLeechMode: j.AntiLeechMode, TokenAntiLeech: j.TokenAntiLeech, ReferWl: j.ReferWl, ReferBl: j.ReferBl, SourceEnabled: j.SourceEnabled, NoReferer: j.NoReferer, MacKey: j.MacKey, MacKey2: j.MacKey2, Zone: j.Zone, Region: j.Region, Remark: j.Remark, CreatedAt: j.CreatedAt})
}
func (j *BucketInfo) UnmarshalJSON(data []byte) error {
	var nj jsonBucketInfo
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Source = nj.Source
	j.Host = nj.Host
	j.Expires = nj.Expires
	j.Protected = nj.Protected
	j.Private = nj.Private
	j.NoIndexPage = nj.NoIndexPage
	j.MaxAge = nj.MaxAge
	j.Separator = nj.Separator
	j.Styles = nj.Styles
	j.AntiLeechMode = nj.AntiLeechMode
	j.TokenAntiLeech = nj.TokenAntiLeech
	j.ReferWl = nj.ReferWl
	j.ReferBl = nj.ReferBl
	j.SourceEnabled = nj.SourceEnabled
	j.NoReferer = nj.NoReferer
	j.MacKey = nj.MacKey
	j.MacKey2 = nj.MacKey2
	j.Zone = nj.Zone
	j.Region = nj.Region
	j.Remark = nj.Remark
	j.CreatedAt = nj.CreatedAt
	return nil
}
func (j *BucketInfo) validate() error {
	if j.Source == "" {
		return errors.MissingRequiredFieldError{Name: "Source"}
	}
	if j.Host == "" {
		return errors.MissingRequiredFieldError{Name: "Host"}
	}
	if j.Expires == 0 {
		return errors.MissingRequiredFieldError{Name: "Expires"}
	}
	if j.Protected == 0 {
		return errors.MissingRequiredFieldError{Name: "Protected"}
	}
	if j.Private == 0 {
		return errors.MissingRequiredFieldError{Name: "Private"}
	}
	if j.NoIndexPage == 0 {
		return errors.MissingRequiredFieldError{Name: "NoIndexPage"}
	}
	if j.MaxAge == 0 {
		return errors.MissingRequiredFieldError{Name: "MaxAge"}
	}
	if j.Separator == "" {
		return errors.MissingRequiredFieldError{Name: "Separator"}
	}
	if j.Styles == nil {
		return errors.MissingRequiredFieldError{Name: "Styles"}
	}
	if j.AntiLeechMode == 0 {
		return errors.MissingRequiredFieldError{Name: "AntiLeechMode"}
	}
	if j.TokenAntiLeech == 0 {
		return errors.MissingRequiredFieldError{Name: "TokenAntiLeech"}
	}
	if len(j.ReferWl) == 0 {
		return errors.MissingRequiredFieldError{Name: "ReferWl"}
	}
	if len(j.ReferBl) == 0 {
		return errors.MissingRequiredFieldError{Name: "ReferBl"}
	}
	if j.MacKey == "" {
		return errors.MissingRequiredFieldError{Name: "MacKey"}
	}
	if j.MacKey2 == "" {
		return errors.MissingRequiredFieldError{Name: "MacKey2"}
	}
	if j.Zone == "" {
		return errors.MissingRequiredFieldError{Name: "Zone"}
	}
	if j.Region == "" {
		return errors.MissingRequiredFieldError{Name: "Region"}
	}
	if j.Remark == "" {
		return errors.MissingRequiredFieldError{Name: "Remark"}
	}
	if j.CreatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "CreatedAt"}
	}
	return nil
}

// 所有存储空间信息
type BucketInfosV2 struct {
	Name       string // 存储空间名称
	BucketInfo BucketInfoV2
}
type jsonBucketInfosV2 struct {
	Name       string       `json:"name"` // 存储空间名称
	BucketInfo BucketInfoV2 `json:"info"`
}

func (j *BucketInfosV2) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonBucketInfosV2{Name: j.Name, BucketInfo: j.BucketInfo})
}
func (j *BucketInfosV2) UnmarshalJSON(data []byte) error {
	var nj jsonBucketInfosV2
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Name = nj.Name
	j.BucketInfo = nj.BucketInfo
	return nil
}
func (j *BucketInfosV2) validate() error {
	if j.Name == "" {
		return errors.MissingRequiredFieldError{Name: "Name"}
	}
	if err := j.BucketInfo.validate(); err != nil {
		return err
	}
	return nil
}

// 所有存储空间信息
type AllBucketInfosV2 []BucketInfosV2

func (j *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.AllBucketInfosV2)
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var array AllBucketInfosV2
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}
	j.AllBucketInfosV2 = array
	return nil
}
