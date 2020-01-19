# Meetup Topics

Each of our meetup session typically has 2 speaking tracks: **beginner track** and **intermediate track**. This is because we want to cover both beginner audience, as well as folks who have been using Kubernetes/Cloud Native tech at work wanting to know what else can they apply or learn failure stories from others.

Note: this curriculum can be revised according to Indonesia community members' needs. **Contribution welcome: add more useful reference links, suggested topics, revise no longer valid topics.**

## Beginner Track

### Objective

For monthly meetup, we will have 12 sessions per year. We want to form a yearly curriculum which can be repeated every year. The curriculum can be divided into 2 categories: (1) Kubernetes CKA Prep; (2) Cloud Native Stack Exploration. 

If a meetup decides to pick category #1, a community member who joins all 12 sessions will learn all core Kubernetes concepts which are acccording to [CKA exam curriculum](https://github.com/cncf/curriculum) designed by CNCF. The expectation is that the community member will be ready to take CKA/CKAD exam, excluding the hands-on exercise.

For category #2, there's no really certification for it, but there are more than [25 graduating/incubating projects](https://www.cncf.io/projects/) and [20+ sandbox projects](https://www.cncf.io/sandbox-projects/) under CNCF ecosystem. The expectation is that a community member will gain basic introduction and use case of a particular CNCF project from the session.

### Suggested Topics

#### Category #1: Kubernetes CKA Prep

1. **Intro to Kubernetes: The Why and Motivation**  
This will be the first door towards learning Kubernetes, and the general overview. What's the reason why Kubernetes becomes popular, when should we start adopting Kubernetes, how is it related to microservice and containers.

2. **Kubernetes Core Concepts: Pod, Deployment, ReplicaSet, Service**  
This will cover the core concepts of Kubernetes, what is a Pod, Deployment, ReplicaSet, and Service. Reference: [Brendan Burns, Technical Overview of Kubernetes](https://www.youtube.com/watch?v=WwBdNXt6wO4).

3. **Kubernetes Cluster Architecture and Installation**  
This topic will briefly cover Kubernetes components in both control plane and worker nodes, and go through the hands-on installation using popular tool i.e. [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/), [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/), [kind](https://github.com/kubernetes-sigs/kind) or if make sense [Kelsey's Kubernetes the hardway](https://github.com/kelseyhightower/kubernetes-the-hard-way).

4. **Understanding Kubernetes API Primitives and Objects**  
This will cover how does [Kubernetes object](https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/) look like, the [API primivites](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.17/), what is spec, what is metadata. How is this related to the [architecture](https://www.youtube.com/watch?v=zeS6OyDoy78).

5. **Understanding Kubernetes Service**  
This will explain [how does Service work](https://kubernetes.io/docs/concepts/services-networking/service/), [life of packets in Kubernetes](https://www.youtube.com/watch?v=0Omvgd7Hg1I), the ins and outs of [Kubernetes networking](https://www.youtube.com/watch?v=0Omvgd7Hg1I).

6. **Kubernetes Networking**  
This is a huge topic, an extention to Kubernetes service, [how does networking work in Kubernetes](https://kubernetes.io/docs/concepts/cluster-administration/networking/), pod to pod communication, [ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/), [what's CNI](https://github.com/containernetworking/cni), etc. Reference: [Tim Hockin, Illustrated Guide to Kubernetes Networking](https://speakerdeck.com/thockin/illustrated-guide-to-kubernetes-networking).

7. **Kubernetes Security and RBAC**  
This topic is an intro [how to secure Kubernetes cluster](https://www.youtube.com/watch?v=YRR-kZub0cA), the [API path of Kubernetes](https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/) (auth, authz, admission), [what's RBAC](https://kubernetes.io/docs/reference/access-authn-authz/rbac/) and how should we configure it. If time permits, we can also cover [Kubernetes Network Policy](https://kubernetes.io/docs/tasks/administer-cluster/declare-network-policy/) and [Pod Security Policy](https://kubernetes.io/docs/concepts/policy/pod-security-policy/).

8. **Kubernetes Storage**  
This topic will cover [types of volumes](https://kubernetes.io/docs/concepts/storage/volumes/), [persistent volumes](https://kubernetes.io/docs/concepts/storage/volumes/), what's a [persistent volume claims](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims), and [how to configure an app to use persistent volumes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-volume-storage/).

9. **Troubleshooting Kubernetes**  
Without understanding Kubernetes, troubleshooting issues in the cluster is not an easy task. This topic will give tips and what to see first when things go wrong: [troubleshoot app failure](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-application/), [troubleshoot clusterr failure](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-cluster/). Reference: [Brandon Philips, Kubernetes Day 2: Cluster Operations](https://www.youtube.com/watch?v=U1zR0eDQRYQ).

10. **Kubernetes Logging and Monitoring**  
This topic will cover [basic logging architecture](https://kubernetes.io/docs/concepts/cluster-administration/logging/), [what to monitor](https://www.datadoghq.com/blog/monitoring-kubernetes-era/), which includes [cluster components](https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/). Reference: [kubectl logs](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#logs).

11. **Intro to Helm**  
Helm is a package manager for Kubernetes, like how `apt` is for `ubuntu`. This will not be in CKA exam but will be very useful, necessary, and practical to someone new to Kubernetes. Reference: [Helm docs](https://helm.sh/docs/intro/).

12. **Extending Kubernetes**  
This is the next step of Kubernetes, good for the last session of curriculum. It won't be in CKA exam but good for someone who has completed the entire core curriculum and looking to figure out what to learn next. This topic will cover [extending Kubernetes](https://kubernetes.io/docs/concepts/extend-kubernetes/extend-cluster/) with [kubectl plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/), [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/), what's [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) and [operator SDK](https://github.com/operator-framework/operator-sdk).


#### Category #2: Cloud Native Stack Exploration

This category is as simple as *Intro to X*, replace X with any CNCF or [cloud native projects](http://l.cncf.io) that are open source. We won't go deep but go breadth to cover as many open source projects as we can. Prioritize the projects that have [graduated or are in incubating stages](https://www.cncf.io/projects/). Popular projects that at least should be covered:
* Container runtime: [Docker](https://docker-curriculum.com/), [containerd](https://containerd.io/), [firecracker](https://github.com/firecracker-microvm/firecracker).
* [Kubernetes](https://kubernetes.io/)
* [Istio](https://istio.io/)
* [Prometheus](https://prometheus.io/)
* [Envoy](https://www.envoyproxy.io/)
* [Open Policy Agent](https://www.openpolicyagent.org/)
* [CoreDNS](https://coredns.io/)
* [gRPC](https://grpc.io/)

## Intermediate Track

### Objective

This track will cover practical use cases of Kubernetes and cloud native technologies adopted by companies, or can cover any intro to CNCF projects (like in Category #2 of Beginner Track) if a meetup decides to pick Category #1 already in the Beginner Track.

### Suggested Topics

This is really free depending on speakers' journey and lessons that they learned or explored that they want to share with the community. For the sake of guidance, some interesting topics:
* Failure stories in companies
* Usecases of Kubernetes and CNCF projects (prometheus, envoy, etc.)
* Istio Deep-Dive
* Toolings to improve Production System
* Autoscaling in Kubernetes (HPA, VPA, CA)
* And many more
