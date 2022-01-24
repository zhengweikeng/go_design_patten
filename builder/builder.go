/**
 * @Author: Seedzheng
 * 建造者模式
 */

package builder

import "errors"

const (
    defaultMaxTotal = 10
    defaultMaxIdle  = 5
)

type ResourcePoolConfig struct {
    name     string
    maxTotal uint
    maxIdle  uint
    minIdle  uint
}

func newResourcePoolConfig(build *resourcePoolBuilder) *ResourcePoolConfig {
    return &ResourcePoolConfig{
        name:     build.name,
        maxTotal: build.maxTotal,
        maxIdle:  build.maxIdle,
        minIdle:  build.minIdle,
    }
}

type resourcePoolBuilder struct {
    name     string // 必填
    maxTotal uint   // 必填
    maxIdle  uint   // 必填
    minIdle  uint   // 选填
}

func NewResourcePoolBuilder() *resourcePoolBuilder {
    return &resourcePoolBuilder{
        maxTotal: defaultMaxTotal,
        maxIdle:  defaultMaxIdle,
    }
}

func (builder *resourcePoolBuilder) Build() (*ResourcePoolConfig, error) {
    if builder.name == "" {
        return nil, errors.New("name不能为空")
    }

    if builder.maxTotal == 0 {
        return nil, errors.New("maxTotal不能0")
    }

    if builder.maxIdle == 0 {
        return nil, errors.New("maxIdle不能0")
    }

    if builder.minIdle > builder.maxIdle {
        return nil, errors.New("minIdle不能大于maxIdle")
    }

    return newResourcePoolConfig(builder), nil
}

func (builder *resourcePoolBuilder) SetName(name string) *resourcePoolBuilder {
    builder.name = name
    return builder
}

func (builder *resourcePoolBuilder) SetMaxTotal(maxTotal uint) *resourcePoolBuilder {
    builder.maxTotal = maxTotal
    return builder
}

func (builder *resourcePoolBuilder) SetMaxIdle(maxIdle uint) *resourcePoolBuilder {
    builder.maxIdle = maxIdle
    return builder
}

func (builder *resourcePoolBuilder) SetMinIdle(minIdle uint) *resourcePoolBuilder {
    builder.minIdle = minIdle
    return builder
}
