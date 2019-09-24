package pohsien

import (
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        corev1 "k8s.io/api/core/v1"
        _"k8s.io/apimachinery/pkg/labels"
        "k8s.io/kubernetes/test/e2e/framework"
        "github.com/onsi/ginkgo"
        "github.com/onsi/gomega"
)

var _ = SIGDescribe("Kubernetes Pohsien's Pod [Pohsien]", func() {
        f := framework.NewDefaultFramework("pods")
        var podClient *framework.PodClient
                ginkgo.BeforeEach(func() {
                podClient = f.PodClient()
        })

        ginkgo.It("Make sure the pohsien pod can be deployed", func(){
                ginkgo.By("Create Pod")
                pod := &corev1.Pod{
                        ObjectMeta: metav1.ObjectMeta{
                                Name: "pohsien",
                                Labels: map[string]string{
                                        "name": "pohsien",
                                },
                        },
                        Spec: corev1.PodSpec{
                                Containers: []corev1.Container{
                                        {
                                                Name:  "nginx",
                                                Image: "nginx:1.17.3",
                                        },
                                },
                        },
                }
                podClient.Create(pod)

                ginkgo.By("Get the pod")
		podGetting, err := podClient.Get(pod.Name, metav1.GetOptions{})
                framework.ExpectNoError(err, "Failed to get the pod")
                framework.ExpectNoError(f.WaitForPodRunning(podGetting.Name))
                gomega.Expect(podGetting.Name, "pohsien")

        })
})

