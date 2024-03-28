package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus28/v2"
	"sjcarpenter.com/cdk8s-golang/imports/k8s"
)

type CronJobOptions struct {
	Schedule      string
	Container     k8s.Container
	RestartPolicy string
}

type NginxChartProps struct {
	cdk8s.ChartProps
}

func Redis(scope constructs.Construct, id string) cdk8s.Chart {

	chart := cdk8s.NewChart(scope, jsii.String(id), nil)

	cdk8s.NewHelm(chart, jsii.String("redis"), &cdk8s.HelmProps{
		Chart:   jsii.String("bitnami/redis"),
		Version: jsii.String("19.0.1"),
		Values: &map[string]interface{}{
			"auth": map[string]string{
				"password": "secret",
			},
			"sentinel": map[string]bool{
				"enabled": true,
			},
		},
	})

	return chart
}

//func PubgIngress(scope constructs.Construct, id string) cdk8s.Chart {
//
//	chart := cdk8s.NewChart(scope, jsii.String(id), nil)
//
//	ingress := cdk8splus28.NewIngress(chart, jsii.String("ingress"), nil)
//
//	ingress.AddRule(jsii.String("/"), PubgBackend(chart, "root"), cdk8splus28.HttpIngressPathType_PREFIX)
//	ingress.AddRule(jsii.String("/foo"), PubgBackend(chart, "foo"), cdk8splus28.HttpIngressPathType_PREFIX)
//	ingress.AddRule(jsii.String("/foo/bar"), PubgBackend(chart, "foo-bar"), cdk8splus28.HttpIngressPathType_PREFIX)
//	ingress.AddRule(jsii.String("/pubg/leaderboard/"), PubgBackend(chart, "pubg-leaderboard"), cdk8splus28.HttpIngressPathType_PREFIX)
//
//	ingress.AddHostDefaultBackend(jsii.String("host"), PubgBackend(chart, "host/hey"))
//
//	return chart
//}

func NewNginxChart(scope constructs.Construct, id string, props *NginxChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	dep := cdk8splus28.NewDeployment(chart, jsii.String("deployment"), &cdk8splus28.DeploymentProps{Metadata: &cdk8s.ApiObjectMetadata{Name: jsii.String("nginx-deployment-cdk8s-plus")}})

	ingress := cdk8splus28.NewIngress(chart, jsii.String("ingress"), nil)

	ingress.AddRule(jsii.String("/"), PubgBackend(chart, "root"), cdk8splus28.HttpIngressPathType_PREFIX)
	ingress.AddRule(jsii.String("/foo"), PubgBackend(chart, "foo"), cdk8splus28.HttpIngressPathType_PREFIX)
	ingress.AddRule(jsii.String("/pubg/leaderboard/"), PubgBackend(chart, "pubg-leaderboard"), cdk8splus28.HttpIngressPathType_PREFIX)

	ingress.AddHostDefaultBackend(jsii.String("scjarpenter"), PubgBackend(chart, "scjarpenter"))

	dep.AddContainer(&cdk8splus28.ContainerProps{
		Name:       jsii.String("nginx-container"),
		Image:      jsii.String("nginx"),
		PortNumber: jsii.Number(80),
	})

	dep.AddContainer(&cdk8splus28.ContainerProps{
		Name:       jsii.String("hashicorp-http-echo"),
		Image:      jsii.String("hashicorp/http-echo"),
		Args:       &[]*string{jsii.String("-text"), jsii.String("hello")},
		PortNumber: jsii.Number(6789),
	})

	dep.ExposeViaService(&cdk8splus28.DeploymentExposeViaServiceOptions{
		Name:        jsii.String("nginx-container-service"),
		ServiceType: cdk8splus28.ServiceType_LOAD_BALANCER,
		Ports:       &[]*cdk8splus28.ServicePort{{Port: jsii.Number(9090), TargetPort: jsii.Number(80)}}})

	dep.ExposeViaService(&cdk8splus28.DeploymentExposeViaServiceOptions{
		Name: jsii.String("hashicorp-http-echo-service"),

		ServiceType: cdk8splus28.ServiceType_LOAD_BALANCER,
		Ports:       &[]*cdk8splus28.ServicePort{{Port: jsii.Number(6789), TargetPort: jsii.Number(8090)}}})

	return chart
}

func PubgBackend(scope constructs.Construct, text string) cdk8splus28.IngressBackend {

	deploy := cdk8splus28.NewDeployment(scope, jsii.String(text), &cdk8splus28.DeploymentProps{
		Containers: &[]*cdk8splus28.ContainerProps{
			{
				Image:      jsii.String("hashicorp/http-echo"),
				Args:       &[]*string{jsii.String("-text"), jsii.String(text)},
				PortNumber: jsii.Number(6789),
			},
			//{
			//	Image:      jsii.String("sjcarpenter/pubg"),
			//	Args:       &[]*string{jsii.String("--accountId"), jsii.String(text)},
			//	PortNumber: jsii.Number(9876),
			//},
		},
	})

	return cdk8splus28.IngressBackend_FromService(deploy.ExposeViaService(&cdk8splus28.DeploymentExposeViaServiceOptions{
		Ports: &[]*cdk8splus28.ServicePort{{Port: jsii.Number(6789)}}}), nil)

}

//func LeaderBoardUpdate(scope constructs.Construct, id string) cdk8splus28.CronJob {
//	//chart := cdk8s.NewChart(scope, jsii.String(id), nil)
//
//	return cdk8splus28.CronJob(scope, jsii.String(id), &k8s.KubeCronJobProps{
//		Metadata: k8s.ObjectMeta{
//			Name: jsii.String(id),
//		},
//		Spec: k8s.JobSpec{
//			te,
//		},
//	})
//
//return chart

//}

//
//Schedule: "*/5 * * * *", // Runs every 5 minutes
//JobTemplate: JobSpec{
//Template: PodTemplateSpec{
//Spec: PodSpec{
//Containers: []Container{
//{
//Name:  "example",
//Image: "busybox",
//Args:  []string{"/bin/sh", "-c", "date; echo Hello from the Kubernetes cron job"},
//},
//},
//RestartPolicy: "OnFailure",
//},
//},
//},
//

//func NewCronJob(chart *cdk8s.Chart, name string, options *CronJobOptions) {
//	job := k8s.NewCronJob(chart, name, &k8s.CronJobProps{
//		Metadata: &cdk8s.ApiObjectMetadata{Name: name},
//		Spec: &k8s.CronJobSpec{
//			Schedule: options.Schedule,
//			JobTemplate: &k8s.JobTemplateSpec{
//				Spec: &k8s.JobSpec{
//					Template: &k8s.PodTemplateSpec{
//						Spec: &k8s.PodSpec{
//							Containers:    &[]k8s.Container{options.Container},
//							RestartPolicy: options.RestartPolicy,
//						},
//					},
//				},
//			},
//		},
//	})
//
//	// Additional configuration or manipulation of the job can be done here if needed
//	_ = job // Use or modify the job as needed
//}

//func NewChart(scope constructs.Construct, id string, ns string, appLabel string) cdk8s.Chart {
//
//	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
//		Namespace: jsii.String(ns),
//	})
//
//	labels := map[string]*string{
//		"app": jsii.String(appLabel),
//	}
//
//	k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
//		Spec: &k8s.DeploymentSpec{
//			Replicas: jsii.Number(3),
//			Selector: &k8s.LabelSelector{
//				MatchLabels: &labels,
//			},
//			Template: &k8s.PodTemplateSpec{
//				Metadata: &k8s.ObjectMeta{
//					Labels: &labels,
//				},
//				Spec: &k8s.PodSpec{
//					Containers: &[]*k8s.Container{{
//						Name:  jsii.String("app-container"),
//						Image: jsii.String("nginx:1.19.10"),
//						Ports: &[]*k8s.ContainerPort{{
//							ContainerPort: jsii.Number(80),
//						}},
//					}},
//				},
//			},
//		},
//	})
//
//	return chart
//}

func main() {
	app := cdk8s.NewApp(nil)
	Redis(app, "redis")
	NewNginxChart(app, "nginx-cdk8s-plus", nil)
	//PubgIngress(app, "pubg")
	//PubgBackend(app, "snugmarine")
	//NewChart(app, "pubg", "default", "pubg")

	app.Synth()
}
