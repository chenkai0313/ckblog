package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blog/controllers:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers:SiteController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/backend/site/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers:SiteController"],
        beego.ControllerComments{
            Method: "Site",
            Router: `/backend/site/site`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:SiteController"] = append(beego.GlobalControllerRouter["blog/controllers:SiteController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/backend/site/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/backend/user/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Out",
            Router: `/backend/user/out`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:UserController"] = append(beego.GlobalControllerRouter["blog/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/backend/user/register`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
