/**
 * @Author: Seedzheng
 */

package responsibilitychain

import (
	"fmt"
	"testing"
)

func TestPipeline_DoJobs(t *testing.T) {
	type fields struct {
		jobs []Job
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "chain of responsibility, success",
			fields: fields{jobs: []Job{
				QuantityJob{Quantity: 20},
				QualityJob{Quality: 99},
				PackageJob{IsBroke: false},
			}},
		},
		{
			name: "chain of responsibility, fail",
			fields: fields{jobs: []Job{
				QuantityJob{Quantity: 20},
				QualityJob{Quality: 90},
				PackageJob{IsBroke: false},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pipeline{
				jobs: tt.fields.jobs,
			}
			if p.DoJobs() {
				fmt.Println("流水线成功")
			} else {
				fmt.Println("流水线失败")
			}
		})
	}
}
