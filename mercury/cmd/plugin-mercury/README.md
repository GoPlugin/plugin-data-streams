This directory houses the Mercury LOOPP

# Running Integration Tests Locally

Running the tests is as simple as
- building this binary
- setting the CL_MERCURY_CMD env var to the *fully resolved* binary path
- running the test(s)


The interesting tests are `TestIntegration_MercuryV*` in ` github.com/goplugin/pluginv3.0/core/services/ocr2/plugins/mercury`

In detail:
```
sh

go install # builds `mercury` binary in this dir
CL_MERCURY_CMD=plugin-mercury go test -v -timeout 120s -run ^TestIntegration_MercuryV github.com/goplugin/pluginv3.0/core/services/ocr2/plugins/mercury 2>&1 | tee /tmp/mercury_loop.log
```