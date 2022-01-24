/**
 * @Author: Seedzheng
 * 责任链模式
 */

package responsibilitychain

import "fmt"

type Job interface {
	Check() bool
}

type QuantityJob struct {
	Quantity int
}

func (q QuantityJob) Check() bool {
	if q.Quantity < 10 {
		return false
	}
	fmt.Println("检查数量")
	return true
}

type QualityJob struct {
	Quality int
}

func (q QualityJob) Check() bool {
	if q.Quality < 95 {
		return false
	}
	fmt.Println("检查质量")
	return true
}

type PackageJob struct {
	IsBroke bool
}

func (p PackageJob) Check() bool {
	if p.IsBroke {
		return false
	}
	fmt.Println("检查包装")
	return true
}

type Pipeline struct {
	jobs []Job
}

func (p *Pipeline) AddJob(job Job) {
	p.jobs = append(p.jobs, job)
}

func (p *Pipeline) DoJobs() bool {
	success := false
	for i, job := range p.jobs {
		if !job.Check() {
			break
		}

		if i == len(p.jobs)-1 {
			success = true
		}
	}

	return success
}
