/*
 Copyright 2020 The Magma Authors.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package test_init

import (
	"testing"

	"magma/feg/cloud/go/feg"
	feg_service "magma/feg/cloud/go/services/feg"
	builder_servicers "magma/feg/cloud/go/services/feg/servicers/protected"
	"magma/orc8r/cloud/go/orc8r"
	builder_protos "magma/orc8r/cloud/go/services/configurator/mconfig/protos"
	"magma/orc8r/cloud/go/test_utils"
)

func StartTestService(t *testing.T) {
	labels := map[string]string{
		orc8r.MconfigBuilderLabel: "true",
	}

	srv, lis := test_utils.NewTestOrchestratorService(t, feg.ModuleName, feg_service.ServiceName, labels, nil)
	builder_protos.RegisterMconfigBuilderServer(srv.GrpcServer, builder_servicers.NewBuilderServicer())

	go srv.RunTest(lis)
}
