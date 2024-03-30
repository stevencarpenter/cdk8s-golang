package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus28/v2"
)

type PubgCronChartProps struct {
	cdk8s.ChartProps
}

type NginxChartProps struct {
	cdk8s.ChartProps
}

type ServiceRoleChartProps struct {
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

func NewPubgServerChart(scope constructs.Construct, id string, props *NginxChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	redisPassword := cdk8splus28.Secret_FromSecretName(chart, jsii.String("redis-pass"), jsii.String("redis-pass"))

	dep := cdk8splus28.NewDeployment(chart, jsii.String("deployment"), &cdk8splus28.DeploymentProps{
		Metadata:       &cdk8s.ApiObjectMetadata{Name: jsii.String("pubgserver")},
		ServiceAccount: cdk8splus28.ServiceAccount_FromServiceAccountName(scope, jsii.String("sa-pubgserver"), jsii.String("sa-pubg"), nil)},
	)

	dep.AddContainer(&cdk8splus28.ContainerProps{
		Name:            jsii.String("pubgserver"),
		Image:           jsii.String("pubgserver:latest"),
		ImagePullPolicy: cdk8splus28.ImagePullPolicy_IF_NOT_PRESENT,
		PortNumber:      jsii.Number(8090),
		Liveness: cdk8splus28.Probe_FromTcpSocket(&cdk8splus28.TcpSocketProbeOptions{
			InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(10)),
			PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(5)),
		}),
		Readiness: cdk8splus28.Probe_FromTcpSocket(&cdk8splus28.TcpSocketProbeOptions{
			InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(10)),
			PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(5)),
		}),
		Startup: cdk8splus28.Probe_FromTcpSocket(&cdk8splus28.TcpSocketProbeOptions{
			InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(10)),
			PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(5)),
		}),
		SecurityContext: &cdk8splus28.ContainerSecurityContextProps{EnsureNonRoot: jsii.Bool(false)},
		EnvVariables: &map[string]cdk8splus28.EnvValue{
			"REDIS_PASSWORD": cdk8splus28.EnvValue_FromSecretValue(&cdk8splus28.SecretValue{Key: jsii.String("redis-pass"), Secret: redisPassword}, nil),
		},
	})

	dep.ExposeViaService(&cdk8splus28.DeploymentExposeViaServiceOptions{
		Name:        jsii.String("pubgserver"),
		ServiceType: cdk8splus28.ServiceType_LOAD_BALANCER,
		Ports:       &[]*cdk8splus28.ServicePort{{Port: jsii.Number(8090), TargetPort: jsii.Number(8090)}}})

	return chart
}

func NewPubgCronChart(scope constructs.Construct, id string, props *PubgCronChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}

	var chart = cdk8s.NewChart(scope, &id, &cprops)

	pubgSecret := cdk8splus28.Secret_FromSecretName(chart, jsii.String("pubg-api-token"), jsii.String("pubg-api-token"))
	redisPassword := cdk8splus28.Secret_FromSecretName(chart, jsii.String("redis-pass"), jsii.String("redis-pass"))

	cdk8splus28.NewCronJob(chart, jsii.String(id), &cdk8splus28.CronJobProps{
		Metadata:       &cdk8s.ApiObjectMetadata{Name: jsii.String(id)},
		ServiceAccount: cdk8splus28.ServiceAccount_FromServiceAccountName(scope, jsii.String("sa-pubgcron"), jsii.String("sa-pubg"), nil),
		Containers: &[]*cdk8splus28.ContainerProps{
			{
				Name:            jsii.String(id),
				Image:           jsii.String("pubg:latest"),
				ImagePullPolicy: cdk8splus28.ImagePullPolicy_IF_NOT_PRESENT,
				PortNumber:      jsii.Number(8091),
				SecurityContext: &cdk8splus28.ContainerSecurityContextProps{EnsureNonRoot: jsii.Bool(false)}, //Not sure why the container needs root
				EnvVariables: &map[string]cdk8splus28.EnvValue{
					"PUBG_TOKEN":     cdk8splus28.EnvValue_FromSecretValue(&cdk8splus28.SecretValue{Key: jsii.String("pubg-api-token"), Secret: pubgSecret}, nil),
					"REDIS_PASSWORD": cdk8splus28.EnvValue_FromSecretValue(&cdk8splus28.SecretValue{Key: jsii.String("redis-pass"), Secret: redisPassword}, nil),
				},
			},
		},
		Schedule: cdk8s.Cron_Schedule(&cdk8s.CronOptions{
			Minute: jsii.String("0"),
		}),
	})

	return chart
}

//func NewPubgServiceAccount(scope constructs.Construct, id string, props *ServiceRoleChartProps) cdk8s.Chart {
//	var cprops cdk8s.ChartProps
//	if props != nil {
//		cprops = props.ChartProps
//	}
//	var chart = cdk8s.NewChart(scope, &id, &cprops)
//
//	serviceAccount := cdk8splus28.NewServiceAccount(chart, jsii.String(id), &cdk8splus28.ServiceAccountProps{
//		Metadata:       &cdk8s.ApiObjectMetadata{Name: jsii.String(id)},
//		AutomountToken: jsii.Bool(true),
//	})
//
//	role := cdk8splus28.NewRole(chart, jsii.String(id), &cdk8splus28.RoleProps{
//		Metadata: &cdk8s.ApiObjectMetadata{Name: jsii.String(id)},
//	})
//
//	role.AllowReadWrite()
//
//	roleBinding := cdk8splus28.NewRoleBinding(chart, jsii.String("pubg-role-binding"), &cdk8splus28.RoleBindingProps{
//		Metadata: &cdk8s.ApiObjectMetadata{Name: jsii.String("pubg-role-binding")},
//		Role:     role,
//	})
//
//	roleBinding.AddSubjects(cdk8splus28.SubjectConfiguration{Kind: jsii.String("ServiceAccount"), Name: serviceAccount.Name(), Namespace: jsii.String("default")})
//	return chart
//}

func main() {
	app := cdk8s.NewApp(nil)
	Redis(app, "redis")
	//NewPubgServiceAccount(app, "pubg-role", nil)
	NewPubgServerChart(app, "pubgserver", nil)
	NewPubgCronChart(app, "pubg", nil)

	app.Synth()
}
