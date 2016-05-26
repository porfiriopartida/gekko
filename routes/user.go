package routes

import (
	"fmt"
	//"github.com/tonymtz/gekko/services"
	//"github.com/tonymtz/gekko/server/config"
	"github.com/tonymtz/gekko/repos"
	"github.com/tonymtz/gekko/server/status"
	"github.com/labstack/echo"
	//"github.com/tonymtz/gekko/models"
	//"github.com/tonymtz/gekko/services/oauth2"
	//"encoding/json"
	"strconv"
)


/**
* Used to declare some globals
*/
func init(){
}

//var myProviders map[string]oauth2.IProvider

type UserRoute interface {
	Get(echo.Context) error
	Post(echo.Context) error
}

var User userRoute

type userRoute struct {
	UserRoute
	usersRepository repos.UsersRepository
}
func parseInt(str string) int64{
	//var e error
	intId, e := strconv.ParseInt(str, 10, 32)
	if(e != nil){
		return -1
	}
	return intId
}
func (this *userRoute) Get(ctx echo.Context) error {
	fmt.Println("User.GET");
	//var id string
	id  := ctx.Param("id")
	//var id = ctx.QueryParam("id");
	fmt.Println("User.GET Context: ");
	fmt.Println("id: " + id);

	if id == "" {
		return ctx.JSON(200, ctx)
		//return echo.NewHTTPError(status.BAD_REQUEST, "Must provide a code")
	} else {
		intId := parseInt(id)
		if intId < 0 {
			return echo.NewHTTPError(status.BAD_REQUEST, "Invalid user id (" + id + ")")
		}
		return ctx.JSON(200, "Getting user" + id)
		//return echo.NewHTTPError(status.BAD_REQUEST, "Must provide a code")
	}

	////fmt.Println(ctx);
	////var jsonString string
	////jsonString = json.Marshal(ctx.QueryParams())

}
func (this *userRoute) Post(ctx echo.Context) error {
	fmt.Println("User.POST");
	//fmt.Println(ctx);
	//var jsonString string
	//jsonString = json.Marshal(ctx.QueryParams())
	return ctx.JSON(200, ctx)

	return echo.NewHTTPError(status.OK, "POST")
}
