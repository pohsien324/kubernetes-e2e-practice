# A Practice of  Writing Kubernetes E2E Test 
A simple practice of writing Kubernetes E2E test. This test is marked to `[Pohsien]` tag, and it will create the pod named `pohsien` which uses the Nginx image. This can verifies the resource deployment of Kubernetes cluster.

## Notice
* This test only works for Kubernetes **v1.15.x**.
* Make sure you already pulled the [`Kubernetes` repository](https://github.com/kubernetes/kubernetes).
* Make sure you already installed the [`kubetest`](https://github.com/kubernetes/test-infra/tree/master/kubetest). 
## How to Build the Test
You can use [`kubetest`](https://github.com/kubernetes/test-infra/tree/master/kubetest) to build and execute this test.
```bash
$ go get github.com/pohsienshih/kubernetes-e2e-practice

# Change the version of Kubernetes
$ cd $GOPATH/src/k8s.io/kubernetes
$ git checkout v1.15.x

# Copy the test into Kubernetes repository
$ cd $GOPATH/src/github.com/pohsienshih/kubernetes-e2e-practice/
$ cp -r BUILD e2e_test.go pohsien/  \
$GOPATH/src/k8s.io/kubernetes/test/e2e/

# Build the e2e.test
$ cd $GOPATH/src/k8s.io/kubernetes
$ kubetest --build

# Execute the test
$ kubetest --test --test_args="--ginkgo.focus=\[Pohsien\]"
```

## Add the Test Manually
If you don't want to override the Kubernetes default e2e files in above steps or you want to add your own test. You can follow the below steps to add the new test.

1. Create a folder for the test in kubernets/test/e2e.
2. Place the test files into the folder.
3. Add the BUILD file of your test into the folder. 
    * `kubernets/test/e2e/<yourfolder>/BUILD`
4.  Import your package in 
    * `kubernets/test/e2e/e2e_test.go` 
    * `kubernets/test/e2e/BUILD`
5. Rebuild the e2e.test 
    * `kubetest --build`

> Don’t forget to mark the test with specific tag( [Slow], [Serial], etc.). You can consult #kinds-of-tests to determine how your test should be marked.

Refer: 
* [End-to-End Testing in Kubernetes](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/e2e-tests.md)
* [Writing good e2e tests for Kubernetes](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-testing/writing-good-e2e-tests.md)

