package k8s

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"

	batch "k8s.io/api/batch/v1"
	"k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
)

type JobConfig struct {
	Name                  string
	Namespace             string
	ServiceAccountName    string
	PodAnnotation         map[string]string
	ActiveDeadlineSeconds int64
	BackoffLimit          int32
}

type Job interface {
	ExecuteJobWithCommand(ctx context.Context, config *JobConfig, imageName string, args map[string]string, commands []string) (executionName string, err error)
	ExecuteJob(ctx context.Context, config *JobConfig, imageName string, args map[string]string) (executionName string, err error)
	JobExecutionStatus(ctx context.Context, config *JobConfig, executionName string) (JobStatus, error)
	WaitForReadyJob(ctx context.Context, config *JobConfig, executionName string, waitTime time.Duration) error
	WaitForReadyPod(ctx context.Context, config *JobConfig, executionName string, waitTime time.Duration) (*v1.Pod, error)
	GetPodLogs(ctx context.Context, pod *v1.Pod) (io.ReadCloser, error)
}

type job struct {
	clientSet        kubernetes.Interface
	maxWaitPoolCount int
}

type JobStatus int

const (
	Failed JobStatus = iota
	Succeed
	FetchError
	NotFound
)

func NewJob(kubeConfigPath string, context string) (Job, error) {
	clientSet, err := NewClientSet(kubeConfigPath, context)

	if err != nil {
		return nil, err
	}
	newClient := &job{
		maxWaitPoolCount: 5,
		clientSet:        clientSet,
	}

	return newClient, nil
}

func watcherError(resource string, listOptions meta.ListOptions) error {
	return fmt.Errorf("watch error when waiting for %s with list option %v", resource, listOptions)
}

func timeoutError(resource string, listOptions meta.ListOptions) error {
	return fmt.Errorf("timeout error when waiting for %s with list option %v", resource, listOptions)
}

func (client *job) ExecuteJob(ctx context.Context, config *JobConfig, imageName string, envMap map[string]string) (string, error) {
	return client.ExecuteJobWithCommand(ctx, config, imageName, envMap, []string{})
}

func (client *job) ExecuteJobWithCommand(ctx context.Context, config *JobConfig, imageName string, envMap map[string]string, command []string) (string, error) {
	toEnvVar := func(envMap map[string]string) []v1.EnvVar {
		var envVars []v1.EnvVar
		for k, v := range envMap {
			envVar := v1.EnvVar{
				Name:  k,
				Value: v,
			}
			envVars = append(envVars, envVar)
		}
		return envVars
	}

	executionName := config.Name + "-" + uuid.New().String()
	label := map[string]string{
		"job": executionName,
	}

	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(config.Namespace)

	container := v1.Container{
		Name:  executionName,
		Image: imageName,
		Env:   toEnvVar(envMap),
	}

	if len(command) != 0 {
		container.Command = command
	}

	podSpec := v1.PodSpec{
		Containers:         []v1.Container{container},
		RestartPolicy:      v1.RestartPolicyNever,
		ServiceAccountName: config.ServiceAccountName,
	}

	objectMeta := meta.ObjectMeta{
		Name:        executionName,
		Labels:      label,
		Annotations: config.PodAnnotation,
	}

	template := v1.PodTemplateSpec{
		ObjectMeta: objectMeta,
		Spec:       podSpec,
	}

	jobSpec := batch.JobSpec{
		Template:              template,
		ActiveDeadlineSeconds: &config.ActiveDeadlineSeconds,
		BackoffLimit:          &config.BackoffLimit,
	}

	jobToRun := &batch.Job{
		TypeMeta: meta.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		ObjectMeta: objectMeta,
		Spec:       jobSpec,
	}

	opts := meta.CreateOptions{}

	_, err := kubernetesJobs.Create(ctx, jobToRun, opts)
	if err != nil {
		return "", err
	}
	return executionName, nil
}

func (client *job) WaitForReadyJob(ctx context.Context, config *JobConfig, executionName string, waitTime time.Duration) error {
	batchV1 := client.clientSet.BatchV1()
	jobs := batchV1.Jobs(config.Namespace)
	listOptions := meta.ListOptions{
		TypeMeta: meta.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		LabelSelector: fmt.Sprintf("job=%s", executionName),
	}

	var err error
	for i := 0; i < client.maxWaitPoolCount; i += 1 {
		watchJob, watchErr := jobs.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var job *batch.Job
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("job", listOptions)
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				job = event.Object.(*batch.Job)
				if job.Status.Active >= 1 || job.Status.Succeeded >= 1 || job.Status.Failed >= 1 {
					watchJob.Stop()
					return nil
				}
			case <-timeoutChan:
				err = timeoutError("job", listOptions)
				break
			}

			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	return err
}

func (client *job) WaitForReadyPod(ctx context.Context, config *JobConfig, executionName string, waitTime time.Duration) (*v1.Pod, error) {
	coreV1 := client.clientSet.CoreV1()
	kubernetesPods := coreV1.Pods(config.Namespace)
	listOptions := meta.ListOptions{
		LabelSelector: fmt.Sprintf("job=%s", executionName),
	}

	var err error
	for i := 0; i < client.maxWaitPoolCount; i += 1 {
		watchJob, watchErr := kubernetesPods.Watch(ctx, listOptions)
		if watchErr != nil {
			err = watchErr
			continue
		}

		timeoutChan := time.After(waitTime)
		resultChan := watchJob.ResultChan()

		var pod *v1.Pod
		for {
			select {
			case event := <-resultChan:
				if event.Type == watch.Error {
					err = watcherError("pod", listOptions)
					watchJob.Stop()
					break
				}

				// Ignore empty events
				if event.Object == nil {
					continue
				}

				pod = event.Object.(*v1.Pod)
				if pod.Status.Phase == v1.PodRunning || pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
					watchJob.Stop()
					return pod, nil
				}
			case <-timeoutChan:
				err = timeoutError("pod", listOptions)
				watchJob.Stop()
				break
			}
			if err != nil {
				watchJob.Stop()
				break
			}
		}
	}

	return nil, err
}

func (client *job) JobExecutionStatus(ctx context.Context, config *JobConfig, executionName string) (JobStatus, error) {
	batchV1 := client.clientSet.BatchV1()
	kubernetesJobs := batchV1.Jobs(config.Namespace)
	listOptions := meta.ListOptions{
		TypeMeta: meta.TypeMeta{
			Kind:       "Job",
			APIVersion: "batch/v1",
		},
		LabelSelector: fmt.Sprintf("job=%s", executionName),
	}

	watchJob, err := kubernetesJobs.Watch(ctx, listOptions)
	if err != nil {
		return Failed, err
	}

	resultChan := watchJob.ResultChan()
	defer watchJob.Stop()
	var event watch.Event
	var jobEvent *batch.Job

	for event = range resultChan {
		if event.Type == watch.Error {
			return FetchError, nil
		}

		jobEvent = event.Object.(*batch.Job)
		if jobEvent.Status.Succeeded >= int32(1) {
			return Succeed, nil
		} else if jobEvent.Status.Failed >= int32(1) {
			return Failed, nil
		}
	}

	return NotFound, nil
}

func (client *job) GetPodLogs(ctx context.Context, pod *v1.Pod) (io.ReadCloser, error) {
	podLogOpts := v1.PodLogOptions{
		Follow: true,
	}
	request := client.clientSet.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
	response, err := request.Stream(ctx)

	if err != nil {
		return nil, err
	}
	return response, nil
}
