// Copyright 2019 Mark Spicer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package shopmanager provides an application for managing the prepared.org workshop.
package shopmanager

import (
	pb "github.com/theprepared-dot-org/ste332/api/shopmanager"
	"github.com/theprepared-dot-org/ste332/pkg/application"
)

const (
	name = "shop-manager"
	version = "0.1.0"
	banner = `
   SSSSSSSSSSSSSSS      tttt                             333333333333333    333333333333333    222222222222222
 SS:::::::::::::::S  ttt:::t                            3:::::::::::::::33 3:::::::::::::::33 2:::::::::::::::22
S:::::SSSSSS::::::S  t:::::t                            3::::::33333::::::33::::::33333::::::32::::::222222:::::2
S:::::S     SSSSSSS  t:::::t                            3333333     3:::::33333333     3:::::32222222     2:::::2
S:::::S        ttttttt:::::ttttttt        eeeeeeeeeeee              3:::::3            3:::::3            2:::::2
S:::::S        t:::::::::::::::::t      ee::::::::::::ee            3:::::3            3:::::3            2:::::2
 S::::SSSS     t:::::::::::::::::t     e::::::eeeee:::::ee  33333333:::::3     33333333:::::3          2222::::2
  SS::::::SSSSStttttt:::::::tttttt    e::::::e     e:::::e  3:::::::::::3      3:::::::::::3      22222::::::22
    SSS::::::::SS    t:::::t          e:::::::eeeee::::::e  33333333:::::3     33333333:::::3   22::::::::222
       SSSSSS::::S   t:::::t          e:::::::::::::::::e           3:::::3            3:::::3 2:::::22222
            S:::::S  t:::::t          e::::::eeeeeeeeeee            3:::::3            3:::::32:::::2
            S:::::S  t:::::t    tttttte:::::::e                     3:::::3            3:::::32:::::2
SSSSSSS     S:::::S  t::::::tttt:::::te::::::::e        3333333     3:::::33333333     3:::::32:::::2       222222
S::::::SSSSSS:::::S  tt::::::::::::::t e::::::::eeeeeeee3::::::33333::::::33::::::33333::::::32::::::2222222:::::2
S:::::::::::::::SS     tt:::::::::::tt  ee:::::::::::::e3:::::::::::::::33 3:::::::::::::::33 2::::::::::::::::::2
 SSSSSSSSSSSSSSS         ttttttttttt      eeeeeeeeeeeeee 333333333333333    333333333333333   22222222222222222222

Ste332 Shop Manager
Version %s

`
)

// ShopManager provides an application definition for the shopmanager gRPC server.
type ShopManager struct {
	app *application.Application
}

// NewShopManager provides an instantiated ShopManager server ready to run.
func NewShopManager() (*ShopManager, error) {
	app, err := application.NewApplication(name, banner, version)
	if err != nil {
		return nil, err
	}

	shopManger := &ShopManager{
		app: app,
	}

	pb.RegisterShopManagerServer(app.Server, shopManger)

	return shopManger, nil
}

// Run starts the ShopManager gRPC server and admin interface.
func (s *ShopManager) Run() error {
	return s.app.Run()
}

// ListUsers returns a list of users.
func (s *ShopManager) ListUsers(empty *pb.Empty, stream pb.ShopManager_ListUsersServer) error {
	return nil
}
